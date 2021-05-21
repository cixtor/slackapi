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

type TeamBillableInfoResponse struct {
	Response
	BillableInfo map[string]BillableInfo `json:"billable_info"`
}

type BillableInfo struct {
	BillingActive bool `json:"billing_active"`
}

// TeamBillableInfo is https://api.slack.com/methods/team.billableInfo
func (s *SlackAPI) TeamBillableInfo(teamID string, user string) TeamBillableInfoResponse {
	in := url.Values{}

	if teamID != "" {
		in.Add("team_id", teamID)
	}

	if user != "" {
		in.Add("user", user)
	}

	var out TeamBillableInfoResponse
	if err := s.baseGET("/api/team.billableInfo", in, &out); err != nil {
		return TeamBillableInfoResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type TeamBillingInfoResponse struct {
	Response
	Plan string `json:"plan"`
}

// TeamBillingInfo is https://api.slack.com/methods/team.billing.info
func (s *SlackAPI) TeamBillingInfo() TeamBillingInfoResponse {
	in := url.Values{}
	var out TeamBillingInfoResponse
	if err := s.baseGET("/api/team.billing.info", in, &out); err != nil {
		return TeamBillingInfoResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type TeamInfoResponse struct {
	Response
	Team Team `json:"team"`
}

// Team defines the expected data from the JSON-encoded API response.
type Team struct {
	Domain      string   `json:"domain"`
	EmailDomain string   `json:"email_domain"`
	Icon        TeamIcon `json:"icon"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
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

// TeamInfo gets information about the current team.
//
// The team parameter is to get info on, if omitted, will return information
// about the current team. Will only return team that the authenticated token
// is allowed to see through external shared channels
func (s *SlackAPI) TeamInfo(team string) TeamInfoResponse {
	in := url.Values{}
	if team != "" {
		in.Add("team", team)
	}
	var out TeamInfoResponse
	if err := s.baseGET("/api/team.info", in, &out); err != nil {
		return TeamInfoResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// ResponseTeamProfile defines the JSON-encoded output for TeamProfile.
type ResponseTeamProfile struct {
	Response
	Profile TeamProfile `json:"profile"`
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

// TeamProfileGet retrieve a team's profile.
func (s *SlackAPI) TeamProfileGet() ResponseTeamProfile {
	var response ResponseTeamProfile
	s.getRequest(&response, "team.profile.get", nil)
	return response
}
