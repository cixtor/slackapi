package slackapi

import (
	"encoding/json"
	"log"
	"os"
)

// ChatDelete deletes a message.
func (s *SlackAPI) ChatDelete(channel string, timestamp string) ModifiedMessage {
	var response ModifiedMessage
	s.PostRequest(&response,
		"chat.delete",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

// ChatMeMessage share a me message into a channel.
func (s *SlackAPI) ChatMeMessage(channel string, message string) ModifiedMessage {
	var response ModifiedMessage
	s.PostRequest(&response,
		"chat.meMessage",
		"token",
		"channel="+channel,
		"text="+message)
	return response
}

// ChatPostAttachment sends a message to a channel.
func (s *SlackAPI) ChatPostAttachment(channel string, attachment Attachment) Post {
	return s.SendMessage(map[string]interface{}{
		"channel":     channel,
		"attachments": []Attachment{attachment},
	})
}

// ChatPostMessage sends a message to a channel.
func (s *SlackAPI) ChatPostMessage(channel string, message string) Post {
	return s.SendMessage(map[string]interface{}{
		"channel": channel,
		"text":    message,
	})
}

// ChatReplyMessage sends a message to a channel.
func (s *SlackAPI) ChatReplyMessage(channel string, timestamp string, message string) Post {
	return s.SendMessage(map[string]interface{}{
		"thread_ts": timestamp,
		"channel":   channel,
		"text":      message,
	})
}

// ChatRobotMessage sends a message to a channel as a robot.
func (s *SlackAPI) ChatRobotMessage(channel string, text string) Post {
	s.RobotIsActive = true
	s.RobotName = os.Getenv("SLACK_ROBOT_NAME")
	s.RobotImage = os.Getenv("SLACK_ROBOT_IMAGE")

	if s.RobotName == "" {
		s.RobotName = "foobar"
	}

	if s.RobotImage == "" {
		s.RobotImage = ":slack:"
	}

	data := map[string]interface{}{
		"text":     text,
		"channel":  channel,
		"username": s.RobotName,
	}

	if s.RobotImage[0] == ':' {
		s.RobotImageType = EMOJI
		data[EMOJI] = s.RobotImage
	} else {
		s.RobotImageType = ICONURL
		data[ICONURL] = s.RobotImage
	}

	return s.SendMessage(data)
}

// ChatUpdate updates a message.
func (s *SlackAPI) ChatUpdate(channel string, timestamp string, message string) Post {
	var response Post
	s.PostRequest(&response,
		"chat.update",
		"token",
		"parse=none",
		"channel="+channel,
		"text="+message,
		"ts="+timestamp,
		"link_names=true")
	return response
}

// SendMessage sends a message to a channel.
func (s *SlackAPI) SendMessage(data map[string]interface{}) Post {
	var response Post

	params := []string{
		"token",
		"parse=none",
		"as_user=true",
		"link_names=true",
	}

	for name, value := range data {
		switch value.(type) {
		case string:
			params = append(params, name+"="+value.(string))
		case []Attachment:
			if out, err := json.Marshal(value); err == nil {
				params = append(params, "attachments="+string(out))
			}
		}
	}

	s.PostRequest(&response, "chat.postMessage", params...)

	return response
}
