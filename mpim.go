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
	s.PostRequest(&response, "mpim.close", struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}

// MultiPartyInstantMessageHistory fetches history of messages and events from a multiparty direct message.
func (s *SlackAPI) MultiPartyInstantMessageHistory(channel string, latest string) History {
	return s.ResourceHistory("mpim.history", HistoryArgs{
		Channel: channel,
		Latest:  latest,
	})
}

// MultiPartyInstantMessageList lists multiparty direct message channels for the calling user.
func (s *SlackAPI) MultiPartyInstantMessageList() ResponseMultiPartyInstantMessageList {
	var response ResponseMultiPartyInstantMessageList
	s.GetRequest(&response, "mpim.list", nil)
	return response
}

// MultiPartyInstantMessageListSimple lists ID and members in a multiparty direct message channels.
func (s *SlackAPI) MultiPartyInstantMessageListSimple() ResponseMultiPartyInstantMessageListSimple {
	var response ResponseMultiPartyInstantMessageList
	output := make(map[string]string)
	s.GetRequest(&response, "mpim.list", nil)
	for _, data := range response.Groups {
		output[data.ID] = data.Purpose.Value
	}
	return output
}

// MultiPartyInstantMessageMark sets the read cursor in a multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessageMark(channel string, latest string) Response {
	return s.ResourceMark("mpim.mark", channel, latest)
}

// MultiPartyInstantMessageMyHistory displays messages of the current user from multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessageMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("mpim.history", channel, latest)
}

// MultiPartyInstantMessageOpen this method opens a multiparty direct message.
func (s *SlackAPI) MultiPartyInstantMessageOpen(users []string) ResponseMultiPartyInstantMessageOpen {
	var response ResponseMultiPartyInstantMessageOpen
	s.GetRequest(&response, "mpim.open", struct {
		Users []string `json:"users"`
	}{users})
	return response
}

// MultiPartyInstantMessagePurgeHistory deletes history of messages and events from multiparty direct message channel.
func (s *SlackAPI) MultiPartyInstantMessagePurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("mpim.history", channel, latest, verbose)
}
