package slackapi

// APITest is https://api.slack.com/methods/api.test
func (s *SlackAPI) APITest(error string) Response {
	in := struct {
		Error string `json:"error"`
	}{
		Error: error,
	}
	var out Response
	if err := s.baseJSONPOST("/api/api.test", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
