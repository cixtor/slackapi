package slackapi

import (
	"net/url"
)

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

type ConversationsHistoryInput struct {
	// Conversation ID to fetch history for.
	Channel string `json:"channel"`
	// Paginate through collections of data by setting the cursor parameter to
	// a next_cursor attribute returned by a previous request's response_metadata.
	// Default value fetches the first "page" of the collection. See pagination
	// for more detail.
	Cursor string `json:"cursor"`
	// Include messages with latest or oldest timestamp in results only when
	// either timestamp is specified.
	Inclusive bool `json:"inclusive"`
	// End of time range of messages to include in results.
	Latest string `json:"latest"`
	// The maximum number of items to return. Fewer than the requested number
	// of items may be returned, even if the end of the users list hasn't been
	// reached.
	Limit int `json:"limit"`
	// Start of time range of messages to include in results.
	Oldest string `json:"oldest"`
}

// ConversationsHistory fetches a conversation's history of messages and events.
func (s *SlackAPI) ConversationsHistory(input ConversationsHistoryInput) History {
	var out History
	if err := s.basePOST("/api/conversations.history", input, &out); err != nil {
		return History{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsInfo retrieve information about a conversation.
func (s *SlackAPI) ConversationsInfo(channel string) ResponseChannelsInfo {
	in := url.Values{
		"channel":             []string{channel},
		"include_locale":      []string{"true"},
		"include_num_members": []string{"true"},
	}
	var out ResponseChannelsInfo
	if err := s.baseGET("/api/conversations.info", in, &out); err != nil {
		return ResponseChannelsInfo{Response: Response{Error: err.Error()}}
	}
	return out
}
