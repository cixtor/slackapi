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
	}
}

func (s *SlackAPI) SendUserMessage() {
	// Send the message to the remote service.
	if s.IsConnected {
		response := s.ChatPostMessage(s.Channel, s.UserInput)
		s.History = append(s.History, response)
		s.PrintInlineJson(response)
	} else {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
	}
}

func (s *SlackAPI) ProcessCommandOpen() {
	response := s.InstantMessagingOpen(s.UserInput)
	s.PrintInlineJson(response)

	if response.Ok == true {
		s.Channel = response.Channel.Id
		s.IsConnected = response.Ok
	}
}

func (s *SlackAPI) ProcessCommandDelete() {
	var forDeletion int = len(s.History) - 1
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

func (s *SlackAPI) CloseSession() {
	if s.IsConnected {
		response := s.InstantMessagingClose(s.Channel)
		s.PrintInlineJson(response)
		s.IsConnected = false
	}
}
