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

// ChatPostMessage sends a message to a channel.
func (s *SlackAPI) ChatPostMessage(data MessageArgs) Post {
	var response Post
	s.postRequest(&response, "chat.postMessage", data)
	return response
}

// ChatUpdate updates a message.
func (s *SlackAPI) ChatUpdate(data MessageArgs) Post {
	var response Post
	s.postRequest(&response, "chat.update", data)
	return response
}
