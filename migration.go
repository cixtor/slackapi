package slackapi

// ResponseMigrationExchange defines the expected data from the JSON-encoded API response.
type ResponseMigrationExchange struct {
	Response
	TeamID         string            `json:"team_id"`
	EnterpriseID   string            `json:"enterprise_id"`
	UserIDMap      map[string]string `json:"user_id_map"`
	InvalidUserIds []string          `json:"invalid_user_ids"`
}

// MigrationExchange for Enterprise Grid workspaces, map local user IDs to global user IDs.
func (s *SlackAPI) MigrationExchange(users []string, order bool) ResponseMigrationExchange {
	var response ResponseMigrationExchange
	s.getRequest(&response, "migration.exchange", struct {
		Users []string `json:"users"`
		ToOld bool     `json:"to_old"`
	}{users, order})
	return response
}
