package slackapi

// ResponseEventlogs defines the JSON-encoded output for Eventlogs.
type ResponseEventlogs struct {
	Response
	Events  []ResponseEvent `json:"events"`
	HasMore bool            `json:"has_more"`
	Total   int             `json:"total"`
}

// ResponseEvent defines the JSON-encoded output for Event.
type ResponseEvent struct {
	Type      string  `json:"type"`
	Channel   string  `json:"channel"`
	FileID    string  `json:"file_id"`
	UserID    string  `json:"user_id"`
	User      string  `json:"user"`
	ItemUser  string  `json:"item_user"`
	Subtype   string  `json:"subtype"`
	Hidden    bool    `json:"hidden"`
	IsMpim    bool    `json:"is_mpim"`
	Message   Message `json:"message,omitempty"`
	File      File    `json:"file,omitempty"`
	Reaction  string  `json:"reaction"`
	EventTS   string  `json:"event_ts"`
	Latest    string  `json:"latest"`
	Timestamp string  `json:"ts"`

	Item ResponseEventItem `json:"item"`
}

// ResponseEventItem defines the JSON-encoded output for EventItem.
type ResponseEventItem struct {
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
}

// EventlogHistory lists all the events since the specified time.
func (s *SlackAPI) EventlogHistory(start string) ResponseEventlogs {
	var response ResponseEventlogs
	s.getRequest(&response, "eventlog.history", struct {
		Start string `json:"start"`
	}{start})
	return response
}
