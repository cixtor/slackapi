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
		// Send the message to the remote service.
		if s.IsConnected {
			response := s.ChatPostMessage(s.Channel, s.UserInput)
			s.History = append(s.History, response)
			s.PrintInlineJson(response)
		} else {
			fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
		}
	}
}

func (s *SlackAPI) ProcessCommand() {
	if s.Command == ":open" {
		s.ProcessCommandOpen()
	} else if s.Command == ":delete" {
		s.ProcessCommandDelete()
	} else if s.Command == ":history" {
		s.PrintFormattedJson(s.History)
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

func (s *SlackAPI) CloseSession() {
	if s.IsConnected {
		response := s.InstantMessagingClose(s.Channel)
		s.PrintInlineJson(response)
		s.IsConnected = false
	}
}
