package slackapi

// ResponseGroupsInfo defines the JSON-encoded output for GroupsInfo.
type ResponseGroupsInfo struct {
	Response
	Group Channel `json:"group"`
}

// ResponseGroupsList defines the JSON-encoded output for GroupsList.
type ResponseGroupsList struct {
	Response
	Groups []Channel `json:"groups"`
}

// GroupsClose closes a private channel.
func (s *SlackAPI) GroupsClose(channel string) Response {
	var response Response
	s.getRequest(&response, "groups.close", struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}

// GroupsID gets private channel identifier from readable name.
func (s *SlackAPI) GroupsID(query string) string {
	response := s.GroupsList()

	if response.Ok {
		for _, room := range response.Groups {
			if room.Name == query {
				return room.ID
			}
		}
	}

	return query
}

// GroupsInvite invites a user to a private channel.
func (s *SlackAPI) GroupsInvite(channel string, user string) Response {
	return s.ResourceInvite("groups.invite", s.GroupsID(channel), user)
}

// GroupsKick removes a user from a private channel.
func (s *SlackAPI) GroupsKick(channel string, user string) Response {
	return s.ResourceKick("groups.kick", s.GroupsID(channel), user)
}

// GroupsLeave leaves a private channel.
func (s *SlackAPI) GroupsLeave(channel string) Response {
	return s.ResourceLeave("groups.leave", s.GroupsID(channel))
}

// GroupsList lists private channels that the calling user has access to.
func (s *SlackAPI) GroupsList() ResponseGroupsList {
	if s.teamGroups.Ok {
		return s.teamGroups
	}

	var response ResponseGroupsList
	s.getRequest(&response, "groups.list", struct {
		ExcludeArchived bool `json:"exclude_archived"`
	}{false})
	s.teamGroups = response

	return response
}

// GroupsMark sets the read cursor in a private channel.
func (s *SlackAPI) GroupsMark(channel string, timestamp string) Response {
	return s.ResourceMark("groups.mark", channel, timestamp)
}

// GroupsMyHistory displays messages of the current user from a private channel.
func (s *SlackAPI) GroupsMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("groups.history", channel, latest)
}

// GroupsOpen opens a private channel.
func (s *SlackAPI) GroupsOpen(channel string) Session {
	var response Session
	channel = s.GroupsID(channel)
	s.getRequest(&response, "groups.open", struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}

// GroupsPurgeHistory deletes history of messages and events from a private channel.
func (s *SlackAPI) GroupsPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("groups.history", channel, latest, verbose)
}

// GroupsRename renames a private channel.
func (s *SlackAPI) GroupsRename(channel string, name string) ChannelRename {
	return s.ResourceRename("groups.rename", s.GroupsID(channel), name)
}

// GroupsSetPurpose sets the purpose for a private channel.
func (s *SlackAPI) GroupsSetPurpose(channel string, purpose string) ChannelPurposeNow {
	return s.ResourceSetPurpose("groups.setPurpose", channel, purpose)
}

// GroupsSetRetention sets the retention time of the messages.
func (s *SlackAPI) GroupsSetRetention(channel string, duration int) Response {
	return s.ResourceSetRetention("groups.setRetention", channel, duration)
}

// GroupsSetTopic sets the topic for a private channel.
func (s *SlackAPI) GroupsSetTopic(channel string, topic string) ChannelTopicNow {
	return s.ResourceSetTopic("groups.setTopic", channel, topic)
}

// GroupsUnarchive unarchives a private channel.
func (s *SlackAPI) GroupsUnarchive(channel string) Response {
	return s.ResourceUnarchive("groups.unarchive", s.GroupsID(channel))
}
