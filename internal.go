package main

type ResponseIssues struct {
	Response
	Issues []string `json:"issues"`
}

func (s *SlackAPI) HelpIssuesList() ResponseIssues {
	var response ResponseIssues
	s.GetRequest(&response, "help.issues.list", "token")
	return response
}
