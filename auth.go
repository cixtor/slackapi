package slackapi

import (
	"net/url"
	"strconv"
)

// ResponseRevocation defines the JSON-encoded output for Revocation.
type ResponseRevocation struct {
	Response
	Revoked bool `json:"revoked"`
}

// APITest checks API calling code.
func (s *SlackAPI) APITest() Response {
	var response Response
	s.getRequest(&response, "api.test", nil)
	return response
}

// AuthRevoke revokes a token.
func (s *SlackAPI) AuthRevoke() ResponseRevocation {
	var response ResponseRevocation
	s.getRequest(&response, "auth.revoke", nil)
	return response
}

type AuthTeamsListInput struct {
	Cursor      string `json:"cursor"`
	IncludeIcon bool   `json:"include_icon"`
	Limit       int    `json:"limit"`
}

type AuthTeamsListResponse struct {
	Response
	Teams            []Team           `json:"teams"`
	ResponseMetadata ResponseMetadata `json:"response_metadata"`
}

// AuthTeamsList is https://api.slack.com/methods/auth.teams.list
func (s *SlackAPI) AuthTeamsList(input AuthTeamsListInput) AuthTeamsListResponse {
	in := url.Values{}

	if input.Cursor != "" {
		in.Add("cursor", input.Cursor)
	}

	if input.IncludeIcon {
		in.Add("include_icon", "true")
	}

	if input.Limit > 0 {
		in.Add("limit", strconv.Itoa(input.Limit))
	}

	var out AuthTeamsListResponse
	if err := s.baseGET("/api/auth.teams.list", in, &out); err != nil {
		return AuthTeamsListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// AuthTest checks authentication and identity.
func (s *SlackAPI) AuthTest() (*Owner, error) {
	var output Owner
	input := url.Values{}
	if err := s.baseGET("/api/auth.test", input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
