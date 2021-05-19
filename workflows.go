package slackapi

// WorkflowsStepCompleted https://api.slack.com/methods/workflows.stepCompleted
func (s *SlackAPI) WorkflowsStepCompleted() Response {
	var out Response
	if err := s.basePOST("/api/workflows.stepCompleted", nil, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// WorkflowsStepFailed https://api.slack.com/methods/workflows.stepFailed
func (s *SlackAPI) WorkflowsStepFailed() Response {
	var out Response
	if err := s.basePOST("/api/workflows.stepFailed", nil, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// WorkflowsUpdateStep https://api.slack.com/methods/workflows.updateStep
func (s *SlackAPI) WorkflowsUpdateStep() Response {
	var out Response
	if err := s.basePOST("/api/workflows.updateStep", nil, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
