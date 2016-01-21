package main

func (s *SlackAPI) ApiTest() {
	var response interface{}
	s.GetRequest(&response, "api.test")
	s.PrintJson(response)
}

func (s *SlackAPI) AuthTest() {
	var response interface{}
	s.GetRequest(&response, "auth.test", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) EmojiList() {
	var response interface{}
	s.GetRequest(&response, "emoji.list", "token")
	s.PrintJson(response)
}

func (s *SlackAPI) TeamInfo() {
	var response interface{}
	s.GetRequest(&response, "team.info", "token")
	s.PrintJson(response)
}
