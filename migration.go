package slackapi

type MigrationExchangeResponse struct {
	Response
	TeamID         string            `json:"team_id"`
	EnterpriseID   string            `json:"enterprise_id"`
	UserIDMap      map[string]string `json:"user_id_map"`
	InvalidUserIds []string          `json:"invalid_user_ids"`
}

// MigrationExchange for Enterprise Grid workspaces, map local user IDs to global user IDs.
func (s *SlackAPI) MigrationExchange(users []string, order bool) MigrationExchangeResponse {
	in := struct {
		Users []string `json:"users"`
		ToOld bool     `json:"to_old"`
	}{
		Users: users,
		ToOld: order,
	}
	var out MigrationExchangeResponse
	if err := s.baseJSONPOST("/api/migration.exchange", in, &out); err != nil {
		return MigrationExchangeResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
