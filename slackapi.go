package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SlackAPI struct {
}

func (s *SlackAPI) ReportError(err error) {
	fmt.Printf("Error: %s", err)
	os.Exit(1)
}

func (s *SlackAPI) GetRequest(action string, params ...string) []byte {
	var url string = fmt.Sprintf("https://slack.com/api/%s", action)

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

	fmt.Printf("%s\n", response)
	os.Exit(0)
}
