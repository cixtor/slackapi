package slackapi

// ResponseRevocation defines the JSON-encoded output for Revocation.
type ResponseRevocation struct {
	Response
	Revoked bool `json:"revoked"`
}

// APITest checks API calling code.
func (s *SlackAPI) APITest() Response {
	var response Response
	s.GetRequest(&response, "api.test", nil)
	return response
}

// AppsList lists associated applications.
func (s *SlackAPI) AppsList() AppsList {
	var response AppsList
	s.GetRequest(&response, "apps.list", nil)
	return response
}

// AuthRevoke revokes a token.
func (s *SlackAPI) AuthRevoke() ResponseRevocation {
	var response ResponseRevocation
	s.GetRequest(&response, "auth.revoke", nil)
	return response
}

// AuthTest checks authentication and identity.
func (s *SlackAPI) AuthTest() Owner {
	if s.Owner.Ok {
		return s.Owner
	}

	var response Owner
	s.GetRequest(&response, "auth.test", nil)
	s.Owner = response

	return response
}
