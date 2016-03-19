package main

type IMList struct {
	Response
	Ims []IMObject `json:"ims"`
}

type IMHistory struct {
	Response
	HasMore            bool             `json:"has_more"`
	Messages           []HistoryMessage `json:"messages"`
	UnreadCountDisplay int              `json:"unread_count_display"`
}

type IMMyHistory struct {
	Filtered int
	Latest   string
	Messages []HistoryMessage
	Oldest   string
	Total    int
	Username string
}

type IMDeletedHistory struct {
	Deleted    int
	NotDeleted int
	Messages   []IMDeletedMessage
}

type IMDeletedMessage struct {
	Deleted bool
	Text    string
	Ts      string
}

type IMObject struct {
	Created       int    `json:"created"`
	Id            string `json:"id"`
	IsIm          bool   `json:"is_im"`
	IsUserDeleted bool   `json:"is_user_deleted"`
	User          string `json:"user"`
}

type HistoryMessage struct {
	Attachments  []Attachment `json:"attachments"`
	BotId        string       `json:"bot_id"`
	DisplayAsBot bool         `json:"display_as_bot"`
	File         File         `json:"file"`
	Subtype      string       `json:"subtype"`
	Text         string       `json:"text"`
	Ts           string       `json:"ts"`
	Type         string       `json:"type"`
	Upload       bool         `json:"upload"`
	User         string       `json:"user"`
	Username     string       `json:"username"`
}

type Attachment struct {
	Fallback    string `json:"fallback"`
	FromUrl     string `json:"from_url"`
	Id          int    `json:"id"`
	ServiceName string `json:"service_name"`
	Text        string `json:"text"`
	ThumbHeight int    `json:"thumb_height"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	Title       string `json:"title"`
	TitleLink   string `json:"title_link"`
}

func (s *SlackAPI) InstantMessagingClose(channel string) Response {
	var response Response
	s.GetRequest(&response, "im.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) InstantMessagingHistory(channel string, latest string) IMHistory {
	var response IMHistory
	s.GetRequest(&response,
		"im.history",
		"token",
		"channel="+channel,
		"latest="+latest,
		"inclusive=1",
		"count=1000",
		"unreads=1")
	return response
}

func (s *SlackAPI) InstantMessagingList() IMList {
	var response IMList
	s.GetRequest(&response, "im.list", "token")
	return response
}

func (s *SlackAPI) InstantMessagingMark(channel string, timestamp string) Response {
	var response Response
	s.GetRequest(&response,
		"im.mark",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

func (s *SlackAPI) InstantMessagingMyHistory(channel string, latest string) IMMyHistory {
	var owner Owner = s.AuthTest()
	var rhistory IMMyHistory

	response := s.InstantMessagingHistory(channel, latest)

	for _, message := range response.Messages {
		rhistory.Total += 1

		if message.User == owner.UserId {
			rhistory.Messages = append(rhistory.Messages, message)
			rhistory.Filtered += 1
		}
	}

	if rhistory.Total > 0 {
		var offset int = len(response.Messages) - 1

		rhistory.Username = owner.User
		rhistory.Latest = response.Messages[0].Ts
		rhistory.Oldest = response.Messages[offset].Ts
	}

	return rhistory
}

func (s *SlackAPI) InstantMessagingOpen(userid string) Session {
	var response Session
	userid = s.UsersId(userid)
	s.GetRequest(&response, "im.open", "token", "user="+userid)
	return response
}

func (s *SlackAPI) InstantMessagingPurgeHistory(channel string, latest string) IMDeletedHistory {
	var delhist IMDeletedHistory
	var delmsg IMDeletedMessage

	response := s.InstantMessagingMyHistory(channel, latest)

	if response.Filtered > 0 {
		for _, message := range response.Messages {
			result := s.ChatDelete(channel, message.Ts)
			delmsg.Text = message.Text
			delmsg.Ts = message.Ts

			if result.Ok == true {
				delhist.Deleted++
				delmsg.Deleted = true
			} else {
				delhist.NotDeleted++
				delmsg.Deleted = false
			}

			delhist.Messages = append(delhist.Messages, delmsg)
		}
	}

	return delhist
}
