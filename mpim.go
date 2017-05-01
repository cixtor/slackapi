package main

// ResponseMultiPartyInstantMessageList defines the JSON-encoded output for MultiPartyInstantMessageList.
type ResponseMultiPartyInstantMessageList struct {
	Response
	Groups []Channel `json:"groups"`
}

// ResponseMultiPartyInstantMessageListSimple defines the JSON-encoded output for MultiPartyInstantMessageListSimple.
type ResponseMultiPartyInstantMessageListSimple map[string]string

// MultiPartyInstantMessageHistory fetches history of messages and events from a multiparty direct message.
func (s *SlackAPI) MultiPartyInstantMessageHistory(channel string, latest string) History {
	return s.ResourceHistory("mpim.history", channel, latest)
}

// MultiPartyInstantMessageList lists multiparty direct message channels for the calling user.
func (s *SlackAPI) MultiPartyInstantMessageList() ResponseMultiPartyInstantMessageList {
	var response ResponseMultiPartyInstantMessageList
	s.GetRequest(&response, "mpim.list", "token")
	return response
}

// MultiPartyInstantMessageListSimple lists ID and members in a multiparty direct message channels.
func (s *SlackAPI) MultiPartyInstantMessageListSimple() ResponseMultiPartyInstantMessageListSimple {
	var response ResponseMultiPartyInstantMessageList
	output := make(map[string]string)
	s.GetRequest(&response, "mpim.list", "token")
	for _, data := range response.Groups {
		output[data.ID] = data.Purpose.Value
	}
	return output
}

// MultiPartyInstantMessageMyHistory displays messages of the current user from multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessageMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("mpim.history", channel, latest)
}

// MultiPartyInstantMessagePurgeHistory deletes history of messages and events from multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessagePurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("mpim.history", channel, latest, verbose)
}
