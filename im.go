package main

type IMList struct {
	Response
	Ims []IMObject `json:"ims"`
}

type IMObject struct {
	Created       int    `json:"created"`
	Id            string `json:"id"`
	IsIm          bool   `json:"is_im"`
	IsUserDeleted bool   `json:"is_user_deleted"`
	User          string `json:"user"`
}

func (s *SlackAPI) InstantMessagingClose(channel string) Response {
	var response Response
	s.GetRequest(&response, "im.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) InstantMessagingHistory(channel string, latest string) History {
	return s.ResourceHistory("im.history", channel, latest)
}

func (s *SlackAPI) InstantMessagingList() IMList {
	var response IMList
	s.GetRequest(&response, "im.list", "token")
	return response
}

func (s *SlackAPI) InstantMessagingMark(channel string, timestamp string) Response {
	return s.ResourceMark("im.mark", channel, timestamp)
}

func (s *SlackAPI) InstantMessagingMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("im.history", channel, latest)
}

func (s *SlackAPI) InstantMessagingOpen(userid string) Session {
	var response Session
	userid = s.UsersId(userid)
	s.GetRequest(&response, "im.open", "token", "user="+userid)
	return response
}

func (s *SlackAPI) InstantMessagingPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("im.history", channel, latest, verbose)
}
