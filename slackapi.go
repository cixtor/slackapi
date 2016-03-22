package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type SlackAPI struct {
	Token          string
	Username       string
	Channel        string
	Command        string
	UserInput      string
	MethodName     string
	RobotName      string
	RobotImage     string
	RobotImageType string
	RobotIsActive  bool
	IsConnected    bool
	IsUserConn     bool
	IsGroupConn    bool
	IsChannelConn  bool
	Owner          Owner
	TeamUsers      ResponseUsersList
	TeamRooms      Rooms
	TeamGroups     ResponseGroupsList
	History        []Post
}

func (s *SlackAPI) AutoConfigure() {
	s.Token = os.Getenv("SLACK_TOKEN")

	s.RobotName = "bender"
	s.RobotImage = ":robot_face:"
	s.RobotImageType = "emoji"
	s.RobotIsActive = false
}

func (s *SlackAPI) ReportError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
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

func (s *SlackAPI) ExecuteRequest(req *http.Request, data interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		s.ReportError(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		s.ReportError(err)
	}
}

func (s *SlackAPI) GetRequest(data interface{}, action string, params ...string) {
	req := s.HttpRequest("GET", nil, action, params)
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

		if len(parts) == 2 {
			if parts[0] == "file" {
				// Get real filepath.
				var filepath string = parts[1]
				filepath = filepath[1:len(filepath)]

				// Get short filename.
				parts = strings.Split(filepath, "/")
				filename := parts[len(parts)-1]

				// Read local file data.
				resource, _ := os.Open(filepath)
				defer resource.Close()

				// Attach the data read from the local file.
				fwriter, _ := writer.CreateFormFile("file", filepath)
				io.Copy(fwriter, resource)

				// Append another HTTP parameter with the filename.
				fwriter, _ = writer.CreateFormField("filename")
				fwriter.Write([]byte(filename))
			} else {
				fwriter, _ := writer.CreateFormField(parts[0])
				fwriter.Write([]byte(parts[1]))
			}
		}
	}

	writer.Close()

	// Now that you have a form, you can submit it to your handler.
	req := s.HttpRequest("POST", &buffer, action, nil)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	s.ExecuteRequest(req, &data)
}
