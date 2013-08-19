package main

type ResponseReactionsGet struct {
	Response
	ReactionItem
}

type ResponseReactionsList struct {
	Response
	Items  []ReactionItem `json:"items"`
	Paging Pagination     `json:"paging"`
}

type ReactionItem struct {
	Channel string          `json:"channel"`
	File    ReactionFile    `json:"file"`
	Message ReactionMessage `json:"message"`
	Type    string          `json:"type"`
}

type ReactionMessage struct {
	Reactions []Reaction `json:"reactions"`
	Text      string     `json:"text"`
	Ts        string     `json:"ts"`
	Type      string     `json:"type"`
	User      string     `json:"user"`
}

type ReactionFile struct {
	File
	Reactions []Reaction `json:"reactions"`
}

type Reaction struct {
	Count int      `json:"count"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

func (s *SlackAPI) ReactionsAdd(name string, resource string, unique string) Response {
	var response Response

	if resource[0] == 'F' {
		if unique == "" {
			s.GetRequest(&response,
				"reactions.add",
				"token",
				"name="+name,
				"file="+resource)
		} else {
			s.GetRequest(&response,
				"reactions.add",
				"token",
				"name="+name,
				"file_comment="+unique)
		}
	} else {
		s.GetRequest(&response,
			"reactions.add",
			"token",
			"name="+name,
			"channel="+resource,
			"timestamp="+unique)
	}

	return response
}

func (s *SlackAPI) ReactionsGet(resource string, unique string) ResponseReactionsGet {
	var response ResponseReactionsGet

	if resource[0] == 'F' {
		if unique == "" {
			s.GetRequest(&response,
				"reactions.get",
				"token",
				"file="+resource)
		} else {
			s.GetRequest(&response,
				"reactions.get",
				"token",
				"file_comment="+unique)
		}
	} else {
		s.GetRequest(&response,
			"reactions.get",
			"token",
			"channel="+resource,
			"timestamp="+unique)
	}

	return response
}

func (s *SlackAPI) ReactionsList(userid string) ResponseReactionsList {
	var response ResponseReactionsList
	s.GetRequest(&response,
		"reactions.list",
		"token",
		"full=true",
		"count=100",
		"user="+userid)
	return response
}

func (s *SlackAPI) ReactionsRemove(name string, resource string, unique string) Response {
	var response Response

	if resource[0] == 'F' {
		if unique == "" {
			s.GetRequest(&response,
				"reactions.remove",
				"token",
				"name="+name,
				"file="+resource)
		} else {
			s.GetRequest(&response,
				"reactions.remove",
				"token",
				"name="+name,
				"file_comment="+unique)
		}
	} else {
		s.GetRequest(&response,
			"reactions.remove",
			"token",
			"name="+name,
			"channel="+resource,
			"timestamp="+unique)
	}

	return response
}
