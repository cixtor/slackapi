package slackapi

import (
	"encoding/json"
	"net/url"
)

// PinsAdd is https://api.slack.com/methods/pins.add
func (s *SlackAPI) PinsAdd(channel string, itemid string) Response {
	in := url.Values{"channel": {channel}}
	var out Response

	if len(itemid) >= 3 && itemid[0:2] == "Fc" {
		in.Add("file_comment", itemid) // remove pinned file comment.
	} else if len(itemid) >= 2 && itemid[0] == 'F' {
		in.Add("file", itemid) // remove pinned file.
	} else {
		in.Add("timestamp", itemid) // remove pinned message.
	}

	if err := s.baseFormPOST("/api/pins.add", in, &out); err != nil {
		return Response{Error: err.Error()}
	}

	return out
}

type PinsListResponse struct {
	Response
	Items []PinsListItem `json:"items"`
}

type PinsListItem struct {
	Channel   string      `json:"channel"`
	Created   json.Number `json:"created"`
	CreatedBy string      `json:"created_by"`
	Message   Message     `json:"message"`
	Type      string      `json:"type"`
}

// PinsList is https://api.slack.com/methods/pins.list
func (s *SlackAPI) PinsList(channel string) PinsListResponse {
	in := url.Values{"channel": {channel}}
	var out PinsListResponse
	if err := s.baseGET("/api/pins.list", in, &out); err != nil {
		return PinsListResponse{Response: Response{Error: err.Error()}}
	}
	return out
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
