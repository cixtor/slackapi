package main

import (
	"log"
	"os"
)

// ChatDelete deletes a message.
func (s *SlackAPI) ChatDelete(channel string, timestamp string) ModifiedMessage {
	var response ModifiedMessage
	s.GetRequest(&response,
		"chat.delete",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

// ChatMeMessage share a me message into a channel.
func (s *SlackAPI) ChatMeMessage(channel string, message string) ModifiedMessage {
	var response ModifiedMessage
	s.GetRequest(&response,
		"chat.meMessage",
		"token",
		"channel="+channel,
		"text="+message)
	return response
}

// ChatPostMessage sends a message to a channel.
func (s *SlackAPI) ChatPostMessage(channel string, message string) Post {
	var response Post

	if s.RobotIsActive == true {
		var imageType string

		if s.RobotImageType == EMOJI {
			imageType = "icon_emoji"
		} else {
			imageType = ICONURL
		}

		s.GetRequest(&response,
			"chat.postMessage",
			"token",
			"parse=none",
			"channel="+channel,
			"text="+message,
			"as_user=false",
			"link_names=1",
			"username="+s.RobotName,
			imageType+"="+s.RobotImage)
	} else {
		s.GetRequest(&response,
			"chat.postMessage",
			"token",
			"parse=none",
			"channel="+channel,
			"text="+message,
			"as_user=true",
			"link_names=1")
	}

	return response
}

// ChatRobotMessage sends a message to a channel as a robot.
func (s *SlackAPI) ChatRobotMessage(channel string, message string) Post {
	s.RobotIsActive = true
	s.RobotName = os.Getenv("SLACK_ROBOT_NAME")
	s.RobotImage = os.Getenv("SLACK_ROBOT_IMAGE")

	if s.RobotName == "" {
		log.Fatal("Missing SLACK_ROBOT_NAME environment variable")
	}

	if s.RobotImage == "" {
		log.Fatal("Missing SLACK_ROBOT_IMAGE environment variable")
	}

	if s.RobotImage[0] == ':' {
		s.RobotImageType = EMOJI
	} else {
		s.RobotImageType = ICONURL
	}

	return s.ChatPostMessage(channel, message)
}

// ChatUpdate updates a message.
func (s *SlackAPI) ChatUpdate(channel string, timestamp string, message string) Post {
	var response Post
	s.GetRequest(&response,
		"chat.update",
		"token",
		"parse=none",
		"channel="+channel,
		"text="+message,
		"ts="+timestamp,
		"link_names=1")
	return response
}
