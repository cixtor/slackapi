package main

func (s *SlackAPI) MultiPartyInstantMessagingList() {
	var response interface{}
	s.GetRequest(&response, "mpim.list", "token")
	s.PrintJson(response)
}
