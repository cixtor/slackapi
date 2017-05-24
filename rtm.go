package slackapi

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"os"
)

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

// NewRTM connects to the real time messaging websocket.
func (s *SlackAPI) NewRTM() (*websocket.Conn, error) {
	resp := s.RTMStart()

	return websocket.Dial(resp.URL, "", "https://api.slack.com")
}

// RTMStart Starts a Real Time Messaging session.
func (s *SlackAPI) RTMStart() RTMResponse {
	var response RTMResponse
	s.GetRequest(&response, "rtm.start", "token")
	return response
}

// RTMMessage reads and returns a message from the websocket connection.
func (s *SlackAPI) RTMMessage(ws *websocket.Conn, res interface{}) error {
	return websocket.JSON.Receive(ws, &res)
}

// RTMEvents prints all the websocket events.
func (s *SlackAPI) RTMEvents() {
	ws, err := s.NewRTM()

	if err != nil {
		log.Fatal(err)
		return
	}

	var data interface{}

	for {
		if err := s.RTMMessage(ws, &data); err != nil {
			log.Println(err)
			continue
		}

		if err := json.NewEncoder(os.Stdout).Encode(data); err != nil {
			log.Println("JSON encode;", err)
			continue
		}
	}

}
