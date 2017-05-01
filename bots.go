package main

// ResponseBot defines the JSON-encoded output for Bot.
type ResponseBot struct {
	Response
	Bot BotEntity `json:"bot"`
}

// BotEntity defines the expected data from the JSON-encoded API response.
type BotEntity struct {
	ID      string            `json:"id"`
	Deleted bool              `json:"deleted"`
	Name    string            `json:"name"`
	Icons   map[string]string `json:"icons"`
}

// BotsInfo gets information about a bot user.
func (s *SlackAPI) BotsInfo(bot string) ResponseBot {
	var response ResponseBot
	s.GetRequest(&response, "bots.info", "token", "bot="+s.UsersID(bot))
	return response
}
