package slackapi

type HelpIssuesListResponse struct {
	Response
	Issues []string `json:"issues"`
}

// HelpIssuesList is https://api.slack.com/methods/help.issues.list
func (s *SlackAPI) HelpIssuesList() HelpIssuesListResponse {
	var out HelpIssuesListResponse
	if err := s.baseGET("/api/help.issues.list", nil, &out); err != nil {
		return HelpIssuesListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
