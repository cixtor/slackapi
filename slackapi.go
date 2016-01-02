package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

func (s *SlackAPI) PrintInlineJson(data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
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

func (s *SlackAPI) GetRequest(data interface{}, action string, params ...string) interface{} {
	var url string = s.Url(action, params)
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

func (s *SlackAPI) ChatPostMessage(channel string, message string) Message {
	var response Message
	s.GetRequest(&response,
		"chat.postMessage",
		"token",
		"channel="+channel,
		"text="+message,
		"as_user=true",
		"link_names=1")
	return response
}

func (s *SlackAPI) ChatPostMessageVerbose(channel string, message string) {
	response := s.ChatPostMessage(channel, message)
	s.PrintJson(response)
}

func (s *SlackAPI) ChatSession() {
	var command string
	var parts []string
	var channel string = "unknown"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("slack:%s> ", channel)
		message, err := reader.ReadString('\n')

		if err != nil {
			s.ReportError(err)
		}

		message = strings.TrimSpace(message)

		if message == ":exit" {
			// Close the chat session and exit the loop.
			fmt.Printf("Closing...")
			response := s.InstantMessagingClose(channel)
			fmt.Printf("\r")
			s.PrintInlineJson(response)
			fmt.Println("Closed")
			break
		} else if message[0] == ':' {
			// Execute a custom command with the message.
			parts = strings.SplitN(message, "\x20", 2)
			command = parts[0]
			message = parts[1]

			if command == ":open" {
				fmt.Printf("Opening session...")
				response := s.InstantMessagingOpen(message)
				fmt.Printf("\r")
				s.PrintInlineJson(response)

				if response.Ok == true {
					channel = response.Channel.Id
				}
			}
		} else {
			// Send the message to the remote service.
			fmt.Println(message)
		}
	}

	os.Exit(0)
}

func (s *SlackAPI) EmojiList() {
	var response interface{}
	s.GetRequest(&response, "emoji.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingClose(query string) Base {
	var response Base
	s.GetRequest(&response, "im.close", "token", "channel="+query)
	return response
}

func (s *SlackAPI) InstantMessagingCloseVerbose(query string) {
	response := s.InstantMessagingClose(query)
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingOpen(query string) Session {
	var response Session

	if query == "slackbot" {
		query = "USLACKBOT"
	}

	s.GetRequest(&response, "im.open", "token", "user="+query)

	return response
}

func (s *SlackAPI) InstantMessagingOpenVerbose(query string) {
	response := s.InstantMessagingOpen(query)
	s.PrintJson(response)
}

func (s *SlackAPI) UsersGetPresence(query string) {
	var response interface{}
	s.GetRequest(&response, "users.getPresence", "token", "user="+query)
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

func (s *SlackAPI) UsersSetActive() {
	var response interface{}
	s.GetRequest(&response, "users.setActive", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) UsersSetPresence(query string) {
	var response interface{}
	s.GetRequest(&response, "users.setPresence", "token", "presence="+query)
	s.PrintJson(response)
}
