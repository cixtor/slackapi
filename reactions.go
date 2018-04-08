package slackapi

// ReactionArgs defines the data to send to the API service.
type ReactionArgs struct {
	Name        string `json:"name"`
	Channel     string `json:"channel"`
	File        string `json:"file"`
	FileComment string `json:"file_comment"`
	Full        bool   `json:"full"`
	Timestamp   string `json:"timestamp"`
}

// ReactionListArgs defines the data to send to the API service.
type ReactionListArgs struct {
	Count int    `json:"count"`
	Full  bool   `json:"full"`
	Page  int    `json:"page"`
	User  string `json:"user"`
}

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
func (s *SlackAPI) ReactionsAdd(data ReactionArgs) Response {
	var response Response
	s.postRequest(&response, "reactions.add", data)
	return response
}

// ReactionsGet gets reactions for an item.
func (s *SlackAPI) ReactionsGet(data ReactionArgs) ResponseReactionsGet {
	var response ResponseReactionsGet
	s.postRequest(&response, "reactions.get", data)
	return response
}

// ReactionsList lists reactions made by a user.
func (s *SlackAPI) ReactionsList(data ReactionListArgs) ResponseReactionsList {
	if data.Count == 0 {
		data.Count = 100
	}
	var response ResponseReactionsList
	s.getRequest(&response, "reactions.list", data)
	return response
}

// ReactionsRemove removes a reaction from an item.
func (s *SlackAPI) ReactionsRemove(data ReactionArgs) Response {
	var response Response
	s.postRequest(&response, "reactions.remove", data)
	return response
}
