package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type SlackAPI struct {
	Owner          Owner
	RequestParams  []string
	RobotImage     string
	RobotImageType string
	RobotIsActive  bool
	RobotName      string
	TeamChannels   ResponseChannelsList
	TeamGroups     ResponseGroupsList
	TeamUsers      ResponseUsersList
	Token          string
}

func (s *SlackAPI) SetToken(token string) {
	s.Token = token
}

func (s *SlackAPI) AutoConfigure() {
	s.Token = os.Getenv("SLACK_TOKEN")

	s.RobotName = "bender"
	s.RobotImage = ":robot_face:"
	s.RobotImageType = "emoji"
	s.RobotIsActive = false
}

func (s *SlackAPI) System(kommand string) []byte {
	var binary string
	var parts []string
	var arguments []string

	if kommand == "" {
		s.ReportError(errors.New("invalid empty command"))
	}

	parts = strings.Fields(kommand)
	binary = parts[0]
	arguments = parts[1:len(parts)]

	response, err := exec.Command(binary, arguments...).Output()

	if err != nil {
		s.ReportError(err)
	}

	return bytes.Trim(response, "\n")
}

func (s *SlackAPI) Url(action string, params []string) string {
	data := url.Values{}
	var parts []string
	var encoded string
	var url string = fmt.Sprintf("https://slack.com/api/%s", action)

	for _, keyvalue := range params {
		if keyvalue == "token" {
			keyvalue += "=" + s.Token
		}

		parts = strings.SplitN(keyvalue, "=", 2)

		if len(parts) == 2 {
			data.Add(parts[0], parts[1])
		}
	}

	encoded = data.Encode()

	if encoded != "" {
		url = fmt.Sprintf("%s?%s", url, encoded)
	}

	return url
}

func (s *SlackAPI) HttpRequest(method string, body io.Reader, action string, params []string) *http.Request {
	if len(s.RequestParams) > 0 {
		params = append(params, s.RequestParams...)
		s.RequestParams = s.RequestParams[:0]
	}

	var url string = s.Url(action, params)
	req, err := http.NewRequest(method, url, body)

	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")

	if err != nil {
		s.ReportError(err)
	}

	return req
}

func (s *SlackAPI) AddRequestParam(param string, value string) {
	var parameter string = fmt.Sprintf("%s=%s", param, value)
	s.RequestParams = append(s.RequestParams, parameter)
}

func (s *SlackAPI) ExecuteRequest(req *http.Request, data interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		s.ReportError(err)
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
		s.ReportError(err)
	}
}

func (s *SlackAPI) PrintCurlCommand(req *http.Request, params []string) {
	if os.Getenv("SLACK_VERBOSE") == "true" {
		fmt.Printf("curl -X %s \"%s\"", req.Method, req.URL)

		for header, values := range req.Header {
			fmt.Printf(" \x5c\n-H \"%s: %s\"", header, values[0])
		}

		fmt.Printf(" \x5c\n-H \"Host: %s\"", req.Host)

		for _, param := range params {
			if param == "token" {
				param = "token=" + s.Token
			}
			fmt.Printf(" \x5c\n-d \"%s\"", param)
		}

		fmt.Println()
	}
}

func (s *SlackAPI) GetRequest(data interface{}, action string, params ...string) {
	req := s.HttpRequest("GET", nil, action, params)
	s.PrintCurlCommand(req, params)
	s.ExecuteRequest(req, &data)
}

func (s *SlackAPI) PostRequest(data interface{}, action string, params ...string) {
	var parts []string
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Append more HTTP request params.
	for _, keyvalue := range params {
		if keyvalue == "token" {
			keyvalue += "=" + s.Token
		}

		parts = strings.SplitN(keyvalue, "=", 2)

		/* Ignore empty parameters */
		if len(parts) != 2 {
			continue
		}

		/* Check if the parameter is referencing a file */
		isfile, fpath, fname := s.CheckFileReference(parts[1])

		if !isfile {
			fwriter, _ := writer.CreateFormField(parts[0])
			fwriter.Write([]byte(parts[1]))
			continue
		}

		if len(parts[1]) > 0 && parts[1][0] == '@' {
		}

		/* Read referenced file and attach to the request */
		resource, err := os.Open(fpath)
		if err != nil {
			s.ReportError(err)
		}
		defer resource.Close()
		fwriter, _ := writer.CreateFormFile(parts[0], fpath)
		io.Copy(fwriter, resource) /* attach file data */
		fwriter, _ = writer.CreateFormField("filename")
		fwriter.Write([]byte(fname))
	}

	writer.Close()

	// Now that you have a form, you can submit it to your handler.
	req := s.HttpRequest("POST", &buffer, action, nil)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	s.PrintCurlCommand(req, params)
	s.ExecuteRequest(req, &data)
}

func (s *SlackAPI) CheckFileReference(text string) (bool, string, string) {
	if len(text) < 2 || text[0] != '@' {
		return false, "", ""
	}

	fpath := text[1:len(text)]

	/* Check if the file actually exists */
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return false, "", ""
	}

	parts := strings.Split(fpath, "/")
	fname := parts[len(parts)-1]

	return true, fpath, fname
}

func (s *SlackAPI) PrintFormattedJson(data interface{}) {
	response, err := json.MarshalIndent(data, "", "\x20\x20")

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
}

func (s *SlackAPI) PrintInlineJson(data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
}

func (s *SlackAPI) PrintAndExit(data interface{}) {
	s.PrintFormattedJson(data)
	os.Exit(0)
}

func (s *SlackAPI) ReportError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
