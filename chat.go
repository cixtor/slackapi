package slackapi

import (
	"net/url"
	"strconv"
)

// ChatDelete deletes a message.
func (s *SlackAPI) ChatDelete(data MessageArgs) ModifiedMessage {
	var response ModifiedMessage
	s.postRequest(&response, "chat.delete", data)
	return response
}

type ChatDeleteAttachmentInput struct {
	Channel    string `json:"channel"`
	Ts         string `json:"ts"`
	Attachment int    `json:"attachment"`
}

type ChatDeleteAttachmentResponse struct {
	Response
	Channel    string `json:"channel"`
	Ts         string `json:"ts"`
	Attachment string `json:"attachment"`
}

// ChatDeleteAttachment is https://api.slack.com/methods/chat.deleteAttachment
func (s *SlackAPI) ChatDeleteAttachment(input ChatDeleteAttachmentInput) ChatDeleteAttachmentResponse {
	in := url.Values{}
	if input.Channel != "" {
		in.Add("channel", input.Channel)
	}
	if input.Ts != "" {
		in.Add("ts", input.Ts)
	}
	if input.Attachment > 0 {
		in.Add("attachment", strconv.Itoa(input.Attachment))
	}
	var out ChatDeleteAttachmentResponse
	if err := s.baseGET("/api/chat.deleteAttachment", in, &out); err != nil {
		return ChatDeleteAttachmentResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// ChatMeMessage share a me message into a channel.
func (s *SlackAPI) ChatMeMessage(data MessageArgs) ModifiedMessage {
	var response ModifiedMessage
	s.postRequest(&response, "chat.meMessage", data)
	return response
}

// ChatPostMessage is https://api.slack.com/methods/chat.postMessage
func (s *SlackAPI) ChatPostMessage(input MessageArgs) Post {
	var out Post
	if err := s.basePOST("/api/chat.postMessage", input, &out); err != nil {
		return Post{Response: Response{Error: err.Error()}}
	}
	return out
}

// ChatUpdate is https://api.slack.com/methods/chat.update
func (s *SlackAPI) ChatUpdate(input MessageArgs) Post {
	var out Post
	if err := s.basePOST("/api/chat.update", input, &out); err != nil {
		return Post{Response: Response{Error: err.Error()}}
	}
	return out
}
