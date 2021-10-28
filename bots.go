package slackapi

import (
	"net/url"
)

type BotResponse struct {
	Response
	Bot Bot `json:"bot"`
}

type Bot struct {
	ID      string            `json:"id"`
	Deleted bool              `json:"deleted"`
	Name    string            `json:"name"`
	Icons   map[string]string `json:"icons"`
}

// BotsInfo gets information about a bot user.
func (s *SlackAPI) BotsInfo(bot string) BotResponse {
	in := url.Values{"bot": {bot}}
	var out BotResponse
	if err := s.baseGET("/api/bots.info", in, &out); err != nil {
		return BotResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
