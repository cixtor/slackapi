package slackapi

// SearchArgs defines the data to send to the API service.
type SearchArgs struct {
	Query     string `json:"query"`
	Count     int    `json:"count"`
	Highlight bool   `json:"highlight"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	SortDir   string `json:"sort_dir"`
}

type ResponseSearch struct {
	Response
	Query    string         `json:"query"`
	Files    SearchFiles    `json:"files"`
	Messages SearchMessages `json:"messages"`
}

type SearchMessages struct {
	Matches    []SearchMatches `json:"matches"`
	Pagination Pagination      `json:"pagination"`
	Paging     Paging          `json:"paging"`
	Total      int             `json:"total"`
}

type SearchFiles struct {
	Matches    []File     `json:"matches"`
	Pagination Pagination `json:"pagination"`
	Paging     Paging     `json:"paging"`
	Total      int        `json:"total"`
}

type SearchMatches struct {
	// TODO(cixtor): find a way to convert this into a Message. Currently, the
	// problem is that Channel is not a string, which is what we get from the
	// API when we get a Message.
	Channel   SearchChannel `json:"channel"`
	IID       string        `json:"iid"`
	Permalink string        `json:"permalink"`
	Team      string        `json:"team"`
	Text      string        `json:"text"`
	Timestamp string        `json:"ts"`
	Type      string        `json:"type"`
	User      string        `json:"user"`
	Username  string        `json:"username"`
}

type SearchChannel struct {
	ID                 string        `json:"id"`
	IsExtShared        bool          `json:"is_ext_shared"`
	IsMpim             bool          `json:"is_mpim"`
	IsOrgShared        bool          `json:"is_org_shared"`
	IsPendingExtShared bool          `json:"is_pending_ext_shared"`
	IsPrivate          bool          `json:"is_private"`
	IsShared           bool          `json:"is_shared"`
	Name               string        `json:"name"`
	PendingShared      []interface{} `json:"pending_shared"`
}

func (s *SlackAPI) searchStuff(action string, data SearchArgs) ResponseSearch {
	// set default value for optional field.
	if data.Sort != "timestamp" {
		data.Sort = "score"
	}

	// set default value for optional field.
	if data.SortDir != "asc" {
		data.SortDir = "desc"
	}

	var response ResponseSearch
	s.getRequest(&response, action, data)
	return response
}

// SearchFiles searches for files matching a query.
func (s *SlackAPI) SearchFiles(data SearchArgs) ResponseSearch {
	return s.searchStuff("search.files", data)
}

// SearchMessages searches for messages matching a query.
func (s *SlackAPI) SearchMessages(data SearchArgs) ResponseSearch {
	return s.searchStuff("search.messages", data)
}
