package main

import (
	"errors"
	"strings"
)

func (s *SlackAPI) ApiTest() {
	var response interface{}
	s.GetRequest(&response, "api.test")
	s.PrintJson(response)
}

func (s *SlackAPI) AuthTest() {
	var response interface{}
	s.GetRequest(&response, "auth.test", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsHistory(channel string, latest string) {
	var response interface{}

	if latest == "" {
		s.GetRequest(&response,
			"channels.history",
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"unreads=1")
	} else {
		s.GetRequest(&response,
			"channels.history",
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"latest="+latest,
			"unreads=1")
	}
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsInfo(channel string) {
	var response interface{}
	s.GetRequest(&response, "channels.info", "token", "channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsList() {
	var response interface{}
	s.GetRequest(&response, "channels.list", "token", "exclude_archived=0")
	s.PrintJson(response)
}

func (s *SlackAPI) ChatDeleteVerbose(channel string, timestamp string) {
	response := s.ChatDelete(channel, timestamp)
	s.PrintJson(response)
}

func (s *SlackAPI) ChatPostMessageVerbose(channel string, message string) {
	response := s.ChatPostMessage(channel, message)
	s.PrintJson(response)
}

func (s *SlackAPI) EmojiList() {
	var response interface{}
	s.GetRequest(&response, "emoji.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsInfo(channel string) {
	var response interface{}
	s.GetRequest(&response, "groups.info", "token", "channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsList() {
	var response interface{}
	s.GetRequest(&response, "groups.list", "token", "exclude_archived=0")
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingCloseVerbose(channel string) {
	response := s.InstantMessagingClose(channel)
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingList() {
	var response interface{}
	s.GetRequest(&response, "im.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingOpenVerbose(userid string) {
	response := s.InstantMessagingOpen(userid)
	s.PrintJson(response)
}

func (s *SlackAPI) MultiPartyInstantMessagingList() {
	var response interface{}
	s.GetRequest(&response, "mpim.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsAdd(name string, channel string, timestamp string) {
	var response interface{}

	// Remove reaction from a file, file comment, or message.
	if channel[0] == 'F' {
		s.GetRequest(&response,
			"reactions.add",
			"token",
			"name="+name,
			"file="+channel)
	} else {
		s.GetRequest(&response,
			"reactions.add",
			"token",
			"name="+name,
			"channel="+channel,
			"timestamp="+timestamp)
	}

	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsGet(channel string, timestamp string) {
	var response interface{}

	// Remove reaction from a file, file comment, or message.
	if channel[0] == 'F' {
		s.GetRequest(&response,
			"reactions.get",
			"token",
			"file="+channel)
	} else {
		s.GetRequest(&response,
			"reactions.get",
			"token",
			"channel="+channel,
			"timestamp="+timestamp)
	}

	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsList(userid string) {
	var response interface{}

	if userid == "" {
		s.GetRequest(&response,
			"reactions.list",
			"token",
			"full=true",
			"count=100")
	} else {
		s.GetRequest(&response,
			"reactions.list",
			"token",
			"full=true",
			"count=100",
			"user="+userid)
	}
	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsRemove(name string, channel string, timestamp string) {
	var response interface{}

	// Remove reaction from a file, file comment, or message.
	if channel[0] == 'F' {
		s.GetRequest(&response,
			"reactions.remove",
			"token",
			"name="+name,
			"file="+channel)
	} else {
		s.GetRequest(&response,
			"reactions.remove",
			"token",
			"name="+name,
			"channel="+channel,
			"timestamp="+timestamp)
	}

	s.PrintJson(response)
}

func (s *SlackAPI) TeamInfo() {
	var response interface{}
	s.GetRequest(&response, "team.info", "token")
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

func (s *SlackAPI) UsersSetPresence(value string) {
	var response interface{}
	s.GetRequest(&response, "users.setPresence", "token", "presence="+value)
	s.PrintJson(response)
}
