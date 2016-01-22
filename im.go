package main

func (s *SlackAPI) InstantMessagingClose(channel string) Base {
	var response Base
	s.GetRequest(&response, "im.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) InstantMessagingCloseVerbose(channel string) {
	response := s.InstantMessagingClose(channel)
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingList() {
	var response interface{}
	s.GetRequest(&response, "im.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) InstantMessagingOpen(userid string) Session {
	var response Session

	if userid == "slackbot" {
		userid = "USLACKBOT"
	}

	s.GetRequest(&response, "im.open", "token", "user="+userid)

	return response
}

func (s *SlackAPI) InstantMessagingOpenVerbose(userid string) {
	response := s.InstantMessagingOpen(userid)
	s.PrintJson(response)
}