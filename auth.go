package main

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
