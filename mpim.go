package slackapi

// ResponseMultiPartyInstantMessageList defines the JSON-encoded output for MultiPartyInstantMessageList.
type ResponseMultiPartyInstantMessageList struct {
	Response
	Groups []Channel `json:"groups"`
}

// ResponseMultiPartyInstantMessageListSimple defines the JSON-encoded output for MultiPartyInstantMessageListSimple.
type ResponseMultiPartyInstantMessageListSimple map[string]string

// ResponseMultiPartyInstantMessageOpen defines the JSON-encoded output for ResponseMultiPartyInstantMessageOpen.
type ResponseMultiPartyInstantMessageOpen struct {
	Response
	Group Channel `json:"group"`
}

// MultiPartyInstantMessageClose closes a multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessageClose(channel string) Response {
	var response Response
	s.postRequest(&response, "mpim.close", struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}

// MultiPartyInstantMessageHistory fetches history of messages and events from a multiparty direct message.
func (s *SlackAPI) MultiPartyInstantMessageHistory(data HistoryArgs) History {
	return s.ResourceHistory("mpim.history", data)
}

// MultiPartyInstantMessageList lists multiparty direct message channels for the calling user.
func (s *SlackAPI) MultiPartyInstantMessageList() ResponseMultiPartyInstantMessageList {
	var response ResponseMultiPartyInstantMessageList
	s.getRequest(&response, "mpim.list", nil)
	return response
}

// MultiPartyInstantMessageListSimple lists ID and members in a multiparty direct message channels.
func (s *SlackAPI) MultiPartyInstantMessageListSimple() ResponseMultiPartyInstantMessageListSimple {
	var response ResponseMultiPartyInstantMessageList
	output := make(map[string]string)
	s.getRequest(&response, "mpim.list", nil)
	for _, data := range response.Groups {
		output[data.ID] = data.Purpose.Value
	}
	return output
}

// MultiPartyInstantMessageMyHistory displays messages of the current user from multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessageMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("mpim.history", channel, latest)
}

// MultiPartyInstantMessageOpen this method opens a multiparty direct message.
func (s *SlackAPI) MultiPartyInstantMessageOpen(users []string) ResponseMultiPartyInstantMessageOpen {
	var response ResponseMultiPartyInstantMessageOpen
	s.getRequest(&response, "mpim.open", struct {
		Users []string `json:"users"`
	}{users})
	return response
}

// MultiPartyInstantMessagePurgeHistory deletes history of messages and events from multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessagePurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("mpim.history", channel, latest, verbose)
}
