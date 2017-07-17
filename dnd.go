package slackapi

// SnoozeDebug defines the JSON-encoded output for SnoozeDebug.
type SnoozeDebug struct {
	SnoozeEndDate string `json:"snooze_end_date"`
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

// ResponseSetSnooze defines the JSON-encoded output for Snooze.
type ResponseSetSnooze struct {
	Response
	SnoozeInfo
}

// DNDEndDnd ends the current user's "Do Not Disturb" session immediately.
func (s *SlackAPI) DNDEndDnd() Response {
	var response Response
	s.PostRequest(&response, "dnd.endDnd", nil)
	return response
}

// DNDSetSnooze turns on "Do Not Disturb" mode for the current user.
func (s *SlackAPI) DNDSetSnooze(minutes int) ResponseSetSnooze {
	var response ResponseSetSnooze
	s.PostRequest(&response, "dnd.setSnooze", struct {
		NumMinutes int `json:"num_minutes"`
	}{minutes})
	return response
}
