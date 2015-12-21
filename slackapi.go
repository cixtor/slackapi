package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SlackAPI struct {
	Token string
}

func (s *SlackAPI) AutoConfigure() {
	s.Token = os.Getenv("SLACK_TOKEN")
}

func (s *SlackAPI) ReportError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}

func (s *SlackAPI) PrintResponse(data []byte) {
	var temp interface{}
	err := json.Unmarshal(data, &temp)

	if err != nil {
		s.ReportError(err)
	}

	response, err := json.MarshalIndent(temp, "", "\x20\x20")

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
	os.Exit(0)
}

func (s *SlackAPI) GetRequest(action string, params ...string) []byte {
	var url string = fmt.Sprintf("https://slack.com/api/%s", action)

	if len(params) > 0 {
		var anchor string

		for key, keyvalue := range params {
			if key == 0 {
				anchor = "?"
			} else {
				anchor = "&"
			}

			if keyvalue == "token" {
				keyvalue += "=" + s.Token
			}

			url += anchor + keyvalue
		}
	}

	req, err := http.NewRequest("GET", url, nil)
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
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		s.ReportError(err)
	}

	return body
}

func (s *SlackAPI) Test() {
	response := s.GetRequest("api.test")
	s.PrintResponse(response)
}

func (s *SlackAPI) AuthTest() {
	response := s.GetRequest("auth.test", "token")
	s.PrintResponse(response)
}

func (s *SlackAPI) UsersList() {
	response := s.GetRequest("users.list", "token", "presence=1")
	s.PrintResponse(response)
}
