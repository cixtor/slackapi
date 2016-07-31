package main

type Revocation struct {
	Response
	Revoked bool `json:"revoked"`
}

func (s *SlackAPI) ApiTest() Response {
	var response Response
	s.GetRequest(&response, "api.test")
	return response
}

func (s *SlackAPI) AppsList() AppsList {
	var response AppsList
	s.GetRequest(&response, "apps.list", "token")
	return response
}

func (s *SlackAPI) AuthTest() Owner {
	if s.Owner.Ok == true {
		return s.Owner
	}

	var response Owner
	s.GetRequest(&response, "auth.test", "token")
	s.Owner = response

	return response
}

func (s *SlackAPI) AuthRevoke() Revocation {
	var response Revocation
	s.GetRequest(&response, "auth.revoke", "token")
	return response
}

func (s *SlackAPI) AuthRevokeTest() Revocation {
	var response Revocation
	s.GetRequest(&response, "auth.revoke", "token", "test=true")
	return response
}
