package slackapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ChatSession holds information about the user account associated to the passed
// API token, the message history, and the latest user input that will be passed
// to the API service. Here is also included multiple properties denotating if
// the session is currently in a public channel, a group or a private user chat.
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

// NewSession instantiates a new object.
func NewSession() *ChatSession {
	return &ChatSession{}
}

// StartChatSession initiates a loop to pass user input to the API.
func (s *ChatSession) StartChatSession() {
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
		var shortHistory []Post

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
	var message Post
	var shortHistory []Post

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
			s.RobotImageType = EMOJI
		} else {
			s.RobotImageType = ICONURL
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
