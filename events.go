package slackapi

import "encoding/json"

/* inspired by nlopes/slack/rtm */

// Event defines the JSON-encoded output for RTM events.
type Event struct {
	Type string `json:"type,omitempty"`
	Data interface{}
}

// ErrorEvent represents the websocket errors.
type ErrorEvent struct {
	Text string
}

// AccountsChangedEvent represents the accounts changed event.
type AccountsChangedEvent struct {
	Type string `json:"type"`
}

// BotAddedEvent represents the bot added event.
type BotAddedEvent struct {
	Type string `json:"type"`
	Bot  Bot    `json:"bot"`
}

// BotChangedEvent represents the bot changed event.
type BotChangedEvent struct {
	Type string `json:"type"`
	Bot  Bot    `json:"bot"`
}

// ChannelInfoEvent represents the Channel info event.
type ChannelInfoEvent struct {
	// channel_left
	// channel_deleted
	// channel_archive
	// channel_unarchive
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	User      string `json:"user,omitempty"`
	Timestamp string `json:"ts,omitempty"`
}

// ChannelCreatedInfo represents the information associated with the Channel created event
type ChannelCreatedInfo struct {
	ID        string `json:"id"`
	IsChannel bool   `json:"is_channel"`
	Name      string `json:"name"`
	Created   int    `json:"created"`
	Creator   string `json:"creator"`
}

// ChannelRenameInfo represents the information associated with a Channel rename event
type ChannelRenameInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

// ChannelArchiveEvent represents the Channel archive event.
type ChannelArchiveEvent ChannelInfoEvent

// ChannelCreatedEvent represents the Channel created event.
type ChannelCreatedEvent struct {
	Type           string             `json:"type"`
	Channel        ChannelCreatedInfo `json:"channel"`
	EventTimestamp string             `json:"event_ts"`
}

// ChannelDeletedEvent represents the Channel deleted event.
type ChannelDeletedEvent ChannelInfoEvent

// ChannelHistoryChangedEvent represents the Channel history changed event.
type ChannelHistoryChangedEvent struct {
	Type           string `json:"type"`
	Latest         string `json:"latest"`
	Timestamp      string `json:"ts"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelJoinedEvent represents the Channel joined event.
type ChannelJoinedEvent struct {
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
}

// ChannelLeftEvent represents the Channel left event.
type ChannelLeftEvent ChannelInfoEvent

// ChannelMarkedEvent represents the Channel marked event.
type ChannelMarkedEvent ChannelInfoEvent

// ChannelRenameEvent represents the Channel rename event.
type ChannelRenameEvent struct {
	Type      string            `json:"type"`
	Channel   ChannelRenameInfo `json:"channel"`
	Timestamp string            `json:"event_ts"`
}

// ChannelUnarchiveEvent represents the Channel unarchive event.
type ChannelUnarchiveEvent ChannelInfoEvent

// CommandsChangedEvent represents the commands changed event.
type CommandsChangedEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
}

// DNDUpdatedEvent represents the update event for Do Not Disturb
type DNDUpdatedEvent struct {
	Type   string    `json:"type"`
	User   string    `json:"user"`
	Status DNDStatus `json:"dnd_status"`
}

// EmailDomainChangedEvent represents the email domain changed event.
type EmailDomainChangedEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
	EmailDomain    string `json:"email_domain"`
}

// EmojiChangedEvent represents the emoji changed event.
type EmojiChangedEvent struct {
	Type           string   `json:"type"`
	SubType        string   `json:"subtype"`
	Name           string   `json:"name"`
	Names          []string `json:"names"`
	Value          string   `json:"value"`
	EventTimestamp string   `json:"event_ts"`
}

// FileActionEvent represents the File action event.
type fileActionEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
	File           File   `json:"file"`
	// FileID is used for FileDeletedEvent
	FileID string `json:"file_id,omitempty"`
}

// FileChangeEvent represents the File change event.
type FileChangeEvent fileActionEvent

// FileCommentAddedEvent represents the File comment added event.
type FileCommentAddedEvent struct {
	fileActionEvent
	Comment Comment `json:"comment"`
}

// FileCommentDeletedEvent represents the File comment deleted event.
type FileCommentDeletedEvent struct {
	fileActionEvent
	Comment string `json:"comment"`
}

// FileCommentEditedEvent represents the File comment edited event.
type FileCommentEditedEvent struct {
	fileActionEvent
	Comment Comment `json:"comment"`
}

// FileCreatedEvent represents the File created event.
type FileCreatedEvent fileActionEvent

// FileDeletedEvent represents the File deleted event.
type FileDeletedEvent fileActionEvent

// FilePrivateEvent represents the File private event.
type FilePrivateEvent fileActionEvent

// FilePublicEvent represents the File public event.
type FilePublicEvent fileActionEvent

// FileSharedEvent represents the File shared event.
type FileSharedEvent fileActionEvent

// FileUnsharedEvent represents the File unshared event.
type FileUnsharedEvent fileActionEvent

// GroupCreatedEvent represents the Group created event.
type GroupCreatedEvent struct {
	Type    string             `json:"type"`
	User    string             `json:"user"`
	Channel ChannelCreatedInfo `json:"channel"`
}

// GroupRenameInfo represents the group info related to the renamed group
type GroupRenameInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

// GroupArchiveEvent represents the Group archive event.
type GroupArchiveEvent ChannelInfoEvent

// GroupCloseEvent represents the Group close event.
type GroupCloseEvent ChannelInfoEvent

// GroupHistoryChangedEvent represents the Group history changed event.
type GroupHistoryChangedEvent ChannelHistoryChangedEvent

// GroupJoinedEvent represents the Group joined event.
type GroupJoinedEvent ChannelJoinedEvent

// GroupLeftEvent represents the Group left event.
type GroupLeftEvent ChannelInfoEvent

// GroupMarkedEvent represents the Group marked event.
type GroupMarkedEvent ChannelInfoEvent

// GroupOpenEvent represents the Group open event.
type GroupOpenEvent ChannelInfoEvent

// GroupRenameEvent represents the Group rename event.
type GroupRenameEvent struct {
	Type      string          `json:"type"`
	Group     GroupRenameInfo `json:"channel"`
	Timestamp string          `json:"ts"`
}

// GroupUnarchiveEvent represents the Group unarchive event.
type GroupUnarchiveEvent ChannelInfoEvent

// HelloEvent represents the hello event.
type HelloEvent struct{}

// IMCloseEvent represents the IM close event.
type IMCloseEvent ChannelInfoEvent

// IMCreatedEvent represents the IM created event.
type IMCreatedEvent struct {
	Type    string             `json:"type"`
	User    string             `json:"user"`
	Channel ChannelCreatedInfo `json:"channel"`
}

// IMHistoryChangedEvent represents the IM history changed event.
type IMHistoryChangedEvent ChannelHistoryChangedEvent

// IMMarkedEvent represents the IM marked event.
type IMMarkedEvent ChannelInfoEvent

// IMMarkedHistoryChanged represents the IM marked history changed event
type IMMarkedHistoryChanged ChannelInfoEvent

// IMOpenEvent represents the IM open event.
type IMOpenEvent ChannelInfoEvent

// ManualPresenceChangeEvent represents the manual presence change event.
type ManualPresenceChangeEvent struct {
	Type     string `json:"type"`
	Presence string `json:"presence"`
}

// MessageEvent represents the message event.
type MessageEvent Message

type pinEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Item           Item   `json:"item"`
	Channel        string `json:"channel_id"`
	EventTimestamp string `json:"event_ts"`
	HasPins        bool   `json:"has_pins,omitempty"`
}

// PinAddedEvent represents the Pin added event.
type PinAddedEvent pinEvent

// PinRemovedEvent represents the Pin removed event.
type PinRemovedEvent pinEvent

// PrefChangeEvent represents a user preferences change event
type PrefChangeEvent struct {
	Type  string          `json:"type"`
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

// PresenceChangeEvent represents the presence change event.
type PresenceChangeEvent struct {
	Type     string `json:"type"`
	Presence string `json:"presence"`
	User     string `json:"user"`
}

type reactionEvent struct {
	Type           string       `json:"type"`
	User           string       `json:"user"`
	ItemUser       string       `json:"item_user"`
	Item           ReactionItem `json:"item"`
	Reaction       string       `json:"reaction"`
	EventTimestamp string       `json:"event_ts"`
}

// ReactionAddedEvent represents the Reaction added event.
type ReactionAddedEvent reactionEvent

// ReactionRemovedEvent represents the Reaction removed event.
type ReactionRemovedEvent reactionEvent

// ReconnectURLEvent represents the receiving reconnect url event.
type ReconnectURLEvent struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type starEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Item           Item   `json:"item"`
	EventTimestamp string `json:"event_ts"`
}

// StarAddedEvent represents the Star added event.
type StarAddedEvent starEvent

// StarRemovedEvent represents the Star removed event.
type StarRemovedEvent starEvent

// TeamDomainChangeEvent represents the Team domain change event.
type TeamDomainChangeEvent struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Domain string `json:"domain"`
}

// TeamJoinEvent represents the Team join event.
type TeamJoinEvent struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// TeamMigrationStartedEvent represents the Team migration started event.
type TeamMigrationStartedEvent struct {
	Type string `json:"type"`
}

// TeamPrefChangeEvent represents the Team preference change event.
type TeamPrefChangeEvent struct {
	Type  string   `json:"type"`
	Name  string   `json:"name,omitempty"`
	Value []string `json:"value,omitempty"`
}

// TeamRenameEvent represents the Team rename event.
type TeamRenameEvent struct {
	Type           string `json:"type"`
	Name           string `json:"name,omitempty"`
	EventTimestamp string `json:"event_ts,omitempty"`
}

// UserChangeEvent represents the user change event.
type UserChangeEvent struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// UserTypingEvent represents the user typing event.
type UserTypingEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}
