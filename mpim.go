package main

type ResponseMPIMList struct {
	Response
	Groups []Channel `json:"groups"`
}

func (s *SlackAPI) MultiPartyInstantMessagingList() ResponseMPIMList {
	var response ResponseMPIMList
	s.GetRequest(&response, "mpim.list", "token")
	return response
}
