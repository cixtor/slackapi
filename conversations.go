package slackapi

import (
	"net/url"
	"strconv"
	"strings"
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

type ResponseConversationsClose struct {
	Response
	NoOp          bool `json:"no_op,omitempty"`
	AlreadyClosed bool `json:"already_closed,omitempty"`
}

// ConversationsClose closes a direct message or multi-person direct message.
// Ref: https://api.slack.com/methods/conversations.close
func (s *SlackAPI) ConversationsClose(channel string) ResponseConversationsClose {
	in := struct {
		Token   string `json:"token"`
		Channel string `json:"channel"`
	}{
		Token:   s.token,
		Channel: channel,
	}
	var out ResponseConversationsClose
	if err := s.basePOST("/api/conversations.close", in, &out); err != nil {
		return ResponseConversationsClose{Response: Response{Error: err.Error()}}
	}
	return out
}

type ConversationsCreateInput struct {
	// Authentication token bearing required scopes. Tokens should be passed
	// as an HTTP Authorization header or alternatively, as a POST parameter.
	Token string `json:"token"`
	// Name of the public or private channel to create.
	Name string `json:"name"`
	// Create a private channel instead of a public one.
	IsPrivate bool `json:"is_private"`
	// Encoded team id to create the channel in, required if org token is used.
	TeamID string `json:"team_id"`
}

type ResponseChannelsInfo struct {
	Response
	Channel Channel `json:"channel"`
}

// ConversationsCreate creates a channel.
func (s *SlackAPI) ConversationsCreate(input ConversationsCreateInput) ResponseChannelsInfo {
	var out ResponseChannelsInfo
	if err := s.basePOST("/api/conversations.create", input, &out); err != nil {
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

// ResponseChannelsList defines the JSON-encoded output for ChannelsList.
type ResponseChannelsList struct {
	Response
	Channels []Channel `json:"channels"`
}

// ConversationsGenericInfo retrieve information about a generic conversation.
func (s *SlackAPI) ConversationsGenericInfo(channels string) ResponseChannelsList {
	in := url.Values{
		"channels": []string{channels},
	}
	var out ResponseChannelsList
	if err := s.baseGET("/api/conversations.genericInfo", in, &out); err != nil {
		return ResponseChannelsList{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsInvite invites users to a channel.
func (s *SlackAPI) ConversationsInvite(channel string, users ...string) ResponseChannelsInfo {
	in := struct {
		Channel string `json:"channel"`
		Users   string `json:"users"`
	}{
		Channel: channel,
		Users:   strings.Join(users, ","),
	}
	var out ResponseChannelsInfo
	if err := s.basePOST("/api/conversations.invite", in, &out); err != nil {
		return ResponseChannelsInfo{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsJoin joins an existing conversation.
func (s *SlackAPI) ConversationsJoin(channel string) ResponseChannelsInfo {
	in := struct {
		Channel string `json:"channel"`
	}{
		Channel: channel,
	}
	var out ResponseChannelsInfo
	if err := s.basePOST("/api/conversations.join", in, &out); err != nil {
		return ResponseChannelsInfo{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsKick removes a user from a conversation.
func (s *SlackAPI) ConversationsKick(channel string, user string) Response {
	in := struct {
		Channel string `json:"channel"`
		User    string `json:"user"`
	}{
		Channel: channel,
		User:    user,
	}
	var out Response
	if err := s.basePOST("/api/conversations.kick", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// ConversationsLeave leaves a conversation.
func (s *SlackAPI) ConversationsLeave(channel string) Response {
	in := struct {
		Channel string `json:"channel"`
	}{
		Channel: channel,
	}
	var out Response
	if err := s.basePOST("/api/conversations.leave", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

type ConversationsMarkInput struct {
	// Authentication token bearing required scopes. Tokens should be passed
	// as an HTTP Authorization header or alternatively, as a POST parameter.
	Token string `json:"token"`
	// Conversation ID to fetch thread from.
	Channel string `json:"channel"`
	// Unique identifier of a thread's parent message.
	Timestamp string `json:"ts"`
}

// ConversationsMark sets the read cursor in a channel.
func (s *SlackAPI) ConversationsMark(input ConversationsMarkInput) Response {
	var out Response
	if err := s.basePOST("/api/conversations.mark", input, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

type ConversationsListInput struct {
	// Paginate through collections of data by setting the cursor parameter to
	// a next_cursor attribute returned by a previous request's response_metadata.
	// Default value fetches the first "page" of the collection. See pagination
	// for more detail.
	Cursor string `json:"cursor"`
	// Set to true to exclude archived channels from the list
	ExcludeArchived bool `json:"exclude_archived"`
	// The maximum number of items to return. Fewer than the requested number
	// of items may be returned, even if the end of the list hasn't been reached.
	// Must be an integer no larger than 1000.
	Limit int `json:"limit"`
	// Mix and match channel types by providing a comma-separated list of any
	// combination of public_channel, private_channel, mpim, im.
	Types []string `json:"types"`
}

// ConversationsList lists all channels in a Slack team.
func (s *SlackAPI) ConversationsList(input ConversationsListInput) ResponseChannelsList {
	in := url.Values{}

	if input.Cursor != "" {
		in.Add("cursor", input.Cursor)
	}

	if input.ExcludeArchived {
		in.Add("exclude_archived", "true")
	}

	if input.Limit > 0 {
		in.Add("limit", strconv.Itoa(input.Limit))
	}

	if len(input.Types) > 0 {
		in.Add("types", strings.Join(input.Types, ","))
	}

	var out ResponseChannelsList
	if err := s.baseGET("/api/conversations.list", in, &out); err != nil {
		return ResponseChannelsList{Response: Response{Error: err.Error()}}
	}
	return out
}

type ConversationsMembersInput struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Cursor  string `json:"cursor"`
	Limit   int    `json:"limit"`
}

type ResponseConversationsMembers struct {
	Response
	Members          []string         `json:"members"`
	ResponseMetadata ResponseMetadata `json:"response_metadata"`
}

// ConversationsMembers retrieve members of a conversation.
func (s *SlackAPI) ConversationsMembers(input ConversationsMembersInput) ResponseConversationsMembers {
	if input.Token == "" {
		input.Token = s.token
	}
	var out ResponseConversationsMembers
	if err := s.basePOST("/api/conversations.members", input, &out); err != nil {
		return ResponseConversationsMembers{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsRename renames a conversation.
func (s *SlackAPI) ConversationsRename(channel string, name string) ResponseChannelsInfo {
	in := struct {
		Channel string `json:"channel"`
		Name    string `json:"name"`
	}{
		Channel: channel,
		Name:    name,
	}
	var out ResponseChannelsInfo
	if err := s.basePOST("/api/conversations.rename", in, &out); err != nil {
		return ResponseChannelsInfo{Response: Response{Error: err.Error()}}
	}
	return out
}

type ConversationsRepliesInput struct {
	// Conversation ID to fetch thread from.
	Channel string `json:"channel"`
	// Unique identifier of a thread's parent message.
	Timestamp string `json:"ts"`
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

// ConversationsReplies lists all channels in a Slack team.
func (s *SlackAPI) ConversationsReplies(input ConversationsRepliesInput) History {
	in := url.Values{}

	in.Add("channel", input.Channel)
	in.Add("ts", input.Timestamp)

	if input.Cursor != "" {
		in.Add("cursor", input.Cursor)
	}

	if input.Inclusive {
		in.Add("inclusive", "true")
	}
	in.Add("latest", input.Latest)

	if input.Limit > 0 {
		in.Add("limit", strconv.Itoa(input.Limit))
	}

	in.Add("oldest", input.Oldest)

	var out History
	if err := s.baseGET("/api/conversations.replies", in, &out); err != nil {
		return History{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsSetPurpose sets the purpose for a conversation.
func (s *SlackAPI) ConversationsSetPurpose(channel string, purpose string) ChannelPurposeNow {
	in := struct {
		Channel string `json:"channel"`
		Purpose string `json:"purpose"`
	}{
		Channel: channel,
		Purpose: purpose,
	}
	var out ChannelPurposeNow
	if err := s.basePOST("/api/conversations.setPurpose", in, &out); err != nil {
		return ChannelPurposeNow{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsSetTopic sets the topic for a conversation.
func (s *SlackAPI) ConversationsSetTopic(channel string, topic string) ChannelTopicNow {
	in := struct {
		Channel string `json:"channel"`
		Topic   string `json:"topic"`
	}{
		Channel: channel,
		Topic:   topic,
	}
	var out ChannelTopicNow
	if err := s.basePOST("/api/conversations.setTopic", in, &out); err != nil {
		return ChannelTopicNow{Response: Response{Error: err.Error()}}
	}
	return out
}

// ConversationsUnarchive reverses conversation archival.
func (s *SlackAPI) ConversationsUnarchive(channel string) Response {
	in := struct {
		Channel string `json:"channel"`
	}{
		Channel: channel,
	}
	var out Response
	if err := s.basePOST("/api/conversations.unarchive", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
