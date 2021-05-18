package slackapi

// WorkflowsStepCompleted https://api.slack.com/methods/workflows.stepCompleted
func (s *SlackAPI) WorkflowsStepCompleted() Response {
	var out Response
	if err := s.basePOST("/api/workflows.stepCompleted", nil, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
