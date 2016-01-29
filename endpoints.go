package main

func (s *SlackAPI) ApiTest() {
	var response interface{}
	s.GetRequest(&response, "api.test")
	s.PrintJson(response)
}

func (s *SlackAPI) AuthTest() {
	var response interface{}
	s.GetRequest(&response, "auth.test", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) EmojiList() {
	var response interface{}
	s.GetRequest(&response, "emoji.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) ResourceHistory(action string, channel string, latest string) {
	var response interface{}

	if latest == "" {
		s.GetRequest(&response,
			action,
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"unreads=1")
	} else {
		s.GetRequest(&response,
			action,
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"latest="+latest,
			"unreads=1")
	}

	s.PrintJson(response)
}

func (s *SlackAPI) ResourceMark(action string, channel string, timestamp string) {
	var response interface{}
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"ts="+timestamp)
	s.PrintJson(response)
}

func (s *SlackAPI) TeamInfo() {
	var response interface{}
	s.GetRequest(&response, "team.info", "token")
	s.PrintJson(response)
}
