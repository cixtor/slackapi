package slackapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type TeamAccessLogsInput struct {
	// End of time range of logs to include in results (inclusive).
	Before string `json:"before"`
	// Number of items to return per page.
	Count int `json:"count"`
	// Page number of results to return.
	Page int `json:"page"`
	// Encoded team id to get logs from, required if org token is used.
	TeamID string `json:"team_id"`
}

// TeamAccessLogsResponse defines the JSON-encoded output for TeamAccessLogs.
type TeamAccessLogsResponse struct {
	Response
	Logins []AccessLog `json:"logins"`
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

// TeamAccessLogs is https://api.slack.com/methods/team.accessLogs
func (s *SlackAPI) TeamAccessLogs(input TeamAccessLogsInput) TeamAccessLogsResponse {
	in := url.Values{}

	if input.Before != "" {
		in.Add("before", input.Before)
	}

	if input.Count > 0 {
		in.Add("count", strconv.Itoa(input.Count))
	} else {
		in.Add("count", "100")
	}

	if input.Page > 0 {
		in.Add("page", strconv.Itoa(input.Page))
	}

	if input.TeamID != "" {
		in.Add("team_id", input.TeamID)
	}

	var out TeamAccessLogsResponse
	if err := s.baseGET("/api/team.accessLogs", in, &out); err != nil {
		return TeamAccessLogsResponse{Response: Response{Error: err.Error()}}
	}
	return out
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

// TeamBillableInfo gets billable users information for the current team.
func (s *SlackAPI) TeamBillableInfo(user string) ResponseBillableInfo {
	var response ResponseBillableInfo
	s.getRequest(&response, "team.billableInfo", struct {
		User string `json:"user"`
	}{user})
	return response
}

// TeamInfo gets information about the current team.
func (s *SlackAPI) TeamInfo() ResponseTeamInfo {
	var response ResponseTeamInfo
	s.getRequest(&response, "team.info", nil)
	return response
}

// TeamProfileGet retrieve a team's profile.
func (s *SlackAPI) TeamProfileGet() ResponseTeamProfile {
	var response ResponseTeamProfile
	s.getRequest(&response, "team.profile.get", nil)
	return response
}
