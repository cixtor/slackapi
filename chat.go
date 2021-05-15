package slackapi

// ChatDelete deletes a message.
func (s *SlackAPI) ChatDelete(data MessageArgs) ModifiedMessage {
	var response ModifiedMessage
	s.postRequest(&response, "chat.delete", data)
	return response
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

// ChatUpdate updates a message.
func (s *SlackAPI) ChatUpdate(data MessageArgs) Post {
	var response Post
	s.postRequest(&response, "chat.update", data)
	return response
}
