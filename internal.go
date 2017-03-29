package main

type ResponseEventlogs struct {
	Response
	Events  []ResponseEvent `json:"events"`
	HasMore bool            `json:"has_more"`
	Total   int             `json:"total"`
}

type ResponseEvent struct {
	Type     string  `json:"type"`
	Channel  string  `json:"channel"`
	FileID   string  `json:"file_id"`
	UserID   string  `json:"user_id"`
	User     string  `json:"user"`
	ItemUser string  `json:"item_user"`
	Subtype  string  `json:"subtype"`
	Hidden   bool    `json:"hidden"`
	IsMpim   bool    `json:"is_mpim"`
	Message  Message `json:"message,omitempty"`
	File     File    `json:"file,omitempty"`
	Reaction string  `json:"reaction"`
	EventTS  string  `json:"event_ts"`
	Latest   string  `json:"latest"`
	Ts       string  `json:"ts"`

	Item ResponseEventItem `json:"item"`
}

type ResponseEventItem struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

type ResponseIssues struct {
	Response
	Issues []string `json:"issues"`
}

func (s *SlackAPI) EventlogHistory(timestamp string) ResponseEventlogs {
	var response ResponseEventlogs
	s.GetRequest(&response, "eventlog.history", "token", "start="+timestamp)
	return response
}

func (s *SlackAPI) HelpIssuesList() ResponseIssues {
	var response ResponseIssues
	s.GetRequest(&response, "help.issues.list", "token")
	return response
}