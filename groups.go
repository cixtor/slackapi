package main

func (s *SlackAPI) GroupsClose(channel string) {
	var response interface{}
	s.GetRequest(&response, "groups.close", "token", "channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsHistory(channel string, latest string) {
	s.ResourceHistory("groups.history", channel, latest)
}

func (s *SlackAPI) GroupsInfo(channel string) {
	var response interface{}
	s.GetRequest(&response, "groups.info", "token", "channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsList() {
	var response interface{}
	s.GetRequest(&response, "groups.list", "token", "exclude_archived=0")
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsMark(channel string, timestamp string) {
	var response interface{}
	s.GetRequest(&response,
		"groups.mark",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsOpen(channel string) {
	var response interface{}
	s.GetRequest(&response,
		"groups.open",
		"token",
		"channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsSetPurpose(channel string, purpose string) {
	var response interface{}
	s.GetRequest(&response,
		"groups.setPurpose",
		"token",
		"channel="+channel,
		"purpose="+purpose)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsSetTopic(channel string, topic string) {
	var response interface{}
	s.GetRequest(&response,
		"groups.setTopic",
		"token",
		"channel="+channel,
		"topic="+topic)
	s.PrintJson(response)
}
