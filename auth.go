package slackapi

import (
	"net/url"
)

// ResponseRevocation defines the JSON-encoded output for Revocation.
type ResponseRevocation struct {
	Response
	Revoked bool `json:"revoked"`
}

// APITest checks API calling code.
func (s *SlackAPI) APITest() Response {
	var response Response
	s.getRequest(&response, "api.test", nil)
	return response
}

// AppsList lists associated applications.
func (s *SlackAPI) AppsList() AppsList {
	var response AppsList
	s.getRequest(&response, "apps.list", nil)
	return response
}

// AuthRevoke revokes a token.
func (s *SlackAPI) AuthRevoke() ResponseRevocation {
	var response ResponseRevocation
	s.getRequest(&response, "auth.revoke", nil)
	return response
}

// AuthTest checks authentication and identity.
func (s *SlackAPI) AuthTest() Owner {
	if s.owner.Ok {
		return s.owner
	}

	var response Owner
	s.getRequest(&response, "auth.test", nil)
	s.owner = response

	return response
}
