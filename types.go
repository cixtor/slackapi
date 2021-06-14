package slackapi

// Response defines the expected data from the JSON-encoded API response.
type Response struct {
	Ok               bool             `json:"ok"`
	Error            string           `json:"error,omitempty"`
	Warning          string           `json:"warning,omitempty"`
	Needed           string           `json:"needed,omitempty"`
	Provided         string           `json:"provided,omitempty"`
	Errors           []ErrMsgs        `json:"errors,omitempty"`
	ResponseMetadata ResponseMetadata `json:"response_metadata,omitempty"`
}

type ErrMsgs struct {
	Message string `json:"message"`
	Pointer string `json:"pointer"`
}

type ResponseMetadata struct {
	Messages   []string `json:"messages"`
	Warnings   []string `json:"warnings"`
	NextCursor string   `json:"next_cursor"`
}

// Owner defines the expected data from the JSON-encoded API response.
type Owner struct {
	Response
	Team   string `json:"team"`
	TeamID string `json:"team_id"`
	URL    string `json:"url"`
	User   string `json:"user"`
	UserID string `json:"user_id"`
}

// History defines the expected data from the JSON-encoded API response.
//
// Example:
//
//   {
//       "ok": true,
//       "messages": [
//           {
//               "type": "message",
//               "user": "U012AB3CDE",
//               "text": "I find you punny and would like to smell your nose letter",
//               "ts": "1512085950.000216"
//           },
//           {
//               "type": "message",
//               "user": "U061F7AUR",
//               "text": "What, you want to smell my shoes better?",
//               "ts": "1512104434.000490"
//           }
//       ],
//       "has_more": true,
//       "pin_count": 0,
//       "response_metadata": {
//           "next_cursor": "bmV4dF90czoxNTEyMDg1ODYxMDAwNTQz"
//       }
//   }
type History struct {
	Response
	Messages           []Message        `json:"messages"`
	HasMore            bool             `json:"has_more"`
	PinCount           int              `json:"pin_count"`
	UnreadCountDisplay int              `json:"unread_count_display"`
	ResponseMetadata   ResponseMetadata `json:"response_metadata"`
}

// Post defines the expected data from the JSON-encoded API response.
type Post struct {
	Response
	Channel   string  `json:"channel"`
	Timestamp string  `json:"ts"`
	Message   Message `json:"message"`
}

// Item defines the expected data for: message, file, or file comment.
type Item struct {
	Type      string   `json:"type"`
	Channel   string   `json:"channel,omitempty"`
	Message   *Message `json:"message,omitempty"`
	File      *File    `json:"file,omitempty"`
	Comment   *Comment `json:"comment,omitempty"`
	Timestamp string   `json:"ts,omitempty"`
}

// Channel defines the expected data from the JSON-encoded API response.
type Channel struct {
	Created            int            `json:"created"`
	Creator            string         `json:"creator"`
	ID                 string         `json:"id"`
	IsArchived         bool           `json:"is_archived"`
	IsChannel          bool           `json:"is_channel"`
	IsGeneral          bool           `json:"is_general"`
	IsGroup            bool           `json:"is_group"`
	IsMember           bool           `json:"is_member"`
	IsIm               bool           `json:"is_im"`
	IsMpim             bool           `json:"is_mpim"`
	IsOrgShared        bool           `json:"is_org_shared"`
	IsOpen             bool           `json:"is_open"`
	User               string         `json:"user"`
	LastRead           string         `json:"last_read"`
	Latest             ChannelLatest  `json:"latest"`
	Members            []string       `json:"members"`
	Name               string         `json:"name"`
	NameNormalized     string         `json:"name_normalized"`
	NumMembers         int            `json:"num_members"`
	Purpose            ChannelPurpose `json:"purpose"`
	Topic              ChannelTopic   `json:"topic"`
	Priority           int            `json:"priority"`
	UnreadCount        int            `json:"unread_count"`
	UnreadCountDisplay int            `json:"unread_count_display"`
}

// ChannelLatest defines the expected data from the JSON-encoded API response.
type ChannelLatest struct {
	Text      string `json:"text"`
	Timestamp string `json:"ts"`
	Type      string `json:"type"`
	User      string `json:"user"`
}

// ChannelPurpose defines the expected data from the JSON-encoded API response.
type ChannelPurpose struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

// ChannelTopic defines the expected data from the JSON-encoded API response.
type ChannelTopic struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

type Pagination struct {
	First      int `json:"first"`
	Last       int `json:"last"`
	Page       int `json:"page"`
	PageCount  int `json:"page_count"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
}

// Paging defines the expected data from the JSON-encoded API response.
type Paging struct {
	Count int `json:"count"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

// Fields defines the expected data from the JSON-encoded API response.
type Fields map[string]Field

// Field defines the expected data from the JSON-encoded API response.
type Field struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Alt   string `json:"alt"`
}
