package main

// InstantMessageList defines the expected data from the JSON-encoded API response.
type InstantMessageList struct {
	Response
	InstantMessages []InstantMessageObject `json:"ims"`
}

// InstantMessageObject defines the expected data from the JSON-encoded API response.
type InstantMessageObject struct {
	Created          int    `json:"created"`
	ID               string `json:"id"`
	IsInstantMessage bool   `json:"is_im"`
	IsUserDeleted    bool   `json:"is_user_deleted"`
	User             string `json:"user"`
}

// InstantMessageClose close a direct message channel.
func (s *SlackAPI) InstantMessageClose(channel string) Response {
	var response Response
	s.GetRequest(&response, "im.close", "token", "channel="+channel)
	return response
}

// InstantMessageHistory fetches history of messages and events from direct message channel.
func (s *SlackAPI) InstantMessageHistory(channel string, latest string) History {
	return s.ResourceHistory("im.history", channel, latest)
}

// InstantMessageList lists direct message channels for the calling user.
func (s *SlackAPI) InstantMessageList() InstantMessageList {
	var response InstantMessageList
	s.GetRequest(&response, "im.list", "token")
	return response
}

// InstantMessageMark sets the read cursor in a direct message channel.
func (s *SlackAPI) InstantMessageMark(channel string, timestamp string) Response {
	return s.ResourceMark("im.mark", channel, timestamp)
}

// InstantMessageMyHistory displays messages of the current user from direct message channel.
func (s *SlackAPI) InstantMessageMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("im.history", channel, latest)
}

// InstantMessageOpen opens a direct message channel.
func (s *SlackAPI) InstantMessageOpen(userid string) Session {
	var response Session
	userid = s.UsersID(userid)
	s.GetRequest(&response, "im.open", "token", "user="+userid)
	return response
}

// InstantMessagePurgeHistory deletes history of messages and events from direct message channel.
func (s *SlackAPI) InstantMessagePurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("im.history", channel, latest, verbose)
}
