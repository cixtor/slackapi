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

type SlackbotResponse struct {
	Response
	AutoResponse AutoResponse `json:"response"`
}

type AutoResponse struct {
	ID        string   `json:"id"`
	Creator   string   `json:"creator"`
	Created   int      `json:"created"`
	Responses []string `json:"responses"`
	Triggers  []string `json:"triggers"`
}

// SlackbotResponsesAdd instructs Slackbot to automatically respond to messages
// that members of your workspace send in channels. By default, all members can
// edit Slackbot responses. You can change this in admin settings.
//
// Source: https://cixtor.slack.com/customize/slackbot
func (s *SlackAPI) SlackbotResponsesAdd(triggers string, responses string) SlackbotResponse {
	in := url.Values{
		"triggers":  {triggers},
		"responses": {responses},
	}
	var out SlackbotResponse
	if err := s.baseFormPOST("/api/slackbot.responses.add", in, &out); err != nil {
		return SlackbotResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

func (s *SlackAPI) SlackbotResponsesEdit(id string, triggers string, responses string) SlackbotResponse {
	in := url.Values{
		"response":  {id},
		"triggers":  {triggers},
		"responses": {responses},
	}
	var out SlackbotResponse
	if err := s.baseFormPOST("/api/slackbot.responses.edit", in, &out); err != nil {
		return SlackbotResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
