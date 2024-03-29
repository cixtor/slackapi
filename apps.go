package slackapi

type AppsListResponse struct {
	Response
	Apps           []AppInfo `json:"apps"`
	CacheTimestamp string    `json:"cache_ts"`
}

type AppInfo struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Icons AppIcons `json:"icons"`
}

type AppIcons struct {
	Image1024 string `json:"image_1024"`
	Image128  string `json:"image_128"`
	Image192  string `json:"image_192"`
	Image32   string `json:"image_32"`
	Image36   string `json:"image_36"`
	Image48   string `json:"image_48"`
	Image512  string `json:"image_512"`
	Image64   string `json:"image_64"`
	Image72   string `json:"image_72"`
	Image96   string `json:"image_96"`
}

// AppsList lists associated applications.
func (s *SlackAPI) AppsList() AppsListResponse {
	var out AppsListResponse
	if err := s.baseGET("/api/apps.list", nil, &out); err != nil {
		return AppsListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type AppsConnectionsOpenResponse struct {
	Response
	URL string `json:"url"`
}

// AppsConnectionsOpen is apps.connections.open
func (s *SlackAPI) AppsConnectionsOpen() AppsConnectionsOpenResponse {
	in := struct{}{}
	var out AppsConnectionsOpenResponse
	if err := s.baseJSONPOST("/api/apps.connections.open", in, &out); err != nil {
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
	if err := s.baseJSONPOST("/api/apps.event.authorizations.list", input, &out); err != nil {
		return AppsEventAuthorizationsListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type AppsManifestCreateResponse struct {
	Response
	AppID             string         `json:"app_id"`
	Credentials       AppCredentials `json:"credentials"`
	OauthAuthorizeURL string         `json:"oauth_authorize_url"`
}

type AppCredentials struct {
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
	if err := s.baseJSONPOST("/api/apps.manifest.create", in, &out); err != nil {
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
	if err := s.baseJSONPOST("/api/apps.manifest.delete", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

type AppsManifestExportResponse struct {
	Response
	Manifest AppManifest `json:"manifest"`
}

type AppManifest struct {
	Metadata           AppManifestMetadata           `json:"_metadata"`
	DisplayInformation AppManifestDisplayInformation `json:"display_information"`
	Features           AppManifestFeatures           `json:"features"`
	OauthConfig        AppManifestOauthConfig        `json:"oauth_config"`
	Settings           AppManifestSettings           `json:"settings"`
}

type AppManifestMetadata struct {
	MajorVersion int `json:"major_version"`
	MinorVersion int `json:"minor_version"`
}

type AppManifestDisplayInformation struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	BackgroundColor string `json:"background_color"`
	LongDescription string `json:"long_description"`
}

type AppManifestFeatures struct {
	AppHome       AppManifestFeaturesAppHome `json:"app_home"`
	BotUser       AppManifestFeaturesBotUser `json:"bot_user"`
	SlashCommands []SlashCommand             `json:"slash_commands"`
	WorkflowSteps []WorkflowStep             `json:"workflow_steps"`
}

type AppManifestFeaturesAppHome struct {
	HomeTabEnabled             bool `json:"home_tab_enabled"`
	MessagesTabEnabled         bool `json:"messages_tab_enabled"`
	MessagesTabReadOnlyEnabled bool `json:"messages_tab_read_only_enabled"`
}

type AppManifestFeaturesBotUser struct {
	DisplayName  string `json:"display_name"`
	AlwaysOnline bool   `json:"always_online"`
}

type SlashCommand struct {
	Command      string `json:"command"`
	Description  string `json:"description"`
	UsageHint    string `json:"usage_hint"`
	ShouldEscape bool   `json:"should_escape"`
}

type WorkflowStep struct {
	Name       string `json:"name"`
	CallbackID string `json:"callback_id"`
}

type AppManifestOauthConfig struct {
	RedirectURLs []string            `json:"redirect_urls"`
	Scopes       map[string][]string `json:"scopes"`
}

type AppManifestSettings struct {
	EventSubscriptions map[string][]string `json:"event_subscriptions"`
	Interactivity      struct {
		IsEnabled bool `json:"is_enabled"`
	} `json:"interactivity"`
	OrgDeployEnabled     bool `json:"org_deploy_enabled"`
	SocketModeEnabled    bool `json:"socket_mode_enabled"`
	IsHosted             bool `json:"is_hosted"`
	TokenRotationEnabled bool `json:"token_rotation_enabled"`
}

// AppsManifestExport is https://api.slack.com/methods/apps.manifest.export
func (s *SlackAPI) AppsManifestExport(appID string) AppsManifestExportResponse {
	in := struct {
		AppID string `json:"app_id"`
	}{
		AppID: appID,
	}
	var out AppsManifestExportResponse
	if err := s.baseJSONPOST("/api/apps.manifest.export", in, &out); err != nil {
		return AppsManifestExportResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type AppsManifestUpdateResponse struct {
	Response
	AppID              string `json:"app_id"`
	PermissionsUpdated bool   `json:"permissions_updated"`
}

// AppsManifestUpdate is https://api.slack.com/methods/apps.manifest.update
func (s *SlackAPI) AppsManifestUpdate(appID string, manifest string) AppsManifestUpdateResponse {
	in := struct {
		AppID    string `json:"app_id"`
		Manifest string `json:"manifest"`
	}{
		AppID:    appID,
		Manifest: manifest,
	}
	var out AppsManifestUpdateResponse
	if err := s.baseJSONPOST("/api/apps.manifest.update", in, &out); err != nil {
		return AppsManifestUpdateResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// AppsManifestValidate is https://api.slack.com/methods/apps.manifest.validate
func (s *SlackAPI) AppsManifestValidate(manifest string, appID string) Response {
	in := struct {
		Manifest string `json:"manifest"`
		AppID    string `json:"app_id"`
	}{
		Manifest: manifest,
		AppID:    appID,
	}
	var out Response
	if err := s.baseJSONPOST("/api/apps.manifest.validate", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
