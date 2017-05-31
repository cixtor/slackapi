package slackapi

// ResponseRevocation defines the JSON-encoded output for Revocation.
type ResponseRevocation struct {
	Response
	Revoked bool `json:"revoked"`
}

// APITest checks API calling code.
func (s *SlackAPI) APITest() Response {
	var response Response
	s.GetRequest(&response, "api.test")
	return response
}

// AppsList lists associated applications.
func (s *SlackAPI) AppsList() AppsList {
	var response AppsList
	s.GetRequest(&response, "apps.list", "token")
	return response
}

// AuthRevoke revokes a token.
func (s *SlackAPI) AuthRevoke() ResponseRevocation {
	var response ResponseRevocation
	s.GetRequest(&response, "auth.revoke", "token")
	return response
}

// AuthRevokeTest rest the token revocation.
func (s *SlackAPI) AuthRevokeTest() ResponseRevocation {
	var response ResponseRevocation
	s.GetRequest(&response, "auth.revoke", "token", "test=true")
	return response
}

// AuthTest checks authentication and identity.
func (s *SlackAPI) AuthTest() Owner {
	if s.Owner.Ok {
		return s.Owner
	}

	var response Owner
	s.GetRequest(&response, "auth.test", "token")
	s.Owner = response

	return response
}
