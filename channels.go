package slackapi

// ResponseChannelsInfo defines the JSON-encoded output for ChannelsInfo.
type ResponseChannelsInfo struct {
	Response
	Channel Channel `json:"channel"`
}

// ResponseChannelsJoin defines the JSON-encoded output for ChannelsJoin.
type ResponseChannelsJoin struct {
	Response
	AlreadyInChannel bool    `json:"already_in_channel"`
	Channel          Channel `json:"channel"`
}

// ResponseChannelsList defines the JSON-encoded output for ChannelsList.
type ResponseChannelsList struct {
	Response
	Channels []Channel `json:"channels"`
}

// ChannelSuggestions defines the expected data from the JSON-encoded API response.
type ChannelSuggestions struct {
	Response
	Status               Response `json:"status"`
	SuggestionTypesTried []string `json:"suggestion_types_tried"`
}

// ChannelsID gets channel identifier from readable name.
func (s *SlackAPI) ChannelsID(query string) string {
	response := s.ChannelsList()

	if response.Ok {
		for _, room := range response.Channels {
			if room.Name == query {
				return room.ID
			}
		}
	}

	return query
}

// ChannelsInvite invites a user to a channel.
func (s *SlackAPI) ChannelsInvite(channel string, user string) Response {
	return s.ResourceInvite("channels.invite", s.ChannelsID(channel), user)
}

// ChannelsJoin joins a channel, creating it if needed.
func (s *SlackAPI) ChannelsJoin(name string) ResponseChannelsJoin {
	var response ResponseChannelsJoin
	s.postRequest(&response, "channels.join", struct {
		Name     string `json:"name"`
		Validate bool   `json:"validate"`
	}{name, true})
	return response
}

// ChannelsKick removes a user from a channel.
func (s *SlackAPI) ChannelsKick(channel string, user string) Response {
	return s.ResourceKick("channels.kick", s.ChannelsID(channel), user)
}

// ChannelsLeave leaves a channel.
func (s *SlackAPI) ChannelsLeave(channel string) Response {
	return s.ResourceLeave("channels.leave", s.ChannelsID(channel))
}

// ChannelsList lists all channels in a Slack team.
func (s *SlackAPI) ChannelsList() ResponseChannelsList {
	if s.teamChannels.Ok {
		return s.teamChannels
	}

	var response ResponseChannelsList
	s.getRequest(&response, "channels.list", struct {
		ExcludeArchived bool `json:"exclude_archived"`
	}{false})
	s.teamChannels = response

	return response
}

// ChannelsMark sets the read cursor in a channel.
func (s *SlackAPI) ChannelsMark(channel string, timestamp string) Response {
	return s.ResourceMark("channels.mark", channel, timestamp)
}

// ChannelsMyHistory displays messages of the current user from a channel.
func (s *SlackAPI) ChannelsMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("channels.history", channel, latest)
}

// ChannelsPurgeHistory deletes history of messages and events from a channel.
func (s *SlackAPI) ChannelsPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("channels.history", channel, latest, verbose)
}

// ChannelsRename renames a channel.
func (s *SlackAPI) ChannelsRename(channel string, name string) ChannelRename {
	return s.ResourceRename("channels.rename", s.ChannelsID(channel), name)
}

// ChannelsSetPurpose sets the purpose for a channel.
func (s *SlackAPI) ChannelsSetPurpose(channel string, purpose string) ChannelPurposeNow {
	return s.ResourceSetPurpose("channels.setPurpose", channel, purpose)
}

// ChannelsSetRetention sets the retention time of the messages.
func (s *SlackAPI) ChannelsSetRetention(channel string, duration int) Response {
	return s.ResourceSetRetention("channels.setRetention", channel, duration)
}

// ChannelsSetTopic sets the topic for a channel.
func (s *SlackAPI) ChannelsSetTopic(channel string, topic string) ChannelTopicNow {
	return s.ResourceSetTopic("channels.setTopic", channel, topic)
}

// ChannelsSuggestions prints a list of suggested channels to join.
func (s *SlackAPI) ChannelsSuggestions() ChannelSuggestions {
	var response ChannelSuggestions
	s.getRequest(&response, "channels.suggestions", nil)
	return response
}

// ChannelsUnarchive unarchives a channel.
func (s *SlackAPI) ChannelsUnarchive(channel string) Response {
	return s.ResourceUnarchive("channels.unarchive", s.ChannelsID(channel))
}
