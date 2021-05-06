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
