package slackapi

// ResponseReactionsGet defines the JSON-encoded output for ReactionsGet.
type ResponseReactionsGet struct {
	Response
	ReactionItem
}

// ResponseReactionsList defines the JSON-encoded output for ReactionsList.
type ResponseReactionsList struct {
	Response
	Items  []ReactionItem `json:"items"`
	Paging Pagination     `json:"paging"`
}

// ReactionItem defines the expected data from the JSON-encoded API response.
type ReactionItem struct {
	Channel     string          `json:"channel"`
	File        ReactionFile    `json:"file"`
	FileComment string          `json:"file_comment"`
	Message     ReactionMessage `json:"message"`
	Type        string          `json:"type"`
	Timestamp   string          `json:"ts"`
}

// ReactionMessage defines the expected data from the JSON-encoded API response.
type ReactionMessage struct {
	Reactions []Reaction `json:"reactions"`
	Text      string     `json:"text"`
	Timestamp string     `json:"ts"`
	Type      string     `json:"type"`
	User      string     `json:"user"`
}

// ReactionFile defines the expected data from the JSON-encoded API response.
type ReactionFile struct {
	File
	Reactions []Reaction `json:"reactions"`
}

// Reaction defines the expected data from the JSON-encoded API response.
type Reaction struct {
	Count int      `json:"count"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

// ReactionsAdd adds a reaction to an item.
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

// ReactionsGet gets reactions for an item.
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

// ReactionsList lists reactions made by a user.
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

// ReactionsRemove removes a reaction from an item.
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
