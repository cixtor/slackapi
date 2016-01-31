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

func (s *SlackAPI) ChatDeleteVerbose(channel string, timestamp string) {
	response := s.ChatDelete(channel, timestamp)
	s.PrintJson(response)
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

func (s *SlackAPI) ChatPostMessageVerbose(channel string, message string) {
	response := s.ChatPostMessage(channel, message)
	s.PrintJson(response)
}

func (s *SlackAPI) ChatSession() {
	var command string
	var parts []string
	var channel string = "unknown"
	var is_connected bool = false
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("slack:%s> ", channel)
		message, err := reader.ReadString('\n')

		if err != nil {
			s.ReportError(err)
		}

		message = strings.TrimSpace(message)

		if message == "" {
			// Do nothing on empty message.
		} else if message == ":exit" {
			// Close the chat session and exit the loop.
			if is_connected {
				fmt.Printf("Closing...")
				response := s.InstantMessagingClose(channel)
				fmt.Printf("\r")
				s.PrintInlineJson(response)
			}
			fmt.Println("Closed")
			break
		} else if message[0] == ':' {
			// Execute a custom command with the message.
			parts = strings.SplitN(message, "\x20", 2)
			command = parts[0]

			if len(parts) == 2 {
				message = parts[1]
			}

			if command == ":open" {
				fmt.Printf("Opening session...")
				response := s.InstantMessagingOpen(message)
				fmt.Printf("\r")
				s.PrintInlineJson(response)

				if response.Ok == true {
					channel = response.Channel.Id
					is_connected = response.Ok
				}
			}
		} else {
			// Send the message to the remote service.
			if is_connected {
				response := s.ChatPostMessage(channel, message)
				s.PrintInlineJson(response)
			} else {
				fmt.Println("{\"ok\":false,\"error\":\"not_connected\"}")
			}
		}
	}

	os.Exit(0)
}

func (s *SlackAPI) ChatUpdate(channel string, timestamp string, message string) Message {
	var response Message
	s.GetRequest(&response,
		"chat.update",
		"token",
		"channel="+channel,
		"text="+message,
		"ts="+timestamp,
		"link_names=1")
	return response
}

func (s *SlackAPI) ChatUpdateVerbose(channel string, timestamp string, message string) {
	response := s.ChatUpdate(channel, timestamp, message)
	s.PrintJson(response)
}
