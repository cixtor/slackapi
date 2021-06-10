package slackapi

type WorkflowsStepCompletedInput struct {
	// Context identifier that maps to the correct workflow step execution.
	WorkflowStepExecuteID string `json:"workflow_step_execute_id"`
	// Key-value object of outputs from your step. Keys of this object
	// reflect the configured key properties of your outputs array from
	// your workflow_step object.
	Outputs map[string]string `json:"outputs"`
}

// WorkflowsStepCompleted https://api.slack.com/methods/workflows.stepCompleted
func (s *SlackAPI) WorkflowsStepCompleted(input WorkflowsStepCompletedInput) Response {
	var out Response
	if err := s.basePOST("/api/workflows.stepCompleted", input, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

type WorkflowsStepFailedInput struct {
	// Context identifier that maps to the correct workflow step execution.
	WorkflowStepExecuteID string `json:"workflow_step_execute_id"`
	// A JSON-based object with a message property that should contain a
	// human readable error message.
	Error WorkflowError
}

type WorkflowError struct {
	Message string `json:"message"`
}

// WorkflowsStepFailed https://api.slack.com/methods/workflows.stepFailed
func (s *SlackAPI) WorkflowsStepFailed(input WorkflowsStepFailedInput) Response {
	var out Response
	if err := s.basePOST("/api/workflows.stepFailed", input, &out); err != nil {
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
