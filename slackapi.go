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

// SlackAPI defines the base object. It holds the API token, the information of
// the user account associated to such API token, the information for the robot
// session (if the user decides has activated it), a list of all the available
// public channels, a list of all the accessible private groups, and a list of
// the users registered into the Slack team.
type SlackAPI struct {
	owner        Owner
	token        string
	params       map[string]string
	teamUsers    ResponseUsersList
	teamGroups   ResponseGroupsList
	teamChannels ResponseChannelsList
}

// New instantiates a new object.
func New() *SlackAPI {
	var s SlackAPI

	s.params = make(map[string]string)

	return &s
}

// SetToken sets the API token for the session.
func (s *SlackAPI) SetToken(token string) {
	s.token = token
}

// URLEndpoint builds and returns the URL to send the HTTP requests.
func (s *SlackAPI) urlEndpoint(action string, params map[string]string) string {
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
func (s *SlackAPI) httpRequest(method string, body io.Reader, action string, params map[string]string) (*http.Request, error) {
	if len(s.params) > 0 {
		for name, value := range s.params {
			params[name] = value
		}
		s.params = map[string]string{}
	}

	url := s.urlEndpoint(action, params)
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
func (s *SlackAPI) dataToParams(data interface{}) map[string]string {
	if data == nil {
		/* no params except for the API token */
		return map[string]string{"token": s.token}
	}

	var name string
	var value interface{}

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	length := t.NumField() /* struct size */
	params := make(map[string]string, length+1)
	params["token"] = s.token /* API token */

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

		case []string:
			params[name] = strings.Join(value.([]string), ",")

		case []Attachment:
			if out, err := json.Marshal(value); err == nil {
				params[name] = string(out)
			}
		}
	}

	return params
}

// ExecuteRequest sends the HTTP request and decodes the JSON response.
func (s *SlackAPI) executeRequest(req *http.Request, data interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("http;", err)
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("http exec; body close;", err)
		}
	}()

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	if err := json.NewDecoder(tee).Decode(&data); err != nil {
		out, _ := ioutil.ReadAll(&buf) /* bad idea; change */

		if strings.Contains(string(out), "too many requests") {
			fake := "{\"ok\":false, \"error\":\"RATELIMIT\"}"
			read := bytes.NewReader([]byte(fake))

			if err2 := json.NewDecoder(read).Decode(&data); err2 != nil {
				log.Println("http exec; json decode;", err)
			}

			return
		}

		log.Println("ratelimit;", err)
	}
}

// PrintCurlCommand prints the HTTP request object into the console.
func (s *SlackAPI) printCurlCommand(req *http.Request, params map[string]string) {
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
func (s *SlackAPI) getRequest(v interface{}, action string, data interface{}) {
	params := s.dataToParams(data)
	req, err := s.httpRequest("GET", nil, action, params)

	if err != nil {
		log.Println("http get;", err)
		return
	}

	s.printCurlCommand(req, params)
	s.executeRequest(req, &v)
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
	params := s.dataToParams(data)
	writer := multipart.NewWriter(&buffer)

	// Append more HTTP request params.
	for name, value := range params {
		/* Check if the parameter is referencing a file */
		isfile, fpath, fname := s.CheckFileReference(value)

		if !isfile {
			fwriter, _ := writer.CreateFormField(name)
			if _, err := fwriter.Write([]byte(value)); err != nil {
				log.Println("http post; create field;", err)
			}
			continue
		}

		/* Read referenced file and attach to the request */
		resource, err := os.Open(fpath)
		if err != nil {
			log.Println("file open;", err)
			return
		}

		defer func() {
			if err := resource.Close(); err != nil {
				log.Println("http post; file close;", err)
			}
		}()

		fwriter, _ := writer.CreateFormFile(name, fpath)
		if _, err := io.Copy(fwriter, resource); err != nil {
			log.Println("http post; copy param;", err)
			continue
		}

		fwriter, _ = writer.CreateFormField("filename")
		if _, err := fwriter.Write([]byte(fname)); err != nil {
			log.Println("http post; write param;", err)
			continue
		}
	}

	if err := writer.Close(); err != nil {
		log.Println("http post; write close;", err)
		return
	}

	// Now that you have a form, you can submit it to your handler.
	req, err := s.httpRequest("POST", &buffer, action, nil)

	if err != nil {
		log.Println("http post;", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	s.printCurlCommand(req, params)
	s.executeRequest(req, &v)
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
