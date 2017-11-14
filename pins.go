package slackapi

import (
	"encoding/json"
)

// ResponsePinsList defines the JSON-encoded output for PinsList.
type ResponsePinsList struct {
	Response
	Items []PinsListItem `json:"items"`
}

// PinsListItem defines the JSON-encoded output for one pinned message.
type PinsListItem struct {
	Channel   string      `json:"channel"`
	Created   json.Number `json:"created"`
	CreatedBy string      `json:"created_by"`
	Message   Message     `json:"message"`
	Type      string      `json:"type"`
}

// PinsList lists items pinned to a channel.
func (s *SlackAPI) PinsList(channel string) ResponsePinsList {
	var response ResponsePinsList
	s.GetRequest(&response, "pins.list", struct {
		Channel string `json:"channel"`
	}{s.ChannelsID(channel)})
	return response
}
