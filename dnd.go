package slackapi

import (
	"fmt"
	"net/url"
	"strings"
)

// SnoozeDebug defines the JSON-encoded output for SnoozeDebug.
type SnoozeDebug struct {
	SnoozeEndDate string `json:"snooze_end_date,omitempty"`
}

// SnoozeInfo defines the JSON-encoded output for SnoozeInfo.
type SnoozeInfo struct {
	SnoozeEnabled      bool        `json:"snooze_enabled,omitempty"`
	SnoozeEndTime      int         `json:"snooze_endtime,omitempty"`
	SnoozeRemaining    int         `json:"snooze_remaining,omitempty"`
	SnoozeIsIndefinite int         `json:"snooze_is_indefinite,omitempty"`
	SnoozeDebug        SnoozeDebug `json:"snooze_debug,omitempty"`
}

// DNDStatus defines the status of the do not disturb setting.
type DNDStatus struct {
	Enabled            bool `json:"dnd_enabled"`
	NextStartTimestamp int  `json:"next_dnd_start_ts"`
	NextEndTimestamp   int  `json:"next_dnd_end_ts"`
	SnoozeInfo
}

type DNDStatusResponse struct {
	Response
	DNDStatus
}

type SnoozeStatusResponse struct {
	Response
	SnoozeInfo
}

// DNDEndDnd ends the current user's "Do Not Disturb" session immediately.
func (s *SlackAPI) DNDEndDnd() Response {
	var response Response
	s.postRequest(&response, "dnd.endDnd", nil)
	return response
}

// DNDEndSnooze is https://api.slack.com/methods/dnd.endSnooze
func (s *SlackAPI) DNDEndSnooze() DNDStatusResponse {
	var out DNDStatusResponse
	if err := s.baseFormPOST("/api/dnd.endSnooze", nil, &out); err != nil {
		return DNDStatusResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// DNDInfo is https://api.slack.com/methods/dnd.info
func (s *SlackAPI) DNDInfo(user string) DNDStatusResponse {
	in := url.Values{"user": {user}}
	var out DNDStatusResponse
	if err := s.baseFormPOST("/api/dnd.info", in, &out); err != nil {
		return DNDStatusResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// DNDSetSnooze is https://api.slack.com/methods/dnd.setSnooze
func (s *SlackAPI) DNDSetSnooze(minutes int) SnoozeStatusResponse {
	in := url.Values{"num_minutes": {fmt.Sprintf("%d", minutes)}}
	var out SnoozeStatusResponse
	if err := s.baseFormPOST("/api/dnd.setSnooze", in, &out); err != nil {
		return SnoozeStatusResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type DNDTeamResponse struct {
	Response
	Cached bool                 `json:"cached"`
	Users  map[string]DNDStatus `json:"users"`
}

// DNDTeamInfo is https://api.slack.com/methods/dnd.teamInfo
func (s *SlackAPI) DNDTeamInfo(users []string) DNDTeamResponse {
	in := url.Values{"users": {strings.Join(users, ",")}}
	var out DNDTeamResponse
	if err := s.baseFormPOST("/api/dnd.teamInfo", in, &out); err != nil {
		return DNDTeamResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
