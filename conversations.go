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
