package main

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

// GroupsArchive archives a private channel.
func (s *SlackAPI) GroupsArchive(channel string) Response {
	return s.ResourceArchive("groups.archive", s.GroupsID(channel))
}

// GroupsClose closes a private channel.
func (s *SlackAPI) GroupsClose(channel string) Response {
	var response Response
	s.GetRequest(&response, "groups.close", "token", "channel="+channel)
	return response
}

// GroupsCreate creates a private channel.
func (s *SlackAPI) GroupsCreate(channel string) ResponseGroupsInfo {
	var response ResponseGroupsInfo
	s.GetRequest(&response, "groups.create", "token", "name="+channel)
	return response
}

// GroupsCreateChild clones and archives a private channel.
func (s *SlackAPI) GroupsCreateChild(channel string) ResponseGroupsInfo {
	var response ResponseGroupsInfo
	s.GetRequest(&response, "groups.createChild", "token", "name="+s.GroupsID(channel))
	return response
}

// GroupsHistory fetches history of messages and events from a private channel.
func (s *SlackAPI) GroupsHistory(channel string, latest string) History {
	return s.ResourceHistory("groups.history", channel, latest)
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

// GroupsInfo gets information about a private channel.
func (s *SlackAPI) GroupsInfo(channel string) ResponseGroupsInfo {
	var response ResponseGroupsInfo
	channel = s.GroupsID(channel)
	s.GetRequest(&response, "groups.info", "token", "channel="+channel)
	return response
}

// GroupsInvite invites a user to a private channel.
func (s *SlackAPI) GroupsInvite(channel string, user string) Response {
	return s.ResourceInvite("groups.invite", s.GroupsID(channel), s.UsersID(user))
}

// GroupsKick removes a user from a private channel.
func (s *SlackAPI) GroupsKick(channel string, user string) Response {
	return s.ResourceKick("groups.kick", s.GroupsID(channel), s.UsersID(user))
}

// GroupsLeave leaves a private channel.
func (s *SlackAPI) GroupsLeave(channel string) Response {
	return s.ResourceLeave("groups.leave", s.GroupsID(channel))
}

// GroupsList lists private channels that the calling user has access to.
func (s *SlackAPI) GroupsList() ResponseGroupsList {
	if s.TeamGroups.Ok == true {
		return s.TeamGroups
	}

	var response ResponseGroupsList
	s.GetRequest(&response, "groups.list", "token", "exclude_archived=0")
	s.TeamGroups = response

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
	s.GetRequest(&response, "groups.open", "token", "channel="+channel)
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
func (s *SlackAPI) GroupsSetRetention(channel string, duration string) Response {
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
