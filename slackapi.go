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
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

// SlackAPI defines the base object. It holds the API token, the information of
// the user account associated to such API token, the information for the robot
// session (if the user decides has activated it), a list of all the available
// public channels, a list of all the accessible private groups, and a list of
// the users registered into the Slack team.
type SlackAPI struct {
	token        string
	cookie       string
	client       *http.Client
	debug        bool
	teamUsers    ResponseUsersList
	teamChannels ResponseChannelsList
}

// New instantiates a new object.
func New() *SlackAPI {
	return &SlackAPI{
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

// SetToken sets the API token for the session.
func (s *SlackAPI) SetToken(token string) {
	s.token = token
}

// SetDebug instructs the library to print all HTTP requests.
func (s *SlackAPI) SetDebug(enable bool) {
	s.debug = enable
}

// SetCookie sets the API cookie for the session. Slack changed the permissions
// of all their tokens; now if you inspect the HTTP requests from a web browser
// session and copy the token from there without copying the cookies, the other
// requests will fail.
func (s *SlackAPI) SetCookie(cookie string) {
	s.cookie = cookie
}

// SetTimeout sets the maximum amount of time to wait for the HTTP request.
func (s *SlackAPI) SetTimeout(t time.Duration) {
	s.client.Timeout = t
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

func (s *SlackAPI) sendRequest(req *http.Request, output interface{}) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")

	if s.cookie != "" {
		// NOTES(cixtor): some tokens are only accepted if a valid HTTP cookie
		// is passed with the rest of the request. For example, tokens created
		// by the web authorization flow.
		req.Header.Set("Cookie", s.cookie)
	}

	if s.debug {
		reqText, err := httputil.DumpRequestOut(req, true)

		if err != nil {
			fmt.Println("httputil.DumpRequestOut", err)
		} else {
			fmt.Printf("%s\n\n", reqText)
		}
	}

	res, err := s.client.Do(req)

	if err != nil {
		return fmt.Errorf("cannot http.Client.Do: %s", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(fmt.Errorf("cannot res.Body.Close: %s", err))
		}
	}()

	reader := io.LimitReader(res.Body, 2<<22)
	rawJSON, err := ioutil.ReadAll(reader)

	if err != nil {
		return fmt.Errorf("cannot ioutil.ReadAll: %s", err)
	}

	if s.debug {
		resText, err := httputil.DumpResponse(res, false)

		if err != nil {
			fmt.Println("httputil.DumpResponse", err)
		} else {
			fmt.Printf("%s", resText)
		}

		fmt.Printf("%s\n\n", rawJSON)
	}

	// NOTES(cixtor): output is expected to be a pointer to a variable.
	if err := json.Unmarshal(rawJSON, output); err != nil {
		return fmt.Errorf("cannot json.Unmarshal: %s", err)
	}

	return nil
}

// HTTPRequest builds an HTTP request object and attaches the action parameters.
func (s *SlackAPI) httpRequest(method string, body io.Reader, action string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, s.urlEndpoint(action, params), body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")
	req.Header.Add("Cookie", s.cookie)

	return req, nil
}

// DataToParams converts a template into a HTTP request parameter map.
func (s *SlackAPI) dataToParams(input interface{}) map[string]string {
	if input == nil {
		/* no params except for the API token */
		return map[string]string{"token": s.token}
	}

	var name string
	var value interface{}

	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

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
			if value != "" {
				params[name] = value.(string)
			}

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
func (s *SlackAPI) executeRequest(req *http.Request, input interface{}) {
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

	if err := json.NewDecoder(tee).Decode(&input); err != nil {
		out, _ := ioutil.ReadAll(&buf) /* bad idea; change */

		if strings.Contains(string(out), "too many requests") {
			fake := "{\"ok\":false, \"error\":\"RATELIMIT\"}"
			read := bytes.NewReader([]byte(fake))

			if err2 := json.NewDecoder(read).Decode(&input); err2 != nil {
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
func (s *SlackAPI) getRequest(input interface{}, action string, output interface{}) {
	params := s.dataToParams(output)
	req, err := s.httpRequest("GET", nil, action, params)

	if err != nil {
		log.Println("http get;", err)
		return
	}

	s.printCurlCommand(req, params)
	s.executeRequest(req, &input)
}

// PostRequest sends a HTTP POST request to the API and returns the response. If
// one of the request parameters is prefixed with an AT symbol the method will
// assume that the user is trying to load a local file, it will proceed to check
// this by locating such file in the disk, then will attach the data into the
// HTTP request object and upload it to the API. Alternatively, if the file does
// not exists, the method will send the parameter with the apparent filename as
// a string value.
func (s *SlackAPI) postRequest(input interface{}, action string, output interface{}) {
	var buffer bytes.Buffer
	params := s.dataToParams(output)
	writer := multipart.NewWriter(&buffer)

	for name, value := range params {
		/* Check if the parameter is referencing a file */
		isfile, fpath, fname := s.checkFileReference(value)

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
	s.executeRequest(req, &input)
}

// CheckFileReference checks if a HTTP request parameter is a file reference.
func (s *SlackAPI) checkFileReference(text string) (bool, string, string) {
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

func (s *SlackAPI) anyGET(targetURL string, input url.Values, output interface{}) error {
	if params := input.Encode(); params != "" {
		targetURL += "?" + input.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, targetURL, nil)

	if err != nil {
		return fmt.Errorf("cannot http.NewRequest.GET: %s", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return s.sendRequest(req, output)
}

func (s *SlackAPI) jsonPOST(targetURL string, input interface{}, output interface{}) error {
	binput, err := json.Marshal(input)

	if err != nil {
		return fmt.Errorf("cannot json.Marshal `%#v`: %s", input, err)
	}

	req, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewBuffer(binput))

	if err != nil {
		return fmt.Errorf("cannot http.NewRequest.POST: %s", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return s.sendRequest(req, output)
}

func (s *SlackAPI) paramPOST(targetURL string, input url.Values, output interface{}) error {
	req, err := http.NewRequest(http.MethodPost, targetURL, strings.NewReader(input.Encode()))

	if err != nil {
		return fmt.Errorf("cannot http.NewRequest.POST: %s", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return s.sendRequest(req, output)
}

func (s *SlackAPI) baseGET(endpoint string, input url.Values, output interface{}) error {
	return s.anyGET("https://slack.com"+endpoint, input, output)
}

func (s *SlackAPI) edgeGET(endpoint string, input url.Values, output interface{}) error {
	return s.anyGET("https://edgeapi.slack.com"+endpoint, input, output)
}

func (s *SlackAPI) baseJSONPOST(endpoint string, input interface{}, output interface{}) error {
	return s.jsonPOST("https://slack.com"+endpoint, input, output)
}

func (s *SlackAPI) baseFormPOST(endpoint string, input url.Values, output interface{}) error {
	return s.paramPOST("https://slack.com"+endpoint, input, output)
}

func (s *SlackAPI) edgePOST(endpoint string, input interface{}, output interface{}) error {
	return s.jsonPOST("https://edgeapi.slack.com"+endpoint, input, output)
}

func addFileToReq(writer *multipart.Writer, name string, filename string) (int64, error) {
	w, err := writer.CreateFormFile(name, filename)

	if err != nil {
		return 0, err
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	return io.Copy(w, file)
}
