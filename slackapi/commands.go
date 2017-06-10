package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/cixtor/slackapi"
)

// ProcessCommandClose closes a private user chat.
func (s *ChatSession) ProcessCommandClose() {
	response := s.InstantMessageClose(s.Channel)
	PrintInlineJSON(response)

	if response.Ok {
		s.IsConnected = false
		s.Username = "username"
		s.Channel = "channel"
	}
}

// ProcessCommandDelete deletes the latest message sent from this session.
func (s *ChatSession) ProcessCommandDelete() {
	totalHistory := len(s.History)

	if totalHistory > 0 {
		var shortHistory []slackapi.Post

		forDeletion := totalHistory - 1
		latestMsg := s.History[forDeletion]
		response := s.ChatDelete(slackapi.MessageArgs{
			Channel: latestMsg.Channel,
			Ts:      latestMsg.Timestamp,
		})

		PrintInlineJSON(response)

		if response.Ok {
			for key := 0; key < forDeletion; key++ {
				shortHistory = append(shortHistory, s.History[key])
			}

			s.History = shortHistory
		}
	}
}

// ProcessCommandExec sends the output of a command executed locally.
func (s *ChatSession) ProcessCommandExec() {
	response := ShellExec(s.UserInput)
	s.UserInput = fmt.Sprintf("```%s```", response)
	s.SendUserMessage()
}

// ProcessCommandExecv sends the output of a command with the command itself.
func (s *ChatSession) ProcessCommandExecv() {
	response := ShellExec(s.UserInput)
	s.UserInput = fmt.Sprintf("```$ %s\n%s```", s.UserInput, response)
	s.SendUserMessage()
}

// ProcessCommandFlush deletes all the messages sent from this session.
func (s *ChatSession) ProcessCommandFlush() {
	var message slackapi.Post
	var shortHistory []slackapi.Post

	totalHistory := len(s.History)
	offset := (totalHistory - 1)

	fmt.Printf("@ Deleting %d messages\n", totalHistory)

	for key := offset; key >= 0; key-- {
		message = s.History[key]
		fmt.Printf("\x20 %s from %s ", message.Timestamp, message.Channel)
		response := s.ChatDelete(slackapi.MessageArgs{
			Channel: message.Channel,
			Ts:      message.Timestamp,
		})

		if response.Ok {
			fmt.Println("\u2714")
		} else {
			shortHistory = append(shortHistory, message)
			fmt.Printf("\u2718 %s\n", response.Error)

			if response.Error == "RATELIMIT" {
				time.Sleep(10 * time.Second)
			}
		}
	}

	s.History = shortHistory
}

// ProcessCommandHistory prints the entire history in the API service.
func (s *ChatSession) ProcessCommandHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		PrintFormattedJSON(s.ResourceHistory(
			s.MethodName+".history",
			slackapi.HistoryArgs{
				Channel: s.Channel,
				Latest:  s.UserInput,
			}))
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandMessages prints all the messages sent from this session.
func (s *ChatSession) ProcessCommandMessages() {
	PrintFormattedJSON(s.History)
}

// ProcessCommandMyHistory prints all messages sent by this user account.
func (s *ChatSession) ProcessCommandMyHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		PrintFormattedJSON(s.ResourceMyHistory(
			s.MethodName+".history",
			s.Channel,
			s.UserInput))
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandOpen opens a new chat with a channel, group or user.
func (s *ChatSession) ProcessCommandOpen() {
	if s.UserInput != "" {
		uniqueid := s.UsersID(s.UserInput)
		response := s.InstantMessageOpen(uniqueid)

		if response.Error == "user_not_found" {
			PrintInlineJSON(response)
			uniqueid = s.ChannelsID(s.UserInput)

			if uniqueid != s.UserInput {
				response.Ok = true
				response.Error = ""
				response.Channel.ID = uniqueid
				s.MethodName = "channels"
				s.IsChannelConn = true
			} else {
				response.Error = "channel_not_found"
			}
		}

		if response.Error == "channel_not_found" {
			PrintInlineJSON(response)
			uniqueid = s.GroupsID(s.UserInput)

			if uniqueid != s.UserInput {
				response.Ok = true
				response.Error = ""
				response.Channel.ID = uniqueid
				s.MethodName = "groups"
				s.IsGroupConn = true
			} else {
				response.Error = "group_not_found"
			}
		}

		if response.Ok {
			s.Username = s.UserInput
			s.Channel = response.Channel.ID
			s.IsConnected = response.Ok

			if !s.IsChannelConn && !s.IsGroupConn {
				s.MethodName = "im"
				s.IsUserConn = true
			}
		}

		PrintInlineJSON(response)
	}
}

// ProcessCommandOwner prints information about this session.
func (s *ChatSession) ProcessCommandOwner() {
	PrintFormattedJSON(s.Owner)
}

// ProcessCommandPurge deletes the entire history from the API service.
func (s *ChatSession) ProcessCommandPurge() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		s.ResourcePurgeHistory(
			s.MethodName+".history",
			s.Channel,
			s.UserInput,
			true)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandStatus sets the user profile status message.
func (s *ChatSession) ProcessCommandStatus() {
	parts := strings.SplitN(s.UserInput, "\x20", 2)
	s.UsersSetStatus(parts[0], parts[1])
}

// ProcessCommandToken sets the API token for this session.
func (s *ChatSession) ProcessCommandToken() {
	s.Token = s.UserInput

	author := s.AuthTest()

	s.Username = author.User
	s.Channel = author.TeamID
	s.Owner = author
}

// ProcessCommandUpdate modifies the latest message sent from this session.
func (s *ChatSession) ProcessCommandUpdate() {
	totalHistory := len(s.History)

	if s.UserInput != "" && totalHistory > 0 {
		latest := s.History[totalHistory-1]
		response := s.ChatUpdate(slackapi.MessageArgs{
			Channel: latest.Channel,
			Ts:      latest.Timestamp,
			Text:    s.UserInput,
		})
		PrintInlineJSON(response)
	}
}

// ProcessCommandUserID searches the ID of certain user account.
func (s *ChatSession) ProcessCommandUserID() {
	fmt.Printf("@ User identifier: %s\n", s.UsersID(s.UserInput))
}

// ProcessCommandUserList prints all the users in the API service.
func (s *ChatSession) ProcessCommandUserList() {
	response := s.UsersList()
	PrintFormattedJSON(response)
}

// ProcessCommandUserSearch searches an user account in the API service.
func (s *ChatSession) ProcessCommandUserSearch() {
	response := s.UsersSearch(s.UserInput)
	PrintFormattedJSON(response)
}
