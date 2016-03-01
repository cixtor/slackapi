package main

func (s *SlackAPI) ChannelsHistory(channel string, latest string) {
	s.ResourceHistoryVerbose("channels.history", channel, latest)
}

func (s *SlackAPI) ChannelsId(query string) string {
	response := s.ChannelsList()

	if response.Ok {
		for _, room := range response.Channels {
			if room.Name == query {
				return room.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) ChannelsInfo(channel string) {
	var response interface{}
	channel = s.ChannelsId(channel)
	s.GetRequest(&response, "channels.info", "token", "channel="+channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChannelsList() Rooms {
	if s.TeamRooms.Ok == true {
		return s.TeamRooms
	}

	var response Rooms
	s.GetRequest(&response, "channels.list", "token", "exclude_archived=0")
	s.TeamRooms = response

	return response
}

func (s *SlackAPI) ChannelsListVerbose() {
	response := s.ChannelsList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChannelsMark(channel string, timestamp string) {
	s.ResourceMark("channels.mark", channel, timestamp)
}

func (s *SlackAPI) ChannelsPurgeHistory(channel string, latest string) {
	s.ResourcePurgeHistory("channels.history", channel, latest)
}

func (s *SlackAPI) ChannelsSetPurpose(channel string, purpose string) {
	var response interface{}
	s.GetRequest(&response,
		"channels.setPurpose",
		"token",
		"channel="+channel,
		"purpose="+purpose)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChannelsSetTopic(channel string, topic string) {
	var response interface{}
	s.GetRequest(&response,
		"channels.setTopic",
		"token",
		"channel="+channel,
		"topic="+topic)
	s.PrintAndExit(response)
}
