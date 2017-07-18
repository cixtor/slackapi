package slackapi

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

// ResponseDNDStatus defines the JSON-encoded output for end Snooze.
type ResponseDNDStatus struct {
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
	s.PostRequest(&response, "dnd.endDnd", nil)
	return response
}

// DNDEndSnooze ends the current user's snooze mode immediately.
func (s *SlackAPI) DNDEndSnooze() ResponseDNDStatus {
	var response ResponseDNDStatus
	s.PostRequest(&response, "dnd.endSnooze", nil)
	return response
}

// DNDInfo retrieves a user's current "Do Not Disturb" status
func (s *SlackAPI) DNDInfo(user string) ResponseDNDStatus {
	var response ResponseDNDStatus
	s.PostRequest(&response, "dnd.info", struct {
		User string `json:"user"`
	}{user})
	return response
}

// DNDSetSnooze turns on "Do Not Disturb" mode for the current user.
func (s *SlackAPI) DNDSetSnooze(minutes int) ResponseSnoozeStatus {
	var response ResponseSnoozeStatus
	s.PostRequest(&response, "dnd.setSnooze", struct {
		NumMinutes int `json:"num_minutes"`
	}{minutes})
	return response
}
