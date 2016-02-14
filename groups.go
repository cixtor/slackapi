package main

func (s *SlackAPI) GroupsClose(channel string) {
	var response interface{}
	s.GetRequest(&response, "groups.close", "token", "channel="+channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) GroupsHistory(channel string, latest string) {
	s.ResourceHistory("groups.history", channel, latest)
}

func (s *SlackAPI) GroupsId(query string) string {
	response := s.GroupsList()

	if response.Ok {
		for _, room := range response.Channels {
			if room.Name == query {
				return room.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) GroupsInfo(channel string) {
	var response interface{}
	channel = s.GroupsId(channel)
	s.GetRequest(&response, "groups.info", "token", "channel="+channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) GroupsList() Groups {
	if s.TeamGroups.Ok == true {
		return s.TeamGroups
	}

	var response Groups
	s.GetRequest(&response, "groups.list", "token", "exclude_archived=0")
	s.TeamGroups = response

	return response
}

func (s *SlackAPI) GroupsListVerbose() {
	response := s.GroupsList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) GroupsMark(channel string, timestamp string) {
	s.ResourceMark("groups.mark", channel, timestamp)
}

func (s *SlackAPI) GroupsOpen(channel string) Session {
	var response Session
	s.GetRequest(&response, "groups.open", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsOpenVerbose(channel string) {
	response := s.GroupsOpen(channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) GroupsSetPurpose(channel string, purpose string) {
	var response interface{}
	s.GetRequest(&response,
		"groups.setPurpose",
		"token",
		"channel="+channel,
		"purpose="+purpose)
	s.PrintAndExit(response)
}

func (s *SlackAPI) GroupsSetTopic(channel string, topic string) {
	var response interface{}
	s.GetRequest(&response,
		"groups.setTopic",
		"token",
		"channel="+channel,
		"topic="+topic)
	s.PrintAndExit(response)
}
