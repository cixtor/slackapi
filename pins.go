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

// PinsAdd pins an item to a channel.
func (s *SlackAPI) PinsAdd(channel string, itemid string) Response {
	var response Response

	if len(itemid) >= 3 && itemid[0:2] == "Fc" {
		/* remove pinned file comment */
		s.postRequest(&response, "pins.add", struct {
			Channel     string `json:"channel"`
			FileComment string `json:"file_comment"`
		}{
			Channel:     channel,
			FileComment: itemid,
		})
	} else if len(itemid) >= 2 && itemid[0] == 'F' {
		/* remove pinned file */
		s.postRequest(&response, "pins.add", struct {
			Channel string `json:"channel"`
			File    string `json:"file"`
		}{
			Channel: channel,
			File:    itemid,
		})
	} else {
		/* remove pinned message */
		s.postRequest(&response, "pins.add", struct {
			Channel   string `json:"channel"`
			Timestamp string `json:"timestamp"`
		}{
			Channel:   channel,
			Timestamp: itemid,
		})
	}

	return response
}

// PinsList lists items pinned to a channel.
func (s *SlackAPI) PinsList(channel string) ResponsePinsList {
	var response ResponsePinsList
	s.getRequest(&response, "pins.list", struct {
		Channel string `json:"channel"`
	}{
		Channel: channel,
	})
	return response
}

// PinsRemove lists items pinned to a channel.
func (s *SlackAPI) PinsRemove(channel string, itemid string) Response {
	var response Response

	if len(itemid) >= 3 && itemid[0:2] == "Fc" {
		/* remove pinned file comment */
		s.postRequest(&response, "pins.remove", struct {
			Channel     string `json:"channel"`
			FileComment string `json:"file_comment"`
		}{
			Channel:     channel,
			FileComment: itemid,
		})
	} else if len(itemid) >= 2 && itemid[0] == 'F' {
		/* remove pinned file */
		s.postRequest(&response, "pins.remove", struct {
			Channel string `json:"channel"`
			File    string `json:"file"`
		}{
			Channel: channel,
			File:    itemid,
		})
	} else {
		/* remove pinned message */
		s.postRequest(&response, "pins.remove", struct {
			Channel   string `json:"channel"`
			Timestamp string `json:"timestamp"`
		}{
			Channel:   channel,
			Timestamp: itemid,
		})
	}

	return response
}
