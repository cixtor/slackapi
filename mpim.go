package main

type ResponseMPIMList struct {
	Response
	Groups []Channel `json:"groups"`
}

type ResponseMPIMListSimple map[string]string

func (s *SlackAPI) MultiPartyInstantMessagingList() ResponseMPIMList {
	var response ResponseMPIMList
	s.GetRequest(&response, "mpim.list", "token")
	return response
}

func (s *SlackAPI) MultiPartyInstantMessagingListSimple() ResponseMPIMListSimple {
	var response ResponseMPIMList
	output := make(map[string]string)
	s.GetRequest(&response, "mpim.list", "token")
	for _, data := range response.Groups {
		output[data.Id] = data.Purpose.Value
	}
	return output
}
