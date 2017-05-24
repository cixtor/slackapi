package slackapi

import (
	"encoding/json"
)

// ResponseTeamAccessLogs defines the JSON-encoded output for TeamAccessLogs.
type ResponseTeamAccessLogs struct {
	Response
	AccessLogs []AccessLog `json:"logins"`
	Needed     string      `json:"needed"`
	Provided   string      `json:"provided"`
}

// ResponseTeamInfo defines the JSON-encoded output for TeamInfo.
type ResponseTeamInfo struct {
	Response
	Team Team `json:"team"`
}

// ResponseTeamProfile defines the JSON-encoded output for TeamProfile.
type ResponseTeamProfile struct {
	Response
	Profile TeamProfile `json:"profile"`
}

// ResponseBillableInfo defines the JSON-encoded output for BillableInfo.
type ResponseBillableInfo struct {
	Response
	BillableInfo map[string]BillableInfo `json:"billable_info"`
}

// AccessLog defines the expected data from the JSON-encoded API response.
type AccessLog struct {
	UserID    string      `json:"user_id"`
	Username  string      `json:"username"`
	DateFirst json.Number `json:"date_first"`
	DateLast  json.Number `json:"date_last"`
	Count     int         `json:"count"`
	IP        string      `json:"ip"`
	UserAgent string      `json:"user_agent"`
	ISP       string      `json:"isp"`
	Country   string      `json:"country"`
	Region    string      `json:"region"`
}

// Team defines the expected data from the JSON-encoded API response.
type Team struct {
	Domain      string   `json:"domain"`
	EmailDomain string   `json:"email_domain"`
	Icon        TeamIcon `json:"icon"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
}

// TeamProfile defines the expected data from the JSON-encoded API response.
type TeamProfile struct {
	Fields []TeamProfileField `json:"fields"`
}

// TeamProfileField defines the expected data from the JSON-encoded API response.
type TeamProfileField struct {
	ID             string      `json:"id"`
	Ordering       int         `json:"ordering"`
	FieldName      string      `json:"field_name"`
	Label          string      `json:"label"`
	Hint           string      `json:"hint"`
	Type           string      `json:"type"`
	PossibleValues interface{} `json:"possible_values"`
	Options        interface{} `json:"options"`
	IsHidden       bool        `json:"is_hidden"`
}

// TeamIcon defines the expected data from the JSON-encoded API response.
type TeamIcon struct {
	Image102      string `json:"image_102"`
	Image132      string `json:"image_132"`
	Image34       string `json:"image_34"`
	Image44       string `json:"image_44"`
	Image68       string `json:"image_68"`
	Image88       string `json:"image_88"`
	ImageOriginal string `json:"image_original"`
}

// BillableInfo defines the expected data from the JSON-encoded API response.
type BillableInfo struct {
	BillingActive bool `json:"billing_active"`
}

// TeamAccessLogs gets the access logs for the current team.
func (s *SlackAPI) TeamAccessLogs(count string, page string) ResponseTeamAccessLogs {
	var response ResponseTeamAccessLogs
	s.GetRequest(&response,
		"team.accessLogs",
		"token",
		"count="+count,
		"page="+page)
	return response
}

// TeamBillableInfo gets billable users information for the current team.
func (s *SlackAPI) TeamBillableInfo(user string) ResponseBillableInfo {
	var response ResponseBillableInfo

	if user == "" {
		s.GetRequest(&response, "team.billableInfo", "token")
	} else {
		s.GetRequest(&response, "team.billableInfo", "token", "user="+s.UsersID(user))
	}

	return response
}

// TeamInfo gets information about the current team.
func (s *SlackAPI) TeamInfo() ResponseTeamInfo {
	var response ResponseTeamInfo
	s.GetRequest(&response, "team.info", "token")
	return response
}

// TeamProfileGet retrieve a team's profile.
func (s *SlackAPI) TeamProfileGet() ResponseTeamProfile {
	var response ResponseTeamProfile
	s.GetRequest(&response, "team.profile.get", "token")
	return response
}
