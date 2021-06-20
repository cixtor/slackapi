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

type ClientShouldReloadInput struct {
	VersionTs       int    `json:"version_ts"`
	BuildVersionTs  int    `json:"build_version_ts"`
	ConfigVersionTs int    `json:"config_version_ts"`
	TeamIDs         string `json:"team_ids"`
}

type ClientShouldReloadResponse struct {
	Response
	BuildVersionEnabled    bool `json:"build_version_enabled"`
	ShouldReload           bool `json:"should_reload"`
	ClientMinVersion       int  `json:"client_min_version"`
	ClientMinConfigVersion int  `json:"client_min_config_version"`
}

// ClientShouldReload is https://api.slack.com/methods/client.shouldReload
func (s *SlackAPI) ClientShouldReload(input ClientShouldReloadInput) ClientShouldReloadResponse {
	var out ClientShouldReloadResponse
	if err := s.basePOST("/api/client.shouldReload", input, &out); err != nil {
		return ClientShouldReloadResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
