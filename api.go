package slackapi

import (
	"net/url"
)

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

type APIGetFlannelHTTPURLResponse struct {
	Response
	URL        string `json:"url"`
	TTLSeconds int    `json:"ttl_seconds"`
}

// APIGetFlannelHTTPURL is https://api.slack.com/methods/api.getFlannelHttpUrl
func (s *SlackAPI) APIGetFlannelHTTPURL() APIGetFlannelHTTPURLResponse {
	in := url.Values{}
	in.Add("include_external_workspaces", "1")
	var out APIGetFlannelHTTPURLResponse
	if err := s.baseJSONPOST("/api/api.getFlannelHttpUrl", in, &out); err != nil {
		return APIGetFlannelHTTPURLResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
