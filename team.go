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

type TeamChannelsInfoInput struct {
	TeamID          string   `json:"-"`
	CheckMembership bool     `json:"check_membership"`
	ChannelIDs      []string `json:"ids"`
}

type TeamChannelsInfoResponse struct {
	Response
	Results []Channel `json:"results"`
}

// TeamChannelsInfo is https://api.slack.com/methods/team.channels.info
func (s *SlackAPI) TeamChannelsInfo(input TeamChannelsInfoInput) TeamChannelsInfoResponse {
	var out TeamChannelsInfoResponse
	if err := s.edgePOST("/cache/"+input.TeamID+"/channels/info", input, &out); err != nil {
		return TeamChannelsInfoResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type TeamInfoResponse struct {
	Response
	Team Team `json:"team"`
}

type Team struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	URL            string        `json:"url"`
	Domain         string        `json:"domain"`
	EmailDomain    string        `json:"email_domain"`
	Icon           TeamIcon      `json:"icon"`
	AvatarBaseURL  string        `json:"avatar_base_url"`
	IsVerified     bool          `json:"is_verified"`
	PublicURL      string        `json:"public_url"`
	OrgMigrations  OrgMigrations `json:"external_org_migrations"`
	Channels       []string      `json:"channels"`
	Counts         TeamCounts    `json:"counts"`
	DateCreated    int           `json:"date_created"`
	EnterpriseID   string        `json:"enterprise_id"`
	EnterpriseName string        `json:"enterprise_name"`
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
	Image230      string `json:"image_230"`
}

type OrgMigrations struct {
	DateUpdated int           `json:"date_updated"`
	Current     []interface{} `json:"current"`
}

type TeamCounts struct {
	Im      int `json:"im"`
	Mpim    int `json:"mpim"`
	Private int `json:"private"`
	Public  int `json:"public"`
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

type TeamIntegrationLogsInput struct {
	// Filter logs to this Slack app. Defaults to all logs.
	AppID string `json:"app_id"`
	// Filter logs with this change type. Defaults to all logs.
	ChangeType string `json:"change_type"`
	// Number of items to return per page.
	Count string `json:"count"`
	// Page number of results to return.
	Page string `json:"page"`
	// Filter logs to this service. Defaults to all logs.
	ServiceID string `json:"service_id"`
	// Encoded team id to get logs from, required if org token is used
	TeamID string `json:"team_id"`
	// Filter logs generated by this user’s actions. Defaults to all logs.
	User string `json:"user"`
}

// TeamIntegrationLogsResponse defines the JSON-encoded output for TeamIntegrationLogs.
type TeamIntegrationLogsResponse struct {
	Response
	Logs   []IntegrationLog `json:"logs"`
	Paging Paging           `json:"paging"`
}

type IntegrationLog struct {
	ServiceID   int    `json:"service_id"`
	ServiceType string `json:"service_type"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Channel     string `json:"channel"`
	Date        string `json:"date"`
	ChangeType  string `json:"change_type"`
	Scope       string `json:"scope"`
}

// TeamIntegrationLogs is https://api.slack.com/methods/team.integrationLogs
func (s *SlackAPI) TeamIntegrationLogs(input TeamIntegrationLogsInput) TeamIntegrationLogsResponse {
	in := url.Values{}

	if input.AppID != "" {
		in.Add("app_id", input.AppID)
	}

	if input.ChangeType != "" {
		in.Add("change_type", input.ChangeType)
	}

	if input.Count != "" {
		in.Add("count", input.Count)
	}

	if input.Page != "" {
		in.Add("page", input.Page)
	}

	if input.ServiceID != "" {
		in.Add("service_id", input.ServiceID)
	}

	if input.TeamID != "" {
		in.Add("team_id", input.TeamID)
	}

	if input.User != "" {
		in.Add("user", input.User)
	}

	var out TeamIntegrationLogsResponse
	if err := s.baseGET("/api/team.integrationLogs", in, &out); err != nil {
		return TeamIntegrationLogsResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type TeamListExternalInput struct {
	IncludeAllVisible   int `json:"include_all_visible"`
	IncludeApprovedOrgs int `json:"include_approved_orgs"`
}

type TeamListExternalResponse struct {
	Response
	Teams []Team `json:"teams"`
}

// TeamListExternal is https://api.slack.com/methods/team.listExternal
func (s *SlackAPI) TeamListExternal(input TeamListExternalInput) TeamListExternalResponse {
	var out TeamListExternalResponse
	if err := s.basePOST("/api/team.listExternal", input, &out); err != nil {
		return TeamListExternalResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type TeamPreferencesListResponse struct {
	Response
	AllowMessageDeletion bool   `json:"allow_message_deletion"`
	DisplayRealNames     bool   `json:"display_real_names"`
	DisableFileUploads   string `json:"disable_file_uploads"`
	MsgEditWindowMins    int    `json:"msg_edit_window_mins"`
	WhoCanPostGeneral    string `json:"who_can_post_general"`
}

// TeamPreferencesList https://api.slack.com/methods/team.preferences.list
func (s *SlackAPI) TeamPreferencesList() TeamPreferencesListResponse {
	in := url.Values{}
	var out TeamPreferencesListResponse
	if err := s.baseGET("/api/team.preferences.list", in, &out); err != nil {
		return TeamPreferencesListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type TeamProfileResponse struct {
	Response
	Profile TeamProfile `json:"profile"`
}

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

// TeamProfileGet https://api.slack.com/methods/team.profile.get
func (s *SlackAPI) TeamProfileGet() TeamProfileResponse {
	in := url.Values{}
	var out TeamProfileResponse
	if err := s.baseGET("/api/team.profile.get", in, &out); err != nil {
		return TeamProfileResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
