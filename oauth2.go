package slackapi

import (
	"net/url"
)

// OAuth2Exchange is https://api.slack.com/methods/oauth.v2.exchange
func (s *SlackAPI) OAuth2Exchange(clientID string, clientSecret string) Response {
	in := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}
	var out Response
	if err := s.baseFormPOST("/api/oauth.v2.exchange", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
