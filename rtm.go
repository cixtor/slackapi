package slackapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"

	"golang.org/x/net/websocket"
)

// RTM is the real time messaging.
type RTM struct {
	conn         *websocket.Conn
	connURL      string
	Events       chan Event
	rawEvents    chan json.RawMessage
	stopEvents   chan bool
	stopListener chan bool
}

// RTMResponse defines the JSON-encoded output for RTM connection.
type RTMResponse struct {
	Response
	// Self                    null `json:"self"`
	// Team                    null `json:"team"`
	// LatestEventTS           null `json:"latest_event_ts"`
	// Channels                null `json:"channels"`
	// Groups                  null `json:"groups"`
	// IMS                     null `json:"ims"`
	// CacheTS                 null `json:"cache_ts"`
	// ReadOnlyChannels        null `json:"read_only_channels"`
	// CanManageSharedChannels null `json:"can_manage_shared_channels"`
	// Subteams                null `json:"subteams"`
	// DND                     null `json:"dnd"`
	// Users                   null `json:"users"`
	// CacheVersion            null `json:"cache_version"`
	// CacheTSVersion          null `json:"cache_ts_version"`
	// Bots                    null `json:"bots"`
	URL string `json:"url"`
}

// ACKMessage is used for messages received in reply to other messages
type ACKMessage struct {
	Response
	ReplyTo   int    `json:"reply_to"`
	Text      string `json:"text"`
	Timestamp string `json:"ts"`
}

// EventTypes represents the data type for each websocket event.
var EventTypes = map[string]interface{}{
	"accounts_changed":        AccountsChangedEvent{},
	"bot_added":               BotAddedEvent{},
	"bot_changed":             BotChangedEvent{},
	"channel_archive":         ChannelArchiveEvent{},
	"channel_created":         ChannelCreatedEvent{},
	"channel_deleted":         ChannelDeletedEvent{},
	"channel_history_changed": ChannelHistoryChangedEvent{},
	"channel_joined":          ChannelJoinedEvent{},
	"channel_left":            ChannelLeftEvent{},
	"channel_marked":          ChannelMarkedEvent{},
	"channel_rename":          ChannelRenameEvent{},
	"channel_unarchive":       ChannelUnarchiveEvent{},
	"commands_changed":        CommandsChangedEvent{},
	"dnd_updated":             DNDUpdatedEvent{},
	"dnd_updated_user":        DNDUpdatedEvent{},
	"email_domain_changed":    EmailDomainChangedEvent{},
	"emoji_changed":           EmojiChangedEvent{},
	"file_change":             FileChangeEvent{},
	"file_comment_added":      FileCommentAddedEvent{},
	"file_comment_deleted":    FileCommentDeletedEvent{},
	"file_comment_edited":     FileCommentEditedEvent{},
	"file_created":            FileCreatedEvent{},
	"file_deleted":            FileDeletedEvent{},
	"file_private":            FilePrivateEvent{},
	"file_public":             FilePublicEvent{},
	"file_shared":             FileSharedEvent{},
	"file_unshared":           FileUnsharedEvent{},
	"group_archive":           GroupArchiveEvent{},
	"group_close":             GroupCloseEvent{},
	"group_history_changed":   GroupHistoryChangedEvent{},
	"group_joined":            GroupJoinedEvent{},
	"group_left":              GroupLeftEvent{},
	"group_marked":            GroupMarkedEvent{},
	"group_open":              GroupOpenEvent{},
	"group_rename":            GroupRenameEvent{},
	"group_unarchive":         GroupUnarchiveEvent{},
	"hello":                   HelloEvent{},
	"im_close":                IMCloseEvent{},
	"im_created":              IMCreatedEvent{},
	"im_history_changed":      IMHistoryChangedEvent{},
	"im_marked":               IMMarkedEvent{},
	"im_open":                 IMOpenEvent{},
	"manual_presence_change":  ManualPresenceChangeEvent{},
	"message":                 MessageEvent{},
	"pin_added":               PinAddedEvent{},
	"pin_removed":             PinRemovedEvent{},
	"pref_change":             PrefChangeEvent{},
	"presence_change":         PresenceChangeEvent{},
	"reaction_added":          ReactionAddedEvent{},
	"reaction_removed":        ReactionRemovedEvent{},
	"reconnect_url":           ReconnectURLEvent{},
	"star_added":              StarAddedEvent{},
	"star_removed":            StarRemovedEvent{},
	"team_domain_change":      TeamDomainChangeEvent{},
	"team_join":               TeamJoinEvent{},
	"team_migration_started":  TeamMigrationStartedEvent{},
	"team_pref_change":        TeamPrefChangeEvent{},
	"team_rename":             TeamRenameEvent{},
	"user_change":             UserChangeEvent{},
	"user_typing":             UserTypingEvent{},
}

// NewRTM connects to the real time messaging websocket.
func (s *SlackAPI) NewRTM() (*RTM, error) {
	var response RTMResponse
	s.GetRequest(&response, "rtm.start", "token")

	ws, err := websocket.Dial(response.URL, "", "https://api.slack.com")

	if err != nil {
		return &RTM{}, err
	}

	return &RTM{
		conn:         ws,
		connURL:      response.URL,
		Events:       make(chan Event),
		rawEvents:    make(chan json.RawMessage),
		stopEvents:   make(chan bool, 1),
		stopListener: make(chan bool, 1),
	}, nil
}

// Stop kills the connection.
func (rtm *RTM) Stop() {
	close(rtm.Events)
	close(rtm.rawEvents)

	rtm.stopEvents <- true
	rtm.stopListener <- true
}

// ManageEvents controls the websocket events.
func (rtm *RTM) ManageEvents() {
	go rtm.handleEvents()

	go rtm.handleIncomingEvents()
}

func (rtm *RTM) handleEvents() {
	for {
		select {
		case <-rtm.stopEvents:
			return
		case rawEvent := <-rtm.rawEvents:
			rtm.handleRawEvent(rawEvent)
		}
	}
}

func (rtm *RTM) handleIncomingEvents() {
	for {
		select {
		case <-rtm.stopListener:
			return
		default:
			rtm.receiveIncomingEvent()
		}
	}
}

// receiveIncomingEvent inserts the websocket events into a queue.
func (rtm *RTM) receiveIncomingEvent() {
	event := json.RawMessage{}

	if err := websocket.JSON.Receive(rtm.conn, &event); err != nil {
		if err == io.EOF {
			log.Fatal("Connection dropped")
			return
		}

		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{err.Error()}}
		return
	}

	if len(event) == 0 {
		err := fmt.Errorf("invalid event; %s", event)
		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{err.Error()}}
		return
	}

	rtm.rawEvents <- event
}

func (rtm *RTM) handleRawEvent(rawEvent json.RawMessage) {
	event := &Event{}

	if err := json.Unmarshal(rawEvent, event); err != nil {
		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{err.Error()}}
		return
	}

	switch event.Type {
	case "":
		rtm.handleACK(rawEvent)
	case "hello":
		rtm.Events <- Event{Type: "hello", Data: &HelloEvent{}}
	default:
		rtm.handleEvent(event.Type, rawEvent)
	}
}

func (rtm *RTM) handleEvent(_type string, event json.RawMessage) {
	v, exists := EventTypes[_type]

	if !exists {
		err := fmt.Errorf("unsupported event %q: %s", _type, string(event))
		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{err.Error()}}
		return
	}

	t := reflect.TypeOf(v)
	recvEvent := reflect.New(t).Interface()

	if err := json.Unmarshal(event, recvEvent); err != nil {
		err = fmt.Errorf("unmarshall event %q: %s", _type, string(event))
		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{err.Error()}}
		return
	}
	rtm.Events <- Event{_type, recvEvent}
}

func (rtm *RTM) handleACK(event json.RawMessage) {
	ack := &ACKMessage{}

	if err := json.Unmarshal(event, ack); err != nil {
		err = fmt.Errorf("ack unmarshal; %s: %s", err.Error(), string(event))
		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{err.Error()}}
		return
	}

	if !ack.Ok {
		rtm.Events <- Event{Type: "error", Data: &ErrorEvent{ack.Error}}
		return
	}

	rtm.Events <- Event{Type: "ack", Data: ack}
}
