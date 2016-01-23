package main

func (s *SlackAPI) GroupsClose(channel string) {
	var response interface{}
	s.GetRequest(&response, "groups.close", "token", "channel="+channel)
	s.PrintJson(response)
}

func (s *SlackAPI) GroupsHistory(channel string, latest string) {
	var response interface{}

	if latest == "" {
		s.GetRequest(&response,
			"groups.history",
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"unreads=1")
	} else {
		s.GetRequest(&response,
			"groups.history",
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"latest="+latest,
			"unreads=1")
	}
	s.PrintJson(response)
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
