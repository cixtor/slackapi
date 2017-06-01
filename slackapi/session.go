package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cixtor/slackapi"
)

// ChatSession holds information about the user account associated to the passed
// API token, the message history, and the latest user input that will be passed
// to the API service. Here is also included multiple properties denotating if
// the session is currently in a public channel, a group or a private user chat.
type ChatSession struct {
	slackapi.SlackAPI
	Command       string
	Channel       string
	Username      string
	UserInput     string
	MethodName    string
	History       []slackapi.Post
	IsUserConn    bool
	IsConnected   bool
	IsGroupConn   bool
	IsChannelConn bool
}

// NewSession instantiates a new object.
func NewSession() *ChatSession {
	return &ChatSession{}
}

// StartSession initiates a loop to pass user input to the API.
func (s *ChatSession) StartSession() {
	reader := bufio.NewReader(os.Stdin)

	s.Channel = "channel"
	s.Username = "username"

	if s.Token != "" {
		author := s.AuthTest()
		s.Username = author.User
		s.Channel = author.TeamID
		s.Owner = author
	}

	for {
		fmt.Printf("%s:%s> ", s.Username, s.Channel)
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("input;", err)
			break
		}

		s.UserInput = strings.TrimSpace(message)

		/* stop infinite loop */
		if s.UserInput == ":exit" {
			break
		}

		s.ProcessMessage()
	}
}

// ProcessMessage takes user input and sends it to the API as a message. If the
// message is prefixed with a colon character the method will execute one of the
// supported custom commands with the rest of the message as its input.
func (s *ChatSession) ProcessMessage() {
	// Ignore empty messages.
	if s.UserInput == "" {
		return
	}

	if s.UserInput[0] == ':' {
		var parts []string
		parts = strings.SplitN(s.UserInput, "\x20", 2)
		s.Command = parts[0]
		s.UserInput = ""

		if len(parts) == 2 {
			s.UserInput = parts[1]
		}

		s.ProcessCommandClose(s.Command)
		s.ProcessCommandDelete(s.Command)
		s.ProcessCommandExec(s.Command)
		s.ProcessCommandExecv(s.Command)
		s.ProcessCommandFlush(s.Command)
		s.ProcessCommandHistory(s.Command)
		s.ProcessCommandMessages(s.Command)
		s.ProcessCommandMyHistory(s.Command)
		s.ProcessCommandOpen(s.Command)
		s.ProcessCommandOwner(s.Command)
		s.ProcessCommandPurge(s.Command)
		s.ProcessCommandRobotImage(s.Command)
		s.ProcessCommandRobotInfo(s.Command)
		s.ProcessCommandRobotName(s.Command)
		s.ProcessCommandRobotOff(s.Command)
		s.ProcessCommandRobotOn(s.Command)
		s.ProcessCommandToken(s.Command)
		s.ProcessCommandUpdate(s.Command)
		s.ProcessCommandUserID(s.Command)
		s.ProcessCommandUserList(s.Command)
		s.ProcessCommandUserSearch(s.Command)
		return
	}

	s.SendUserMessage()
}

// SendUserMessage sends the user input to the API as a chat message.
func (s *ChatSession) SendUserMessage() {
	// Send the message to the remote service.
	if !s.IsConnected {
		fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
		return
	}

	response := s.ChatPostMessage(s.Channel, s.UserInput)
	s.History = append(s.History, response)

	if !response.Ok {
		PrintInlineJSON(response)
		return
	}

	fmt.Printf(
		"{\"ok\":true,\"channel\":\"%s\",\"ts\":\"%s\"}\n",
		response.Channel,
		response.Timestamp)
}
