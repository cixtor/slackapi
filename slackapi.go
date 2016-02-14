package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type SlackAPI struct {
	Token          string
	Channel        string
	Command        string
	UserInput      string
	RobotName      string
	RobotImage     string
	RobotImageType string
	RobotIsActive  bool
	IsConnected    bool
	TeamUsers      Users
	TeamRooms      Rooms
	TeamGroups     Groups
	History        []Message
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

func (s *SlackAPI) HttpRequest(method string, data interface{}, action string, params []string) error {
	var url string = s.Url(action, params)
	req, err := http.NewRequest(method, url, nil)
	client := &http.Client{}

	req.Header.Add("DNT", "1")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Language", "en-US,en")
	req.Header.Add("Origin", "https://sucuri.slack.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")

	if err != nil {
		s.ReportError(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		s.ReportError(err)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&data)
}

func (s *SlackAPI) GetRequest(data interface{}, action string, params ...string) {
	err := s.HttpRequest("GET", &data, action, params)

	if err != nil {
		s.ReportError(err)
	}
}
