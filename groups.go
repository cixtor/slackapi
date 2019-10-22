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
	return query
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
