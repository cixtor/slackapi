package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (s *SlackAPI) ChatSession() {
	s.Channel = "unknown"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("slack:%s> ", s.Channel)
		message, err := reader.ReadString('\n')

		if err != nil {
			s.ReportError(err)
		}

		s.UserInput = strings.TrimSpace(message)

		if s.UserInput == ":exit" {
			break
		} else if s.UserInput == ":close" {
			s.CloseSession()
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
	case ":history":
		s.PrintFormattedJson(s.History)
	case ":open":
		s.ProcessCommandOpen()
	case ":delete":
		s.ProcessCommandDelete()
	case ":flush":
		s.ProcessCommandFlush()
	case ":exec":
		s.ProcessCommandExec()
	case ":execv":
		s.ProcessCommandExecv()
	case ":boton":
		s.ProcessCommandRobotOn()
	case ":botoff":
		s.ProcessCommandRobotOff()
	case ":botinfo":
		s.ProcessCommandRobotInfo()
	case ":botname":
		s.ProcessCommandRobotName()
	case ":botimage":
		s.ProcessCommandRobotImage()
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

func (s *SlackAPI) ProcessCommandOpen() {
	uniqueid := s.UsersId(s.UserInput)
	response := s.InstantMessagingOpen(uniqueid)

	if response.Error == "user_not_found" {
		s.PrintInlineJson(response)
		uniqueid = s.ChannelsId(s.UserInput)

		if uniqueid != s.UserInput {
			response.Ok = true
			response.Error = ""
			response.Channel.Id = uniqueid
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
		} else {
			response.Error = "group_not_found"
		}
	}

	if response.Ok == true {
		s.Channel = response.Channel.Id
		s.IsConnected = response.Ok
	}

	s.PrintInlineJson(response)
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

func (s *SlackAPI) ProcessCommandRobotOn() {
	s.RobotIsActive = true
}

func (s *SlackAPI) ProcessCommandRobotOff() {
	s.RobotIsActive = false
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

func (s *SlackAPI) CloseSession() {
	if s.IsConnected {
		response := s.InstantMessagingClose(s.Channel)
		s.PrintInlineJson(response)
		s.IsConnected = false
	}
}
