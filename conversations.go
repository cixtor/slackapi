package slackapi

import (
	"net/url"
	"strconv"
	"strings"
)

type ConversationsAcceptSharedInviteInput struct {
	// Name of the channel. If the channel does not exist already in your
	// workspace, this name is the one that the channel will take.
	ChannelName string `json:"channel_name"`
	// ID of the channel that you'd like to accept.
	// Must provide either invite_id or channel_id.
	ChannelID string `json:"channel_id"`
	// Whether you'd like to use your workspace's free trial for Slack Connect.
	FreeTrialAccepted bool `json:"free_trial_accepted"`
	// See the shared_channel_invite_received event payload for more details
	// on how to retrieve the ID of the invitation.
	InviteID string `json:"invite_id"`
	// Whether the channel should be private.
	IsPrivate bool `json:"is_private"`
	// The ID of the workspace to accept the channel in. If an org-level token
	// is used to call this method, the team_id argument is required.
	TeamID string `json:"team_id"`
}

// ConversationsAcceptSharedInvite is https://api.slack.com/methods/conversations.acceptSharedInvite
func (s *SlackAPI) ConversationsAcceptSharedInvite(input ConversationsAcceptSharedInviteInput) Response {
	var out Response
	if err := s.basePOST("/api/conversations.acceptSharedInvite", input, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// ConversationsApproveSharedInvite is https://api.slack.com/methods/conversations.approveSharedInvite
func (s *SlackAPI) ConversationsApproveSharedInvite(invite_id string, target_team string) Response {
	in := struct {
		InviteID   string `json:"invite_id"`
		TargetTeam string `json:"target_team"`
	}{
		InviteID:   invite_id,
		TargetTeam: target_team,
	}
	var out Response
	if err := s.basePOST("/api/conversations.approveSharedInvite", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

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
		Channel string `json:"channel"`
	}{
		Channel: channel,
	}
	var out ResponseConversationsClose
	if err := s.basePOST("/api/conversations.close", in, &out); err != nil {
		return ResponseConversationsClose{Response: Response{Error: err.Error()}}
	}
	return out
}

type ConversationsCreateInput struct {
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

type ConversationsDeclineSharedInviteInput struct {
	// ID of the Slack Connect invite to decline. Subscribe to the
	// shared_channel_invite_accepted event to receive IDs of Slack Connect
	// channel invites that have been accepted and are awaiting approval.
	InviteID string `json:"invite_id"`
	// The team or enterprise id of the other party involved in the invitation
	// you are declining
	TargetTeam string `json:"target_team"`
}

// ConversationsDeclineSharedInvite is https://api.slack.com/methods/conversations.declineSharedInvite
func (s *SlackAPI) ConversationsDeclineSharedInvite(input ConversationsDeclineSharedInviteInput) Response {
	in := url.Values{}
	if input.InviteID != "" {
		in.Add("invite_id", input.InviteID)
	}
	if input.TargetTeam != "" {
		in.Add("target_team", input.TargetTeam)
	}
	var out Response
	if err := s.baseGET("/api/conversations.declineSharedInvite", in, &out); err != nil {
		return Response{Error: err.Error()}
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

type ConversationsInviteSharedInput struct {
	// ID of the channel on your team that you'd like to share
	Channel string `json:"channel"`
	// Optional email to receive this invite.
	// Either emails or user_ids must be provided.
	Emails []string `json:"emails"`
	// Optional boolean on whether invite is to a external limited member.
	// Defaults to true.
	ExternalLimited bool `json:"external_limited"`
	// Optional user_id to receive this invite. Either emails or user_ids
	// must be provided.
	UserIDs []string `json:"user_ids"`
}

type ConversationsInviteSharedResponse struct {
	Response
	InviteID              string `json:"invite_id"`
	IsLegacySharedChannel bool   `json:"is_legacy_shared_channel"`
}

// ConversationsInviteShared is https://api.slack.com/methods/conversations.inviteShared
func (s *SlackAPI) ConversationsInviteShared(input ConversationsInviteSharedInput) ConversationsInviteSharedResponse {
	in := url.Values{}

	if input.Channel != "" {
		in.Add("channel", input.Channel)
	}

	if len(input.Emails) > 0 {
		in.Add("emails", strings.Join(input.Emails, ","))
	}

	if input.ExternalLimited {
		in.Add("external_limited", "true")
	} else {
		in.Add("external_limited", "false")
	}

	if len(input.UserIDs) > 0 {
		in.Add("user_ids", strings.Join(input.UserIDs, ","))
	}

	var out ConversationsInviteSharedResponse
	if err := s.baseGET("/api/conversations.inviteShared", in, &out); err != nil {
		return ConversationsInviteSharedResponse{Response: Response{Error: err.Error()}}
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

type ConversationsListConnectInvitesInput struct {
	// Maximum number of invites to return.
	Count int `json:"count"`
	// Set to next_cursor returned by previous call to list items in subsequent page.
	Cursor string `json:"cursor"`
	// Encoded team id for the workspace to retrieve invites for.
	TeamID string `json:"team_id"`
}

// ConversationsListConnectInvitesResponse is https://api.slack.com/methods/conversations.listConnectInvites#markdown
type ConversationsListConnectInvitesResponse struct {
	Response
	Invites []interface{} `json:"invites"`
}

// ConversationsListConnectInvites retrieve members of a conversation.
func (s *SlackAPI) ConversationsListConnectInvites(input ConversationsListConnectInvitesInput) ConversationsListConnectInvitesResponse {
	var out ConversationsListConnectInvitesResponse
	if err := s.basePOST("/api/conversations.listConnectInvites", input, &out); err != nil {
		return ConversationsListConnectInvitesResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type ConversationsMembersInput struct {
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
	var out ResponseConversationsMembers
	if err := s.basePOST("/api/conversations.members", input, &out); err != nil {
		return ResponseConversationsMembers{Response: Response{Error: err.Error()}}
	}
	return out
}

type ConversationsOpenInput struct {
	// Resume a conversation by supplying an im or mpim's ID.
	Channel string `json:"channel"`
	// Do not create a direct message or multi-person direct message.
	// This is used to see if there is an existing dm or mpdm.
	PreventCreation bool `json:"prevent_creation"`
	// Indicates you want the full IM channel definition in the response.
	ReturnIm bool `json:"return_im"`
	// Comma separated lists of users. If only one user is included, this
	// creates a 1:1 DM. The ordering of the users is preserved whenever a
	// multi-person direct message is returned. Supply a channel when not
	// supplying users.
	Users string `json:"users"`
}

type ConversationsOpenResponse struct {
	Response
	Channel
	NoOp        bool `json:"no_op"`
	AlreadyOpen bool `json:"already_open"`
}

// ConversationsOpen is https://api.slack.com/methods/conversations.open
func (s *SlackAPI) ConversationsOpen(input ConversationsOpenInput) ConversationsOpenResponse {
	var out ConversationsOpenResponse
	if err := s.basePOST("/api/conversations.open", input, &out); err != nil {
		return ConversationsOpenResponse{Response: Response{Error: err.Error()}}
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

// ChannelPurposeNow defines the expected data from the JSON-encoded API response.
type ChannelPurposeNow struct {
	Response
	Purpose string `json:"purpose"`
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

// ChannelTopicNow defines the expected data from the JSON-encoded API response.
type ChannelTopicNow struct {
	Response
	Topic string `json:"topic"`
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
