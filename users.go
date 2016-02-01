package main

import (
	"errors"
	"strings"
)

func (s *SlackAPI) UsersGetPresence(query string) {
	var response interface{}
	s.GetRequest(&response, "users.getPresence", "token", "user="+query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) UsersInfo(query string) {
	var response interface{}
	s.GetRequest(&response, "users.info", "token", "user="+query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) UsersList() {
	var response interface{}
	s.GetRequest(&response, "users.list", "token", "presence=1")
	s.PrintAndExit(response)
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

	s.PrintAndExit(matches)
}

func (s *SlackAPI) UsersSetActive() {
	var response interface{}
	s.GetRequest(&response, "users.setActive", "token")
	s.PrintAndExit(response)
}

func (s *SlackAPI) UsersSetPresence(value string) {
	var response interface{}
	s.GetRequest(&response, "users.setPresence", "token", "presence="+value)
	s.PrintAndExit(response)
}
