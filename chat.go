package main

import (
	"log"
	"os"
)

func (s *SlackAPI) ChatDelete(channel string, timestamp string) ModifiedMessage {
	var response ModifiedMessage
	s.GetRequest(&response,
		"chat.delete",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

func (s *SlackAPI) ChatMeMessage(channel string, message string) ModifiedMessage {
	var response ModifiedMessage
	s.GetRequest(&response,
		"chat.meMessage",
		"token",
		"channel="+channel,
		"text="+message)
	return response
}

func (s *SlackAPI) ChatPostMessage(channel string, message string) Post {
	var response Post

	if s.RobotIsActive == true {
		var imageType string

		if s.RobotImageType == "emoji" {
			imageType = "icon_emoji"
		} else {
			imageType = "icon_url"
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
		s.RobotImageType = "emoji"
	} else {
		s.RobotImageType = "icon_url"
	}

	return s.ChatPostMessage(channel, message)
}

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
