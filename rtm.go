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
	conn      *websocket.Conn
	connURL   string
	Events    chan Event
	rawEvents chan json.RawMessage
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
var EventTypes = map[string]interface{}{}

// NewRTM connects to the real time messaging websocket.
func (s *SlackAPI) NewRTM() (*RTM, error) {
	var response RTMResponse
	s.GetRequest(&response, "rtm.start", "token")

	ws, err := websocket.Dial(response.URL, "", "https://api.slack.com")

	if err != nil {
		return &RTM{}, err
	}

	return &RTM{
		conn:      ws,
		connURL:   response.URL,
		Events:    make(chan Event, 50),
		rawEvents: make(chan json.RawMessage),
	}, nil
}

// ManageEvents controls the websocket events.
func (rtm *RTM) ManageEvents() {
	keepalive := make(chan bool)

	go rtm.handleIncomingEvents(keepalive)

	rtm.handleEvents(keepalive)
}

func (rtm *RTM) handleEvents(keepRunning chan bool) {
	for {
		select {
		case rawEvent := <-rtm.rawEvents:
			rtm.handleRawEvent(rawEvent)
		}
	}
}

func (rtm *RTM) handleIncomingEvents(keepalive <-chan bool) {
	for {
		select {
		case <-keepalive:
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

	rtm.handleEvent(event.Type, rawEvent)
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
