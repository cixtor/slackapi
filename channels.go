package main

func (s *SlackAPI) ChannelsHistory(channel string, latest string) {
	s.ResourceHistory("channels.history", channel, latest)
}

func (s *SlackAPI) ChannelsInfo(channel string) {
	var response interface{}
	s.GetRequest(&response, "channels.info", "token", "channel="+channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChannelsList() {
	var response interface{}
	s.GetRequest(&response, "channels.list", "token", "exclude_archived=0")
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChannelsMark(channel string, timestamp string) {
	s.ResourceMark("channels.mark", channel, timestamp)
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
