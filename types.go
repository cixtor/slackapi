package slackapi

import (
	"encoding/json"
)

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
type History struct {
	Response
	HasMore            bool      `json:"has_more"`
	Messages           []Message `json:"messages"`
	UnreadCountDisplay int       `json:"unread_count_display"`
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
	Channel string  `json:"channel"`
	Message Message `json:"message"`
	Text    string  `json:"text"`
	Ts      string  `json:"ts"`
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

// Message defines the expected data from the JSON-encoded API response.
type Message struct {
	Attachments  []Attachment `json:"attachments"`
	BotID        string       `json:"bot_id"`
	Channel      string       `json:"channel"`
	DisplayAsBot bool         `json:"display_as_bot"`
	File         File         `json:"file"`
	Subtype      string       `json:"subtype"`
	Text         string       `json:"text"`
	Ts           string       `json:"ts"`
	Type         string       `json:"type"`
	Upload       bool         `json:"upload"`
	User         string       `json:"user"`
	Username     string       `json:"username"`
}

// ModifiedMessage defines the expected data from the JSON-encoded API response.
type ModifiedMessage struct {
	Response
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Ts      string `json:"ts"`
}

// Comment defines the expected data from the JSON-encoded API response.
type Comment struct {
	ID        string `json:"id,omitempty"`
	Created   int64  `json:"created,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	User      string `json:"user,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// Attachment defines the expected data from the JSON-encoded API response.
type Attachment struct {
	Fallback    string      `json:"fallback"`
	FromURL     string      `json:"from_url"`
	ID          int         `json:"id"`
	ImageBytes  int         `json:"image_bytes"`
	ImageHeight int         `json:"image_height"`
	ImageURL    string      `json:"image_url"`
	ImageWidth  int         `json:"image_width"`
	ServiceName string      `json:"service_name"`
	Text        string      `json:"text"`
	ThumbHeight int         `json:"thumb_height"`
	ThumbURL    string      `json:"thumb_url"`
	ThumbWidth  int         `json:"thumb_width"`
	Title       string      `json:"title"`
	TitleLink   string      `json:"title_link"`
	Ts          json.Number `json:"ts"`
}

// DeletedMessage defines the expected data from the JSON-encoded API response.
type DeletedMessage struct {
	Deleted bool
	Text    string
	Ts      string
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
	NumMembers         int            `json:"num_members"`
	Purpose            ChannelPurpose `json:"purpose"`
	Topic              ChannelTopic   `json:"topic"`
	UnreadCount        int            `json:"unread_count"`
	UnreadCountDisplay int            `json:"unread_count_display"`
}

// ChannelLatest defines the expected data from the JSON-encoded API response.
type ChannelLatest struct {
	Text string `json:"text"`
	Ts   string `json:"ts"`
	Type string `json:"type"`
	User string `json:"user"`
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
	Apps    []AppsListApps `json:"apps"`
	CacheTs string         `json:"cache_ts"`
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

// Pagination defines the expected data from the JSON-encoded API response.
type Pagination struct {
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
