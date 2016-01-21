package main

func (s *SlackAPI) ChannelsHistory(channel string, latest string) {
	var response interface{}

	if latest == "" {
		s.GetRequest(&response,
			"channels.history",
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"unreads=1")
	} else {
		s.GetRequest(&response,
			"channels.history",
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"latest="+latest,
			"unreads=1")
	}
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsInfo(channel string) {
	var response interface{}
	s.GetRequest(&response, "channels.info", "token", "channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsList() {
	var response interface{}
	s.GetRequest(&response, "channels.list", "token", "exclude_archived=0")
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsMark(channel string, timestamp string) {
	var response interface{}
	s.GetRequest(&response,
		"channels.mark",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsSetPurpose(channel string, purpose string) {
	var response interface{}
	s.GetRequest(&response,
		"channels.setPurpose",
		"token",
		"channel="+channel,
		"purpose="+purpose)
	s.PrintJson(response)
}

func (s *SlackAPI) ChannelsSetTopic(channel string, topic string) {
	var response interface{}
	s.GetRequest(&response,
		"channels.setTopic",
		"token",
		"channel="+channel,
		"topic="+topic)
	s.PrintJson(response)
}
