package slackapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const version = "2.0.2"

// SlackAPI defines the base object. It holds the API token, the information of
// the user account associated to such API token, the information for the robot
// session (if the user decides has activated it), a list of all the available
// public channels, a list of all the accessible private groups, and a list of
// the users registered into the Slack team.
type SlackAPI struct {
	Owner         Owner
	Token         string
	RequestParams map[string]string
	TeamUsers     ResponseUsersList
	TeamGroups    ResponseGroupsList
	TeamChannels  ResponseChannelsList
}

// New instantiates a new object.
func New() *SlackAPI {
	return &SlackAPI{}
}

// Version returns the package version number.
func (s *SlackAPI) Version() string {
	return version
}

// SetToken sets the API token for the session.
func (s *SlackAPI) SetToken(token string) {
	s.Token = token
}

// AutoConfigure sets the API token from an environment variable.
func (s *SlackAPI) AutoConfigure() {
	s.Token = os.Getenv("SLACK_TOKEN")
}

// URL builds and returns the URL to send the HTTP requests.
func (s *SlackAPI) URL(action string, params map[string]string) string {
	data := url.Values{}
	url := "https://slack.com/api/" + action

	for name, value := range params {
		data.Add(name, value)
	}

	if encoded := data.Encode(); encoded != "" {
		url += "?" + encoded
	}

	return url
}

// HTTPRequest builds an HTTP request object and attaches the action parameters.
func (s *SlackAPI) HTTPRequest(method string, body io.Reader, action string, params map[string]string) (*http.Request, error) {
	if len(s.RequestParams) > 0 {
		for name, value := range s.RequestParams {
			params[name] = value
		}
		s.RequestParams = map[string]string{}
	}

	url := s.URL(action, params)
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")

	return req, nil
}

// DataToParams converts a template into a HTTP request parameter map.
func (s *SlackAPI) DataToParams(data interface{}) map[string]string {
	if data == nil {
		/* no params except for the API token */
		return map[string]string{"token": s.Token}
	}

	var name string
	var value interface{}

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	length := t.NumField() /* struct size */
	params := make(map[string]string, length+1)
	params["token"] = s.Token /* API token */

	for i := 0; i < length; i++ {
		name = t.Field(i).Tag.Get("json")
		value = v.Field(i).Interface()

		switch v.Field(i).Interface().(type) {
		case int:
			params[name] = fmt.Sprintf("%d", value)

		case bool:
			if value.(bool) {
				params[name] = "true"
			} else {
				params[name] = "false"
			}

		case string:
			params[name] = value.(string)

		case []Attachment:
			if out, err := json.Marshal(value); err == nil {
				params[name] = string(out)
			}
		}
	}

	return params
}

// AddRequestParam adds an additional parameter to the HTTP request.
func (s *SlackAPI) AddRequestParam(name string, value string) {
	s.RequestParams[name] = value
}

// ExecuteRequest sends the HTTP request and decodes the JSON response.
func (s *SlackAPI) ExecuteRequest(req *http.Request, data interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("HTTP;", err)
		return
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)
	if err := json.NewDecoder(tee).Decode(&data); err != nil {
		out, _ := ioutil.ReadAll(&buf)
		if strings.Contains(string(out), "too many requests") {
			fake := "{\"ok\":false, \"error\":\"RATELIMIT\"}"
			read := bytes.NewReader([]byte(fake))
			json.NewDecoder(read).Decode(&data)
			return
		}
		log.Println("ratelimit;", err)
	}
}

// PrintCurlCommand prints the HTTP request object into the console.
func (s *SlackAPI) PrintCurlCommand(req *http.Request, params map[string]string) {
	if os.Getenv("SLACK_VERBOSE") != "true" {
		return
	}

	fmt.Printf("curl -X %s \"%s\"", req.Method, req.URL)

	for header, values := range req.Header {
		fmt.Printf(" \x5c\n-H \"%s: %s\"", header, values[0])
	}

	fmt.Printf(" \x5c\n-H \"Host: %s\"", req.Host)

	if req.Method == "POST" {
		for name, value := range params {
			fmt.Printf(" \x5c\n-d \"%s=%s\"", name, value)
		}
	}

	fmt.Println()
}

// GetRequest sends a HTTP GET request to the API and returns the response.
func (s *SlackAPI) GetRequest(v interface{}, action string, data interface{}) {
	params := s.DataToParams(data)
	req, err := s.HTTPRequest("GET", nil, action, params)

	if err != nil {
		log.Println("HTTP GET;", err)
		return
	}

	s.PrintCurlCommand(req, params)
	s.ExecuteRequest(req, &v)
}

// PostRequest sends a HTTP POST request to the API and returns the response. If
// one of the request parameters is prefixed with an AT symbol the method will
// assume that the user is trying to load a local file, it will proceed to check
// this by locating such file in the disk, then will attach the data into the
// HTTP request object and upload it to the API. Alternatively, if the file does
// not exists, the method will send the parameter with the apparent filename as
// a string value.
func (s *SlackAPI) PostRequest(v interface{}, action string, data interface{}) {
	var buffer bytes.Buffer
	params := s.DataToParams(data)
	writer := multipart.NewWriter(&buffer)

	// Append more HTTP request params.
	for name, value := range params {
		/* Check if the parameter is referencing a file */
		isfile, fpath, fname := s.CheckFileReference(value)

		if !isfile {
			fwriter, _ := writer.CreateFormField(name)
			fwriter.Write([]byte(value))
			continue
		}

		/* Read referenced file and attach to the request */
		resource, err := os.Open(fpath)
		if err != nil {
			log.Println("file open;", err)
			return
		}
		defer resource.Close()
		fwriter, _ := writer.CreateFormFile(name, fpath)
		io.Copy(fwriter, resource) /* attach file data */
		fwriter, _ = writer.CreateFormField("filename")
		fwriter.Write([]byte(fname))
	}

	writer.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := s.HTTPRequest("POST", &buffer, action, nil)

	if err != nil {
		log.Println("HTTP POST;", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	s.PrintCurlCommand(req, params)
	s.ExecuteRequest(req, &v)
}

// CheckFileReference checks if a HTTP request parameter is a file reference.
func (s *SlackAPI) CheckFileReference(text string) (bool, string, string) {
	if len(text) < 2 || text[0] != '@' {
		return false, "", ""
	}

	fpath := text[1:]

	/* Check if the file actually exists */
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return false, "", ""
	}

	parts := strings.Split(fpath, "/")
	fname := parts[len(parts)-1]

	return true, fpath, fname
}
