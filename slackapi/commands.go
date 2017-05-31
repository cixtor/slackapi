package main

import (
	"fmt"
	"time"

	"github.com/cixtor/slackapi"
)

// SendUserMessage sends the user input to the API as a chat message.
func (s *ChatSession) SendUserMessage() {
	// Send the message to the remote service.
	if s.IsConnected {
		response := s.ChatPostMessage(s.Channel, s.UserInput)
		s.History = append(s.History, response)

		if response.Ok {
			fmt.Printf("{\"ok\":true,\"channel\":\"%s\",\"ts\":\"%s\"}\n",
				response.Channel,
				response.Timestamp)
		} else {
			s.PrintInlineJSON(response)
		}
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandClose closes a private user chat.
func (s *ChatSession) ProcessCommandClose() {
	if s.IsConnected {
		response := s.InstantMessageClose(s.Channel)
		s.PrintInlineJSON(response)

		if response.Ok {
			s.IsConnected = false
			s.Username = "username"
			s.Channel = "channel"
		}
	}
}

// ProcessCommandDelete deletes the latest message sent from this session.
func (s *ChatSession) ProcessCommandDelete() {
	totalHistory := len(s.History)

	if totalHistory > 0 {
		var shortHistory []slackapi.Post

		forDeletion := totalHistory - 1
		latestMsg := s.History[forDeletion]
		response := s.ChatDelete(latestMsg.Channel, latestMsg.Timestamp)

		s.PrintInlineJSON(response)

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
	response := s.System(s.UserInput)
	s.UserInput = fmt.Sprintf("```%s```", response)
	s.SendUserMessage()
}

// ProcessCommandExecv sends the output of a command with the command itself.
func (s *ChatSession) ProcessCommandExecv() {
	response := s.System(s.UserInput)
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
		response := s.ChatDelete(message.Channel, message.Timestamp)

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
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceHistory(action, s.Channel, s.UserInput)
		s.PrintFormattedJSON(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandMessages prints all the messages sent from this session.
func (s *ChatSession) ProcessCommandMessages() {
	s.PrintFormattedJSON(s.History)
}

// ProcessCommandMyHistory prints all messages sent by this user account.
func (s *ChatSession) ProcessCommandMyHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceMyHistory(action, s.Channel, s.UserInput)
		s.PrintFormattedJSON(response)
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
			s.PrintInlineJSON(response)
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
			s.PrintInlineJSON(response)
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

		s.PrintInlineJSON(response)
	}
}

// ProcessCommandOwner prints information about this session.
func (s *ChatSession) ProcessCommandOwner() {
	s.PrintFormattedJSON(s.Owner)
}

// ProcessCommandPurge deletes the entire history from the API service.
func (s *ChatSession) ProcessCommandPurge() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		s.ResourcePurgeHistory(action, s.Channel, s.UserInput, true)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandRobotImage sets the emoji or URL to use as the avatar.
func (s *ChatSession) ProcessCommandRobotImage() {
	if s.UserInput != "" {
		s.RobotImage = s.UserInput

		if s.UserInput[0] == ':' {
			s.RobotImageType = slackapi.EMOJI
		} else {
			s.RobotImageType = slackapi.ICONURL
		}
	}
}

// ProcessCommandRobotInfo prints information about the robot session.
func (s *ChatSession) ProcessCommandRobotInfo() {
	fmt.Printf("@ Robot info:\n")
	fmt.Printf("  Robot name: %s\n", s.RobotName)
	fmt.Printf("  Robot type: %s\n", s.RobotImageType)
	fmt.Printf("  Robot image: %s\n", s.RobotImage)

	if s.RobotIsActive {
		fmt.Println("  Robot active: true")
	} else {
		fmt.Println("  Robot active: false")
	}
}

// ProcessCommandRobotName sets the name for the robot session.
func (s *ChatSession) ProcessCommandRobotName() {
	if s.UserInput != "" {
		s.RobotName = s.UserInput
	}
}

// ProcessCommandRobotOff turns the robot session off.
func (s *ChatSession) ProcessCommandRobotOff() {
	s.RobotIsActive = false
}

// ProcessCommandRobotOn turns the robot session on.
func (s *ChatSession) ProcessCommandRobotOn() {
	s.RobotIsActive = true
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
		response := s.ChatUpdate(latest.Channel, latest.Timestamp, s.UserInput)
		s.PrintInlineJSON(response)
	}
}

// ProcessCommandUserID searches the ID of certain user account.
func (s *ChatSession) ProcessCommandUserID() {
	uniqueid := s.UsersID(s.UserInput)
	fmt.Printf("@ User identifier: %s\n", uniqueid)
}

// ProcessCommandUserList prints all the users in the API service.
func (s *ChatSession) ProcessCommandUserList() {
	response := s.UsersList()
	s.PrintFormattedJSON(response)
}

// ProcessCommandUserSearch searches an user account in the API service.
func (s *ChatSession) ProcessCommandUserSearch() {
	response := s.UsersSearch(s.UserInput)
	s.PrintFormattedJSON(response)
}
