package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (s *SlackAPI) ChatSession() {
	reader := bufio.NewReader(os.Stdin)
	s.Username = "username"
	s.Channel = "channel"

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

	os.Exit(0)
}

func (s *SlackAPI) ProcessMessage() {
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

func (s *SlackAPI) ProcessCommand() {
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
	case ":userid":
		s.ProcessCommandUserId()
	case ":userlist":
		s.ProcessCommandUserList()
	case ":usersearch":
		s.ProcessCommandUserSearch()
	}
}

func (s *SlackAPI) SendUserMessage() {
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

func (s *SlackAPI) ProcessCommandClose() {
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

func (s *SlackAPI) ProcessCommandDelete() {
	var totalHistory int = len(s.History)

	if totalHistory > 0 {
		var forDeletion int = totalHistory - 1
		var latestMsg Message = s.History[forDeletion]
		var shortHistory []Message

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

func (s *SlackAPI) ProcessCommandExec() {
	response := s.System(s.UserInput)
	s.UserInput = fmt.Sprintf("```%s```", response)
	s.SendUserMessage()
}

func (s *SlackAPI) ProcessCommandExecv() {
	response := s.System(s.UserInput)
	s.UserInput = fmt.Sprintf("```$ %s\n%s```", s.UserInput, response)
	s.SendUserMessage()
}

func (s *SlackAPI) ProcessCommandFlush() {
	var shortHistory []Message
	var totalHistory int = len(s.History)
	var offset int = (totalHistory - 1)
	var message Message

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
		}
	}

	s.History = shortHistory
}

func (s *SlackAPI) ProcessCommandHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceHistory(action, s.Channel, s.UserInput)
		s.PrintFormattedJson(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *SlackAPI) ProcessCommandMessages() {
	s.PrintFormattedJson(s.History)
}

func (s *SlackAPI) ProcessCommandMyHistory() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		response := s.ResourceMyHistory(action, s.Channel, s.UserInput)
		s.PrintFormattedJson(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *SlackAPI) ProcessCommandOpen() {
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

func (s *SlackAPI) ProcessCommandOwner() {
	s.PrintFormattedJson(s.Owner)
}

func (s *SlackAPI) ProcessCommandPurge() {
	if s.IsChannelConn || s.IsGroupConn || s.IsUserConn {
		var action string = fmt.Sprintf("%s.history", s.MethodName)
		s.ResourcePurgeHistory(action, s.Channel, s.UserInput)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *SlackAPI) ProcessCommandRobotImage() {
	if s.UserInput != "" {
		s.RobotImage = s.UserInput

		if s.UserInput[0] == ':' {
			s.RobotImageType = "emoji"
		} else {
			s.RobotImageType = "icon_url"
		}
	}
}

func (s *SlackAPI) ProcessCommandRobotInfo() {
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

func (s *SlackAPI) ProcessCommandRobotName() {
	if s.UserInput != "" {
		s.RobotName = s.UserInput
	}
}

func (s *SlackAPI) ProcessCommandRobotOff() {
	s.RobotIsActive = false
}

func (s *SlackAPI) ProcessCommandRobotOn() {
	s.RobotIsActive = true
}

func (s *SlackAPI) ProcessCommandToken() {
	s.Token = s.UserInput
	s.Owner = s.AuthTest()
}

func (s *SlackAPI) ProcessCommandUserId() {
	uniqueid := s.UsersId(s.UserInput)
	fmt.Printf("@ User identifier: %s\n", uniqueid)
}

func (s *SlackAPI) ProcessCommandUserList() {
	response := s.UsersList()
	s.PrintFormattedJson(response)
}

func (s *SlackAPI) ProcessCommandUserSearch() {
	response := s.UsersSearch(s.UserInput)
	s.PrintFormattedJson(response)
}
