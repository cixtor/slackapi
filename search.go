package slackapi

import (
	"net/url"
	"strconv"
)

// SearchArgs defines the data to send to the API service.
type SearchArgs struct {
	Query     string `json:"query"`
	Count     int    `json:"count"`
	Highlight bool   `json:"highlight"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	SortDir   string `json:"sort_dir"`
}

type SearchUsersArgs struct {
	Query               string `json:"query"`
	Count               int    `json:"count"`
	Fuzz                int    `json:"fuzz"`
	UAX29Tokenizer      bool   `json:"uax29_tokenizer"`
	SearchProfileFields bool   `json:"search_profile_fields"`
}

type ResponseSearch struct {
	Response
	Query    string         `json:"query"`
	Files    SearchFiles    `json:"files"`
	Posts    SearchPosts    `json:"posts"`
	Messages SearchMessages `json:"messages"`
}

type ResponseSearchUsers struct {
	Response
	Results           []User   `json:"results"`
	PresenceActiveIds []string `json:"presence_active_ids"`
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

type SearchPosts struct {
	Matches    []interface{} `json:"matches"`
	Pagination Pagination    `json:"pagination"`
	Paging     Paging        `json:"paging"`
	Total      int           `json:"total"`
}

type SearchMatches struct {
	// TODO(cixtor): find a way to convert this into a Message. Currently, the
	// problem is that Channel is not a string, which is what we get from the
	// API when we get a Message.
	Attachments []Attachment  `json:"attachments"`
	Channel     SearchChannel `json:"channel"`
	IID         string        `json:"iid"`
	Permalink   string        `json:"permalink"`
	Team        string        `json:"team"`
	Text        string        `json:"text"`
	Timestamp   string        `json:"ts"`
	Type        string        `json:"type"`
	User        string        `json:"user"`
	Username    string        `json:"username"`
}

type SearchChannel struct {
	ID                 string        `json:"id"`
	IsChannel          bool          `json:"is_channel"`
	IsExtShared        bool          `json:"is_ext_shared"`
	IsGroup            bool          `json:"is_group"`
	IsIM               bool          `json:"is_im"`
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

// SearchAll searches for messages and files matching a query.
func (s *SlackAPI) SearchAll(data SearchArgs) ResponseSearch {
	return s.searchStuff("search.all", data)
}

// SearchFiles searches for files matching a query.
func (s *SlackAPI) SearchFiles(data SearchArgs) ResponseSearch {
	return s.searchStuff("search.files", data)
}

// SearchMessages searches for messages matching a query.
func (s *SlackAPI) SearchMessages(data SearchArgs) ResponseSearch {
	return s.searchStuff("search.messages", data)
}

type SearchModulesInput struct {
	ExcludeMyChannels    bool   `json:"exclude_my_channels"`
	ExtraMessageData     bool   `json:"extra_message_data"`
	Highlight            bool   `json:"highlight"`
	NoUserProfile        bool   `json:"no_user_profile"`
	Sort                 string `json:"sort"`
	SortDir              string `json:"sort_dir"`
	Browse               string `json:"browse"`
	BrowseSessionID      string `json:"browse_session_id"`
	ClientReqID          string `json:"client_req_id"`
	Count                int    `json:"count"`
	MaxFilterSuggestions int    `json:"max_filter_suggestions"`
	Module               string `json:"module"`
	Page                 int    `json:"page"`
	ChannelType          string `json:"channel_type"`
	Extracts             int    `json:"extracts"`
	Query                string `json:"query"`
	Team                 string `json:"team"`
}

type SearchModulesResponse struct {
	Response
	Module     string      `json:"module"`
	Filters    interface{} `json:"filters"`
	Query      string      `json:"query"`
	Pagination Pagination  `json:"pagination"`
	Items      []Channel   `json:"items"`
}

// SearchModules is https://api.slack.com/methods/search.modules
func (s *SlackAPI) SearchModules(input SearchModulesInput) SearchModulesResponse {
	in := url.Values{}
	if input.ExcludeMyChannels {
		in.Add("exclude_my_channels", "true")
	}
	if input.ExtraMessageData {
		in.Add("extra_message_data", "true")
	}
	if input.Highlight {
		in.Add("highlight", "true")
	}
	if input.NoUserProfile {
		in.Add("no_user_profile", "true")
	}
	if input.Sort != "" {
		in.Add("sort", input.Sort)
	}
	if input.SortDir != "" {
		in.Add("sort_dir", input.SortDir)
	}
	if input.Browse != "" {
		in.Add("browse", input.Browse)
	}
	if input.BrowseSessionID != "" {
		in.Add("browse_session_id", input.BrowseSessionID)
	}
	if input.ClientReqID != "" {
		in.Add("client_req_id", input.ClientReqID)
	}
	if input.Count > 0 {
		in.Add("count", strconv.Itoa(input.Count))
	}
	if input.MaxFilterSuggestions > 0 {
		in.Add("max_filter_suggestions", strconv.Itoa(input.MaxFilterSuggestions))
	}
	if input.Module != "" {
		in.Add("module", input.Module)
	}
	if input.Page > 0 {
		in.Add("page", strconv.Itoa(input.Page))
	}
	if input.ChannelType != "" {
		in.Add("channel_type", input.ChannelType)
	}
	if input.Extracts > 0 {
		in.Add("extracts", strconv.Itoa(input.Extracts))
	}
	if input.Query != "" {
		in.Add("query", input.Query)
	}
	if input.Team != "" {
		in.Add("team", input.Team)
	}
	var out SearchModulesResponse
	if err := s.baseGET("/api/search.modules", in, &out); err != nil {
		return SearchModulesResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// SearchUsers searches for users matching a query.
func (s *SlackAPI) SearchUsers(input SearchUsersArgs) (ResponseSearchUsers, error) {
	owner, err := s.AuthTest()

	if err != nil {
		return ResponseSearchUsers{}, err
	}

	if owner.TeamID == "" {
		return ResponseSearchUsers{Response: owner.Response}, nil
	}

	var response ResponseSearchUsers
	if err := s.edgePOST("/cache/"+owner.TeamID+"/users/search", input, &response); err != nil {
		return ResponseSearchUsers{}, err
	}
	return response, nil
}
