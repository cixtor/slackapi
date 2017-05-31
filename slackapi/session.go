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
	Channel       string
	Command       string
	History       []slackapi.Post
	IsChannelConn bool
	IsConnected   bool
	IsGroupConn   bool
	IsUserConn    bool
	MethodName    string
	UserInput     string
	Username      string
}

// NewSession instantiates a new object.
func NewSession() *ChatSession {
	return &ChatSession{}
}

// StartSession initiates a loop to pass user input to the API.
func (s *ChatSession) StartSession() {
	reader := bufio.NewReader(os.Stdin)

	s.Username = "username"
	s.Channel = "channel"

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

// ProcessMessage takes user input and sends it to the API as a message. If the
// message is prefixed with a colon character the method will execute one of the
// supported custom commands with the rest of the message as its input.
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

// ProcessCommand executes the corresponding custom command.
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
		s.ProcessCommandUserID()
	case ":userlist":
		s.ProcessCommandUserList()
	case ":usersearch":
		s.ProcessCommandUserSearch()
	}
}
