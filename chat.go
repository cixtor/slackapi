package slackapi

import (
	"net/url"
	"strconv"
)

type DeleteMessageInput struct {
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
}

type DeleteMessageResponse struct {
	Response
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
}

// ChatDelete is https://api.slack.com/methods/chat.delete
func (s *SlackAPI) ChatDelete(input DeleteMessageInput) DeleteMessageResponse {
	var out DeleteMessageResponse
	if err := s.basePOST("/api/chat.delete", input, &out); err != nil {
		return DeleteMessageResponse{Response: Response{Error: err.Error()}}
	}
	return out
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

type MeMessageInput struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type MeMessageResponse struct {
	Response
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
}

// ChatMeMessage is https://api.slack.com/methods/chat.meMessage
func (s *SlackAPI) ChatMeMessage(input MeMessageInput) MeMessageResponse {
	var out MeMessageResponse
	if err := s.basePOST("/api/chat.meMessage", input, &out); err != nil {
		return MeMessageResponse{Response: Response{Error: err.Error()}}
	}
	return out
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
