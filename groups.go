package main

type ResponseGroupsInfo struct {
	Response
	Group Channel `json:"group"`
}

type ResponseGroupsList struct {
	Response
	Groups []Channel `json:"groups"`
}

func (s *SlackAPI) GroupsArchive(channel string) Response {
	var response Response
	s.GetRequest(&response, "groups.archive", "token", "channel="+s.GroupsId(channel))
	return response
}

func (s *SlackAPI) GroupsClose(channel string) Response {
	var response Response
	s.GetRequest(&response, "groups.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsHistory(channel string, latest string) History {
	return s.ResourceHistory("groups.history", channel, latest)
}

func (s *SlackAPI) GroupsId(query string) string {
	response := s.GroupsList()

	if response.Ok {
		for _, room := range response.Groups {
			if room.Name == query {
				return room.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) GroupsInfo(channel string) ResponseGroupsInfo {
	var response ResponseGroupsInfo
	channel = s.GroupsId(channel)
	s.GetRequest(&response, "groups.info", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsList() ResponseGroupsList {
	if s.TeamGroups.Ok == true {
		return s.TeamGroups
	}

	var response ResponseGroupsList
	s.GetRequest(&response, "groups.list", "token", "exclude_archived=0")
	s.TeamGroups = response

	return response
}

func (s *SlackAPI) GroupsMark(channel string, timestamp string) Response {
	return s.ResourceMark("groups.mark", channel, timestamp)
}

func (s *SlackAPI) GroupsMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("groups.history", channel, latest)
}

func (s *SlackAPI) GroupsOpen(channel string) Session {
	var response Session
	channel = s.GroupsId(channel)
	s.GetRequest(&response, "groups.open", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("groups.history", channel, latest, verbose)
}

func (s *SlackAPI) GroupsSetPurpose(channel string, purpose string) ChannelPurposeNow {
	return s.ResourceSetPurpose("groups.setPurpose", channel, purpose)
}

func (s *SlackAPI) GroupsSetRetention(channel string, duration string) Response {
	return s.ResourceSetRetention("groups.setRetention", channel, duration)
}

func (s *SlackAPI) GroupsSetTopic(channel string, topic string) ChannelTopicNow {
	return s.ResourceSetTopic("groups.setTopic", channel, topic)
}
