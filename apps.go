package slackapi

type AppsConnectionsOpenResponse struct {
	Response
	URL string `json:"url"`
}

// AppsConnectionsOpen is apps.connections.open
func (s *SlackAPI) AppsConnectionsOpen() AppsConnectionsOpenResponse {
	in := struct{}{}
	var out AppsConnectionsOpenResponse
	if err := s.basePOST("/api/apps.connections.open", in, &out); err != nil {
		return AppsConnectionsOpenResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type AppsEventAuthorizationsListInput struct {
	EventContext string `json:"event_context"`
	Cursor       string `json:"cursor"`
	Limit        int    `json:"limit"`
}

type AppsEventAuthorizationsListResponse struct {
	Response
	Authorizations []AppAuthorization `json:"authorizations"`
}

type AppAuthorization struct {
	EnterpriseID string `json:"enterprise_id"`
	TeamID       string `json:"team_id"`
	UserID       string `json:"user_id"`
	IsBot        string `json:"is_bot"`
}

// AppsEventAuthorizationsList is https://api.slack.com/methods/apps.event.authorizations.list
func (s *SlackAPI) AppsEventAuthorizationsList(input AppsEventAuthorizationsListInput) AppsEventAuthorizationsListResponse {
	var out AppsEventAuthorizationsListResponse
	if err := s.basePOST("/api/apps.event.authorizations.list", input, &out); err != nil {
		return AppsEventAuthorizationsListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type AppsManifestCreateResponse struct {
	Response
	AppID             string                  `json:"app_id"`
	Credentials       AppsManifestCredentials `json:"credentials"`
	OauthAuthorizeURL string                  `json:"oauth_authorize_url"`
}

type AppsManifestCredentials struct {
	ClientID          string `json:"client_id"`
	ClientSecret      string `json:"client_secret"`
	VerificationToken string `json:"verification_token"`
	SigningSecret     string `json:"signing_secret"`
}

// AppsManifestCreate is https://api.slack.com/methods/apps.manifest.create
func (s *SlackAPI) AppsManifestCreate(manifest string) AppsManifestCreateResponse {
	in := struct {
		Manifest string `json:"manifest"`
	}{
		Manifest: manifest,
	}
	var out AppsManifestCreateResponse
	if err := s.basePOST("/api/apps.manifest.create", in, &out); err != nil {
		return AppsManifestCreateResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// AppsManifestDelete is https://api.slack.com/methods/apps.manifest.delete
func (s *SlackAPI) AppsManifestDelete(appID string) Response {
	in := struct {
		AppID string `json:"app_id"`
	}{
		AppID: appID,
	}
	var out Response
	if err := s.basePOST("/api/apps.manifest.delete", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
