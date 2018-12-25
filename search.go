package slackapi

type ResponseSearchMessages struct {
	Response
	Query    string         `json:"query"`
	Messages SearchMessages `json:"messages"`
}

type SearchMessages struct {
	Matches    []SearchMatches `json:"matches"`
	Pagination Pagination      `json:"pagination"`
	Paging     Paging          `json:"paging"`
	Total      int             `json:"total"`
}

type SearchMatches struct {
	Channel   SearchChannel `json:"channel"`
	IID       string        `json:"iid"`
	Permalink string        `json:"permalink"`
	Team      string        `json:"team"`
	Text      string        `json:"text"`
	Ts        string        `json:"ts"`
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

type Pagination struct {
	First      int `json:"first"`
	Last       int `json:"last"`
	Page       int `json:"page"`
	PageCount  int `json:"page_count"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
}

// SearchArgs defines the data to send to the API service.
type SearchArgs struct {
	Query     string `json:"query"`
	Count     int    `json:"count"`
	Highlight bool   `json:"highlight"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	SortDir   string `json:"sort_dir"`
}

// SearchMessages searches for messages matching a query.
func (s *SlackAPI) SearchMessages(data SearchArgs) ResponseSearchMessages {
	// set default value for optional field.
	if data.Sort != "timestamp" {
		data.Sort = "score"
	}

	// set default value for optional field.
	if data.SortDir != "asc" {
		data.SortDir = "desc"
	}

	var response ResponseSearchMessages
	s.getRequest(&response, "search.messages", data)
	return response
}
