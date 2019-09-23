package slackapi

// Response defines the expected data from the JSON-encoded API response.
type Response struct {
	Ok       bool   `json:"ok"`
	Error    string `json:"error,omitempty"`
	Needed   string `json:"needed,omitempty"`
	Provided string `json:"provided,omitempty"`
}

// Session defines the expected data from the JSON-encoded API response.
type Session struct {
	Response
	AlreadyOpen bool `json:"already_open"`
	NoOp        bool `json:"no_op"`
	Channel     struct {
		ID string `json:"id"`
	} `json:"channel"`
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
	Messages           []Message `json:"messages"`
	HasMore            bool      `json:"has_more"`
	PinCount           int       `json:"pin_count"`
	UnreadCountDisplay int       `json:"unread_count_display"`
	ResponseMetadata   struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

// MyHistory defines the expected data from the JSON-encoded API response.
type MyHistory struct {
	Filtered int
	Latest   string
	Messages []Message
	Oldest   string
	Total    int
	Username string
}

// DeletedHistory defines the expected data from the JSON-encoded API response.
type DeletedHistory struct {
	Deleted    int
	NotDeleted int
	Messages   []DeletedMessage
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
	IsMpim             bool           `json:"is_mpim"`
	IsOpen             bool           `json:"is_open"`
	LastRead           string         `json:"last_read"`
	Latest             ChannelLatest  `json:"latest"`
	Members            []string       `json:"members"`
	Name               string         `json:"name"`
	NameNormalized     string         `json:"name_normalized"`
	NumMembers         int            `json:"num_members"`
	Purpose            ChannelPurpose `json:"purpose"`
	Topic              ChannelTopic   `json:"topic"`
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

// ChannelRename defines the expected data from the JSON-encoded API response.
type ChannelRename struct {
	Response
	Channel struct {
		ID        string `json:"id"`
		IsChannel bool   `json:"is_channel"`
		IsGroup   bool   `json:"is_group"`
		Name      string `json:"name"`
		Created   int    `json:"created"`
	} `json:"channel"`
}

// ChannelTopic defines the expected data from the JSON-encoded API response.
type ChannelTopic struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

// ChannelPurposeNow defines the expected data from the JSON-encoded API response.
type ChannelPurposeNow struct {
	Response
	Purpose string `json:"purpose"`
}

// ChannelTopicNow defines the expected data from the JSON-encoded API response.
type ChannelTopicNow struct {
	Response
	Topic string `json:"topic"`
}

// AppsList defines the expected data from the JSON-encoded API response.
type AppsList struct {
	Response
	Apps           []AppsListApps `json:"apps"`
	CacheTimestamp string         `json:"cache_ts"`
}

// AppsListApps defines the expected data from the JSON-encoded API response.
type AppsListApps struct {
	ID    string            `json:"id"`
	Name  string            `json:"name"`
	Icons AppsListAppsIcons `json:"icons"`
}

// AppsListAppsIcons defines the expected data from the JSON-encoded API response.
type AppsListAppsIcons struct {
	Image1024 string `json:"image_1024"`
	Image128  string `json:"image_128"`
	Image192  string `json:"image_192"`
	Image32   string `json:"image_32"`
	Image36   string `json:"image_36"`
	Image48   string `json:"image_48"`
	Image512  string `json:"image_512"`
	Image64   string `json:"image_64"`
	Image72   string `json:"image_72"`
	Image96   string `json:"image_96"`
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
