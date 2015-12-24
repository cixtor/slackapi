package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
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

func (s *SlackAPI) PrintJson(data interface{}) {
	response, err := json.MarshalIndent(data, "", "\x20\x20")

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
	os.Exit(0)
}

func (s *SlackAPI) GetRequest(data interface{}, action string, params ...string) interface{} {
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

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(data)

	if err != nil {
		s.ReportError(err)
	}

	return data
}

func (s *SlackAPI) Test() {
	var response interface{}
	s.GetRequest(&response, "api.test")
	s.PrintJson(response)
}

func (s *SlackAPI) AuthTest() {
	var response interface{}
	s.GetRequest(&response, "auth.test", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) UsersInfo(query string) {
	var response interface{}
	s.GetRequest(&response, "users.info", "token", "user="+query)
	s.PrintJson(response)
}

func (s *SlackAPI) UsersList() {
	var response interface{}
	s.GetRequest(&response, "users.list", "token", "presence=1")
	s.PrintJson(response)
}

func (s *SlackAPI) UsersSearch(query string) {
	if len(query) == 0 {
		s.ReportError(errors.New("empty query is invalid"))
	}

	var response Users
	var matches []User
	s.GetRequest(&response, "users.list", "token", "presence=1")

	for _, user := range response.Members {
		if strings.Contains(user.Name, query) ||
			strings.Contains(user.RealName, query) ||
			strings.Contains(user.Profile.Email, query) {
			matches = append(matches, user)
		}
	}

	s.PrintJson(matches)
}
