package slackapi

// ConversationsArchive archives a conversation.
func (s *SlackAPI) ConversationsArchive(channel string) Response {
	in := struct {
		Channel string `json:"channel"`
	}{
		Channel: channel,
	}
	var out Response
	if err := s.basePOST("/api/conversations.archive", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// ConversationsCreate creates a channel.
func (s *SlackAPI) ConversationsCreate(name string) ResponseChannelsInfo {
	in := struct {
		Name     string `json:"name"`
		Validate bool   `json:"validate"`
	}{
		Name:     name,
		Validate: true,
	}
	var out ResponseChannelsInfo
	if err := s.basePOST("/api/channels.create", in, &out); err != nil {
		return ResponseChannelsInfo{Response: Response{Error: err.Error()}}
	}
	return out
}
