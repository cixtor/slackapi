package main

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
