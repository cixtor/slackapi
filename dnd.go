package slackapi

import (
	"net/url"
	"strings"
)

// SnoozeDebug defines the JSON-encoded output for SnoozeDebug.
type SnoozeDebug struct {
	SnoozeEndDate string `json:"snooze_end_date,omitempty"`
}

// SnoozeInfo defines the JSON-encoded output for SnoozeInfo.
type SnoozeInfo struct {
	SnoozeEnabled   bool        `json:"snooze_enabled,omitempty"`
	SnoozeEndTime   int         `json:"snooze_endtime,omitempty"`
	SnoozeRemaining int         `json:"snooze_remaining,omitempty"`
	SnoozeDebug     SnoozeDebug `json:"snooze_debug,omitempty"`
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

// ResponseSnoozeStatus defines the JSON-encoded output for set Snooze.
type ResponseSnoozeStatus struct {
	Response
	SnoozeInfo
}

// DNDEndDnd ends the current user's "Do Not Disturb" session immediately.
func (s *SlackAPI) DNDEndDnd() Response {
	var response Response
	s.postRequest(&response, "dnd.endDnd", nil)
	return response
}

// DNDEndSnooze ends the current user's snooze mode immediately.
func (s *SlackAPI) DNDEndSnooze() ResponseDNDStatus {
	var response ResponseDNDStatus
	s.postRequest(&response, "dnd.endSnooze", nil)
	return response
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

// DNDSetSnooze turns on "Do Not Disturb" mode for the current user.
func (s *SlackAPI) DNDSetSnooze(minutes int) ResponseSnoozeStatus {
	var response ResponseSnoozeStatus
	s.postRequest(&response, "dnd.setSnooze", struct {
		NumMinutes int `json:"num_minutes"`
	}{minutes})
	return response
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
