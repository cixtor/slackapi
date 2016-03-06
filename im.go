package main

func (s *SlackAPI) InstantMessagingClose(channel string) Base {
	var response Base
	s.GetRequest(&response, "im.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) InstantMessagingCloseVerbose(channel string) {
	response := s.InstantMessagingClose(channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) InstantMessagingHistory(channel string, latest string) {
	s.ResourceHistoryVerbose("im.history", channel, latest)
}

func (s *SlackAPI) InstantMessagingList() {
	var response interface{}
	s.GetRequest(&response, "im.list", "token")
	s.PrintAndExit(response)
}

func (s *SlackAPI) InstantMessagingMark(channel string, timestamp string) {
	s.ResourceMark("im.mark", channel, timestamp)
}

func (s *SlackAPI) InstantMessagingMyHistory(channel string, latest string) {
	s.ResourceMyHistoryVerbose("im.history", channel, latest)
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
	s.PrintAndExit(response)
}

func (s *SlackAPI) InstantMessagingPurgeHistory(channel string, latest string) {
	s.ResourcePurgeHistory("im.history", channel, latest)
}
