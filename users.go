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

func (s *SlackAPI) UsersId(query string) string {
	var identifier string
	response := s.UsersList()

	if response.Ok {
		for _, user := range response.Members {
			if user.Name == query {
				identifier = user.Id
				break
			}
		}
	}

	return identifier
}

func (s *SlackAPI) UsersInfo(query string) {
	var response interface{}
	query = s.UsersId(query)
	s.GetRequest(&response, "users.info", "token", "user="+query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) UsersList() Users {
	if s.TeamUsers.Ok == true {
		return s.TeamUsers
	}

	var response Users
	s.GetRequest(&response, "users.list", "token", "presence=1")
	s.TeamUsers = response

	return response
}

func (s *SlackAPI) UsersListVerbose() {
	response := s.UsersList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) UsersSearch(query string) []User {
	if len(query) == 0 {
		s.ReportError(errors.New("empty query is invalid"))
	}

	var matches []User
	response := s.UsersList()

	for _, user := range response.Members {
		if strings.Contains(user.Name, query) ||
			strings.Contains(user.RealName, query) ||
			strings.Contains(user.Profile.Email, query) {
			matches = append(matches, user)
		}
	}

	return matches
}

func (s *SlackAPI) UsersSearchVerbose(query string) {
	response := s.UsersSearch(query)
	s.PrintAndExit(response)
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
