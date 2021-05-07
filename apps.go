package slackapi

type AppsConnectionsOpenResponse struct {
	Response
	URL string `json:"url"`
}

// AppsConnectionsOpen is apps.connections.open
func (s *SlackAPI) AppsConnectionsOpen() AppsConnectionsOpenResponse {
	in := struct{}{}
	var out AppsConnectionsOpenResponse
	if err := s.basePOST("/api/apps.connections.open", in, &out); err != nil {
		return AppsConnectionsOpenResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type AppsEventAuthorizationsListInput struct {
	EventContext string `json:"event_context"`
	Cursor       string `json:"cursor"`
	Limit        int    `json:"limit"`
}

type AppsEventAuthorizationsListResponse struct {
	Response
	Authorizations []AppAuthorization `json:"authorizations"`
}

type AppAuthorization struct {
	EnterpriseID string `json:"enterprise_id"`
	TeamID       string `json:"team_id"`
	UserID       string `json:"user_id"`
	IsBot        string `json:"is_bot"`
}

// AppsEventAuthorizationsList is https://api.slack.com/methods/apps.event.authorizations.list
func (s *SlackAPI) AppsEventAuthorizationsList(input AppsEventAuthorizationsListInput) AppsEventAuthorizationsListResponse {
	var out AppsEventAuthorizationsListResponse
	if err := s.basePOST("/api/apps.event.authorizations.list", input, &out); err != nil {
		return AppsEventAuthorizationsListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
