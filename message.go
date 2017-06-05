package slackapi

// Message defines the expected data from the JSON-encoded API response.
type Message struct {
	Attachments      []Attachment   `json:"attachments,omitempty"`
	BotID            string         `json:"bot_id,omitempty"`
	Channel          string         `json:"channel,omitempty"`
	Comment          *Comment       `json:"comment,omitempty"`
	DeletedTimestamp string         `json:"deleted_ts,omitempty"`
	DisplayAsBot     bool           `json:"display_as_bot"`
	Edited           *Edited        `json:"edited,omitempty"`
	EventTimestamp   string         `json:"event_ts,omitempty"`
	File             *File          `json:"file,omitempty"`
	Hidden           bool           `json:"hidden,omitempty"`
	Icons            *Icon          `json:"icons,omitempty"`
	Inviter          string         `json:"inviter,omitempty"`
	IsStarred        bool           `json:"is_starred,omitempty"`
	ItemType         string         `json:"item_type,omitempty"`
	Members          []string       `json:"members,omitempty"`
	Name             string         `json:"name,omitempty"`
	OldName          string         `json:"old_name,omitempty"`
	ParentUserID     string         `json:"parent_user_id,omitempty"`
	PinnedTo         []string       `json:"pinned_to, omitempty"`
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
	Upload           bool           `json:"upload,omitempty"`
	User             string         `json:"user,omitempty"`
	Username         string         `json:"username,omitempty"`
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

// DeletedMessage defines the expected data from the JSON-encoded API response.
type DeletedMessage struct {
	Deleted   bool
	Text      string
	Timestamp string
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
