package main

func (s *SlackAPI) ReactionsAdd(name string, channel string, timestamp string) {
	var response interface{}

	// Remove reaction from a file, file comment, or message.
	if channel[0] == 'F' {
		s.GetRequest(&response,
			"reactions.add",
			"token",
			"name="+name,
			"file="+channel)
	} else {
		s.GetRequest(&response,
			"reactions.add",
			"token",
			"name="+name,
			"channel="+channel,
			"timestamp="+timestamp)
	}

	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsGet(channel string, timestamp string) {
	var response interface{}

	// Remove reaction from a file, file comment, or message.
	if channel[0] == 'F' {
		s.GetRequest(&response,
			"reactions.get",
			"token",
			"file="+channel)
	} else {
		s.GetRequest(&response,
			"reactions.get",
			"token",
			"channel="+channel,
			"timestamp="+timestamp)
	}

	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsList(userid string) {
	var response interface{}

	if userid == "" {
		s.GetRequest(&response,
			"reactions.list",
			"token",
			"full=true",
			"count=100")
	} else {
		s.GetRequest(&response,
			"reactions.list",
			"token",
			"full=true",
			"count=100",
			"user="+userid)
	}
	s.PrintJson(response)
}

func (s *SlackAPI) ReactionsRemove(name string, channel string, timestamp string) {
	var response interface{}

	// Remove reaction from a file, file comment, or message.
	if channel[0] == 'F' {
		s.GetRequest(&response,
			"reactions.remove",
			"token",
			"name="+name,
			"file="+channel)
	} else {
		s.GetRequest(&response,
			"reactions.remove",
			"token",
			"name="+name,
			"channel="+channel,
			"timestamp="+timestamp)
	}

	s.PrintJson(response)
}
