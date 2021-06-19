package slackapi

type ClientCountsInput struct {
	OrgWideAware          bool `json:"org_wide_aware"`
	ThreadCountsByChannel bool `json:"thread_counts_by_channel"`
}

type ClientCountsResponse struct {
	Response
	Alerts   ClientAlerts  `json:"alerts"`
	Threads  ClientThreads `json:"threads"`
	Channels []Mention     `json:"channels"`
	Mpims    []Mention     `json:"mpims"`
	Ims      []Mention     `json:"ims"`
}

type ClientAlerts struct{}

type ClientThreads struct {
	HasUnreads            bool        `json:"has_unreads"`
	MentionCount          int         `json:"mention_count"`
	MentionCountByChannel interface{} `json:"mention_count_by_channel"`
	UnreadCountByChannel  interface{} `json:"unread_count_by_channel"`
}

type Mention struct {
	ID           string `json:"id"`
	LastRead     string `json:"last_read"`
	Latest       string `json:"latest"`
	MentionCount int    `json:"mention_count"`
	HasUnreads   bool   `json:"has_unreads"`
}

// ClientCounts is https://api.slack.com/methods/client.counts
func (s *SlackAPI) ClientCounts(input ClientCountsInput) ClientCountsResponse {
	var out ClientCountsResponse
	if err := s.basePOST("/api/client.counts", input, &out); err != nil {
		return ClientCountsResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
