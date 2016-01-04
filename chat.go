package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (s *SlackAPI) ChatDelete(channel string, timestamp string) ChannelEvent {
	var response ChannelEvent
	s.GetRequest(&response,
		"chat.delete",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

func (s *SlackAPI) ChatPostMessage(channel string, message string) Message {
	var response Message
	s.GetRequest(&response,
		"chat.postMessage",
		"token",
		"channel="+channel,
		"text="+message,
		"as_user=true",
		"link_names=1")
	return response
}

func (s *SlackAPI) ChatSession() {
	var command string
	var parts []string
	var channel string = "unknown"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("slack:%s> ", channel)
		message, err := reader.ReadString('\n')

		if err != nil {
			s.ReportError(err)
		}

		message = strings.TrimSpace(message)

		if message == ":exit" {
			// Close the chat session and exit the loop.
			fmt.Printf("Closing...")
			response := s.InstantMessagingClose(channel)
			fmt.Printf("\r")
			s.PrintInlineJson(response)
			fmt.Println("Closed")
			break
		} else if message[0] == ':' {
			// Execute a custom command with the message.
			parts = strings.SplitN(message, "\x20", 2)
			command = parts[0]
			message = parts[1]

			if command == ":open" {
				fmt.Printf("Opening session...")
				response := s.InstantMessagingOpen(message)
				fmt.Printf("\r")
				s.PrintInlineJson(response)

				if response.Ok == true {
					channel = response.Channel.Id
				}
			}
		} else {
			// Send the message to the remote service.
			fmt.Println(message)
		}
	}

	os.Exit(0)
}

func (s *SlackAPI) InstantMessagingClose(channel string) Base {
	var response Base
	s.GetRequest(&response, "im.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) InstantMessagingOpen(userid string) Session {
	var response Session

	if userid == "slackbot" {
		userid = "USLACKBOT"
	}

	s.GetRequest(&response, "im.open", "token", "user="+userid)

	return response
}
