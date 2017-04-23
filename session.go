package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type ChatSession struct {
	SlackAPI
	Channel       string
	Command       string
	History       []Post
	IsChannelConn bool
	IsConnected   bool
	IsGroupConn   bool
	IsUserConn    bool
	MethodName    string
	UserInput     string
	Username      string
}

func (s *ChatSession) StartChatSession() {
	reader := bufio.NewReader(os.Stdin)

	s.Username = "username"
	s.Channel = "channel"

	if s.Token != "" {
		author := s.AuthTest()
		s.Username = author.User
		s.Channel = author.TeamId
		s.Owner = author
	}

	for {
		fmt.Printf("%s:%s> ", s.Username, s.Channel)
		message, err := reader.ReadString('\n')

		if err != nil {
			s.ReportError(err)
		}

		s.UserInput = strings.TrimSpace(message)

		if s.UserInput == ":exit" {
			break
		} else {
			s.ProcessMessage()
		}
	}
}

func (s *ChatSession) ProcessMessage() {
	var parts []string

	if s.UserInput == "" {
		// Ignore empty messages.
	} else if s.UserInput[0] == ':' {
		parts = strings.SplitN(s.UserInput, "\x20", 2)
		s.Command = parts[0]
		s.UserInput = ""

		if len(parts) == 2 {
			s.UserInput = parts[1]
		}

		s.ProcessCommand()
	} else {
		s.SendUserMessage()
	}
}

func (s *ChatSession) ProcessCommand() {
	switch s.Command {
	case ":close":
		s.ProcessCommandClose()
	case ":delete":
		s.ProcessCommandDelete()
	case ":exec":
		s.ProcessCommandExec()
	case ":execv":
		s.ProcessCommandExecv()
	case ":flush":
		s.ProcessCommandFlush()
	case ":history":
		s.ProcessCommandHistory()
	case ":messages":
		s.ProcessCommandMessages()
	case ":myhistory":
		s.ProcessCommandMyHistory()
	case ":open":
		s.ProcessCommandOpen()
	case ":owner":
		s.ProcessCommandOwner()
	case ":purge":
		s.ProcessCommandPurge()
	case ":robotimage":
		s.ProcessCommandRobotImage()
	case ":robotinfo":
		s.ProcessCommandRobotInfo()
	case ":robotname":
		s.ProcessCommandRobotName()
	case ":robotoff":
		s.ProcessCommandRobotOff()
	case ":roboton":
		s.ProcessCommandRobotOn()
	case ":token":
		s.ProcessCommandToken()
	case ":update":
		s.ProcessCommandUpdate()
	case ":userid":
		s.ProcessCommandUserId()
	case ":userlist":
		s.ProcessCommandUserList()
	case ":usersearch":
		s.ProcessCommandUserSearch()
	}
}

func (s *ChatSession) SendUserMessage() {
	// Send the message to the remote service.
	if s.IsConnected {
		response := s.ChatPostMessage(s.Channel, s.UserInput)
		s.History = append(s.History, response)

		if response.Ok == true {
			fmt.Printf("{\"ok\":true,\"channel\":\"%s\",\"ts\":\"%s\"}\n",
				response.Channel,
				response.Ts)
		} else {
			s.PrintInlineJson(response)
		}
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *ChatSession) ProcessCommandClose() {
	if s.IsConnected {
		response := s.InstantMessagingClose(s.Channel)
		s.PrintInlineJson(response)

		if response.Ok {
			s.IsConnected = false
			s.Username = "username"
			s.Channel = "channel"
		}
	}
}

func (s *ChatSession) ProcessCommandDelete() {
	var totalHistory int = len(s.History)

	if totalHistory > 0 {
		var forDeletion int = totalHistory - 1
		var latestMsg Post = s.History[forDeletion]
		var shortHistory []Post

		response := s.ChatDelete(latestMsg.Channel, latestMsg.Ts)
		s.PrintInlineJson(response)

		if response.Ok == true {
			for key := 0; key < forDeletion; key++ {
				shortHistory = append(shortHistory, s.History[key])
			}

			s.History = shortHistory
		}
	}
}

func (s *ChatSession) ProcessCommandExec() {
	response := s.System(s.UserInput)
	s.UserInput = fmt.Sprintf("```%s```", response)
	s.SendUserMessage()
}

func (s *ChatSession) ProcessCommandExecv() {
	response := s.System(s.UserInput)
	s.UserInput = fmt.Sprintf("```$ %s\n%s```", s.UserInput, response)
	s.SendUserMessage()
}

func (s *ChatSession) ProcessCommandFlush() {
	var shortHistory []Post
	var totalHistory int = len(s.History)
	var offset int = (totalHistory - 1)
	var message Post

	fmt.Printf("@ Deleting %d messages\n", totalHistory)

	for key := offset; key >= 0; key-- {
		message = s.History[key]
		fmt.Printf("\x20 %s from %s ", message.Ts, message.Channel)
		response := s.ChatDelete(message.Channel, message.Ts)

		if response.Ok == true {
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

func (s *ChatSession) ProcessCommandHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceHistory(action, s.Channel, s.UserInput)
		s.PrintFormattedJson(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *ChatSession) ProcessCommandMessages() {
	s.PrintFormattedJson(s.History)
}

func (s *ChatSession) ProcessCommandMyHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceMyHistory(action, s.Channel, s.UserInput)
		s.PrintFormattedJson(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *ChatSession) ProcessCommandOpen() {
	if s.UserInput != "" {
		uniqueid := s.UsersId(s.UserInput)
		response := s.InstantMessagingOpen(uniqueid)

		if response.Error == "user_not_found" {
			s.PrintInlineJson(response)
			uniqueid = s.ChannelsId(s.UserInput)

			if uniqueid != s.UserInput {
				response.Ok = true
				response.Error = ""
				response.Channel.Id = uniqueid
				s.MethodName = "channels"
				s.IsChannelConn = true
			} else {
				response.Error = "channel_not_found"
			}
		}

		if response.Error == "channel_not_found" {
			s.PrintInlineJson(response)
			uniqueid = s.GroupsId(s.UserInput)

			if uniqueid != s.UserInput {
				response.Ok = true
				response.Error = ""
				response.Channel.Id = uniqueid
				s.MethodName = "groups"
				s.IsGroupConn = true
			} else {
				response.Error = "group_not_found"
			}
		}

		if response.Ok == true {
			s.Username = s.UserInput
			s.Channel = response.Channel.Id
			s.IsConnected = response.Ok

			if !s.IsChannelConn && !s.IsGroupConn {
				s.MethodName = "im"
				s.IsUserConn = true
			}
		}

		s.PrintInlineJson(response)
	}
}

func (s *ChatSession) ProcessCommandOwner() {
	s.PrintFormattedJson(s.Owner)
}

func (s *ChatSession) ProcessCommandPurge() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		s.ResourcePurgeHistory(action, s.Channel, s.UserInput, true)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *ChatSession) ProcessCommandRobotImage() {
	if s.UserInput != "" {
		s.RobotImage = s.UserInput

		if s.UserInput[0] == ':' {
			s.RobotImageType = "emoji"
		} else {
			s.RobotImageType = "icon_url"
		}
	}
}

func (s *ChatSession) ProcessCommandRobotInfo() {
	fmt.Printf("@ Robot info:\n")
	fmt.Printf("  Robot name: %s\n", s.RobotName)
	fmt.Printf("  Robot type: %s\n", s.RobotImageType)
	fmt.Printf("  Robot image: %s\n", s.RobotImage)

	if s.RobotIsActive == true {
		fmt.Println("  Robot active: true")
	} else {
		fmt.Println("  Robot active: false")
	}
}

func (s *ChatSession) ProcessCommandRobotName() {
	if s.UserInput != "" {
		s.RobotName = s.UserInput
	}
}

func (s *ChatSession) ProcessCommandRobotOff() {
	s.RobotIsActive = false
}

func (s *ChatSession) ProcessCommandRobotOn() {
	s.RobotIsActive = true
}

func (s *ChatSession) ProcessCommandToken() {
	s.Token = s.UserInput

	author := s.AuthTest()

	s.Username = author.User
	s.Channel = author.TeamId
	s.Owner = author
}

func (s *ChatSession) ProcessCommandUpdate() {
	var totalHistory int = len(s.History)

	if s.UserInput != "" && totalHistory > 0 {
		var latest Post = s.History[totalHistory-1]
		response := s.ChatUpdate(latest.Channel, latest.Ts, s.UserInput)
		s.PrintInlineJson(response)
	}
}

func (s *ChatSession) ProcessCommandUserId() {
	uniqueid := s.UsersId(s.UserInput)
	fmt.Printf("@ User identifier: %s\n", uniqueid)
}

func (s *ChatSession) ProcessCommandUserList() {
	response := s.UsersList()
	s.PrintFormattedJson(response)
}

func (s *ChatSession) ProcessCommandUserSearch() {
	response := s.UsersSearch(s.UserInput)
	s.PrintFormattedJson(response)
}
