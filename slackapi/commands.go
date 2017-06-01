package main

import (
	"fmt"
	"time"

	"github.com/cixtor/slackapi"
)

// ProcessCommandClose closes a private user chat.
func (s *ChatSession) ProcessCommandClose(command string) {
	if !s.IsConnected || command != ":close" {
		return
	}

	response := s.InstantMessageClose(s.Channel)
	PrintInlineJSON(response)

	if response.Ok {
		s.IsConnected = false
		s.Username = "username"
		s.Channel = "channel"
	}
}

// ProcessCommandDelete deletes the latest message sent from this session.
func (s *ChatSession) ProcessCommandDelete(command string) {
	if !s.IsConnected || command != ":delete" {
		return
	}

	totalHistory := len(s.History)

	if totalHistory > 0 {
		var shortHistory []slackapi.Post

		forDeletion := totalHistory - 1
		latestMsg := s.History[forDeletion]
		response := s.ChatDelete(latestMsg.Channel, latestMsg.Timestamp)

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
func (s *ChatSession) ProcessCommandExec(command string) {
	if !s.IsConnected || command != ":exec" {
		return
	}

	response := ShellExec(s.UserInput)
	s.UserInput = fmt.Sprintf("```%s```", response)
	s.SendUserMessage()
}

// ProcessCommandExecv sends the output of a command with the command itself.
func (s *ChatSession) ProcessCommandExecv(command string) {
	if !s.IsConnected || command != ":execv" {
		return
	}

	response := ShellExec(s.UserInput)
	s.UserInput = fmt.Sprintf("```$ %s\n%s```", s.UserInput, response)
	s.SendUserMessage()
}

// ProcessCommandFlush deletes all the messages sent from this session.
func (s *ChatSession) ProcessCommandFlush(command string) {
	if !s.IsConnected || command != ":flush" {
		return
	}

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
func (s *ChatSession) ProcessCommandHistory(command string) {
	if !s.IsConnected || command != ":history" {
		return
	}

	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceHistory(action, s.Channel, s.UserInput)
		PrintFormattedJSON(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandMessages prints all the messages sent from this session.
func (s *ChatSession) ProcessCommandMessages(command string) {
	if !s.IsConnected || command != ":messages" {
		return
	}

	PrintFormattedJSON(s.History)
}

// ProcessCommandMyHistory prints all messages sent by this user account.
func (s *ChatSession) ProcessCommandMyHistory(command string) {
	if !s.IsConnected || command != ":myhistory" {
		return
	}

	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceMyHistory(action, s.Channel, s.UserInput)
		PrintFormattedJSON(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandOpen opens a new chat with a channel, group or user.
func (s *ChatSession) ProcessCommandOpen(command string) {
	if !s.IsConnected || command != ":open" {
		return
	}

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
func (s *ChatSession) ProcessCommandOwner(command string) {
	if !s.IsConnected || command != ":owner" {
		return
	}

	PrintFormattedJSON(s.Owner)
}

// ProcessCommandPurge deletes the entire history from the API service.
func (s *ChatSession) ProcessCommandPurge(command string) {
	if !s.IsConnected || command != ":purge" {
		return
	}

	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		s.ResourcePurgeHistory(action, s.Channel, s.UserInput, true)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

// ProcessCommandRobotImage sets the emoji or URL to use as the avatar.
func (s *ChatSession) ProcessCommandRobotImage(command string) {
	if !s.IsConnected || command != ":robotimage" {
		return
	}

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
func (s *ChatSession) ProcessCommandRobotInfo(command string) {
	if !s.IsConnected || command != ":robotinfo" {
		return
	}

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
func (s *ChatSession) ProcessCommandRobotName(command string) {
	if !s.IsConnected || command != ":robotname" {
		return
	}

	if s.UserInput != "" {
		s.RobotName = s.UserInput
	}
}

// ProcessCommandRobotOff turns the robot session off.
func (s *ChatSession) ProcessCommandRobotOff(command string) {
	if !s.IsConnected || command != ":robotoff" {
		return
	}

	s.RobotIsActive = false
}

// ProcessCommandRobotOn turns the robot session on.
func (s *ChatSession) ProcessCommandRobotOn(command string) {
	if !s.IsConnected || command != ":roboton" {
		return
	}

	s.RobotIsActive = true
}

// ProcessCommandToken sets the API token for this session.
func (s *ChatSession) ProcessCommandToken(command string) {
	if !s.IsConnected || command != ":token" {
		return
	}

	s.Token = s.UserInput

	author := s.AuthTest()

	s.Username = author.User
	s.Channel = author.TeamID
	s.Owner = author
}

// ProcessCommandUpdate modifies the latest message sent from this session.
func (s *ChatSession) ProcessCommandUpdate(command string) {
	if !s.IsConnected || command != ":update" {
		return
	}

	totalHistory := len(s.History)

	if s.UserInput != "" && totalHistory > 0 {
		latest := s.History[totalHistory-1]
		response := s.ChatUpdate(latest.Channel, latest.Timestamp, s.UserInput)
		PrintInlineJSON(response)
	}
}

// ProcessCommandUserID searches the ID of certain user account.
func (s *ChatSession) ProcessCommandUserID(command string) {
	if !s.IsConnected || command != ":userid" {
		return
	}

	uniqueid := s.UsersID(s.UserInput)
	fmt.Printf("@ User identifier: %s\n", uniqueid)
}

// ProcessCommandUserList prints all the users in the API service.
func (s *ChatSession) ProcessCommandUserList(command string) {
	if !s.IsConnected || command != ":userlist" {
		return
	}

	response := s.UsersList()
	PrintFormattedJSON(response)
}

// ProcessCommandUserSearch searches an user account in the API service.
func (s *ChatSession) ProcessCommandUserSearch(command string) {
	if !s.IsConnected || command != ":usersearch" {
		return
	}

	response := s.UsersSearch(s.UserInput)
	PrintFormattedJSON(response)
}
