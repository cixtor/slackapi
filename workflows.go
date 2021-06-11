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

type WorkflowsUpdateStepInput struct {
	// Context identifier that maps to the correct workflow step execution.
	WorkflowStepExecuteID string `json:"workflow_step_execute_id"`
	// JSON key-value map of inputs required from a user during configuration.
	// This is the data your app expects to receive when the workflow step
	// starts.
	//
	// Example:
	//
	//   {"title":{"value":"The Title"},"submitter":{"value":"{{user}}"}}
	Inputs interface{} `json:"inputs"`
	// An JSON array of output objects used during step execution. This is the
	// data your app agrees to provide when your workflow step was executed.
	//
	// Example:
	//
	//   [
	//     {"name":"ticket_id","type":"text","label":"Ticket ID"},
	//     {"name":"title","type":"text","label":"Title"}
	//   ]
	Outputs interface{} `json:"outputs"`
	// An optional field that can be used to override app image that is shown
	// in the Workflow Builder.
	StepImageURL string `json:"step_image_url"`
	// An optional field that can be used to override the step name that is
	// shown in the Workflow Builder.
	StepName string `json:"step_name"`
}

// WorkflowsUpdateStep https://api.slack.com/methods/workflows.updateStep
func (s *SlackAPI) WorkflowsUpdateStep(input WorkflowsUpdateStepInput) Response {
	var out Response
	if err := s.basePOST("/api/workflows.updateStep", input, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
