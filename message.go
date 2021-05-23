package slackapi

// MessageArgs defines the data to send to the API service.
type MessageArgs struct {
	Attachments    []Attachment `json:"attachments"`
	Channel        string       `json:"channel"`
	IconEmoji      string       `json:"icon_emoji"`
	IconURL        string       `json:"icon_url"`
	Parse          string       `json:"parse"`
	Text           string       `json:"text"`
	ThreadTs       string       `json:"thread_ts"`
	Ts             string       `json:"ts"`
	Username       string       `json:"username"`
	AsUser         bool         `json:"as_user"`
	LinkNames      bool         `json:"link_names"`
	ReplyBroadcast bool         `json:"reply_broadcast"`
	UnfurlLinks    bool         `json:"unfurl_links"`
	UnfurlMedia    bool         `json:"unfurl_media"`
	Markdown       bool         `json:"mrkdwn"`
}

// Message defines the expected data from the JSON-encoded API response.
type Message struct {
	IID              string         `json:"iid,omitempty"`
	Attachments      []Attachment   `json:"attachments,omitempty"`
	BotID            string         `json:"bot_id,omitempty"`
	Channel          string         `json:"channel,omitempty"`
	Comment          *Comment       `json:"comment,omitempty"`
	DeletedTimestamp string         `json:"deleted_ts,omitempty"`
	Edited           *Edited        `json:"edited,omitempty"`
	EventTimestamp   string         `json:"event_ts,omitempty"`
	File             *File          `json:"file,omitempty"`
	Icons            *Icon          `json:"icons,omitempty"`
	Inviter          string         `json:"inviter,omitempty"`
	ItemType         string         `json:"item_type,omitempty"`
	Members          []string       `json:"members,omitempty"`
	Name             string         `json:"name,omitempty"`
	OldName          string         `json:"old_name,omitempty"`
	ParentUserID     string         `json:"parent_user_id,omitempty"`
	Permalink        string         `json:"permalink,omitempty"`
	PinnedTo         []string       `json:"pinned_to,omitempty"`
	Purpose          string         `json:"purpose,omitempty"`
	Reactions        []ReactionItem `json:"reactions,omitempty"`
	Replies          []Reply        `json:"replies,omitempty"`
	ReplyCount       int            `json:"reply_count,omitempty"`
	ReplyTo          int            `json:"reply_to,omitempty"`
	Subtype          string         `json:"subtype,omitempty"`
	Team             string         `json:"team,omitempty"`
	Text             string         `json:"text,omitempty"`
	ThreadTimestamp  string         `json:"thread_ts,omitempty"`
	Timestamp        string         `json:"ts,omitempty"`
	Topic            string         `json:"topic,omitempty"`
	Type             string         `json:"type,omitempty"`
	User             string         `json:"user,omitempty"`
	Username         string         `json:"username,omitempty"`
	DisplayAsBot     bool           `json:"display_as_bot"`
	Hidden           bool           `json:"hidden,omitempty"`
	IsStarred        bool           `json:"is_starred,omitempty"`
	Upload           bool           `json:"upload,omitempty"`
}

// Comment defines the expected data from the JSON-encoded API response.
type Comment struct {
	ID        string `json:"id,omitempty"`
	Created   int64  `json:"created,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	User      string `json:"user,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// ModifiedMessage defines the expected data from the JSON-encoded API response.
type ModifiedMessage struct {
	Response
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	Timestamp string `json:"ts"`
}

// Edited defines the structure of a modified message.
type Edited struct {
	User      string `json:"user,omitempty"`
	Timestamp string `json:"ts,omitempty"`
}

// Reply defines the structure of a message reply.
type Reply struct {
	User      string `json:"user,omitempty"`
	Timestamp string `json:"ts,omitempty"`
}

// Icon defines the structure of an application icon.
type Icon struct {
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}
