package main

type BotEntity struct {
	ID      string            `json:"id"`
	Deleted bool              `json:"deleted"`
	Name    string            `json:"name"`
	Icons   map[string]string `json:"icons"`
}

type ResponseBot struct {
	Response
	Bot BotEntity `json:"bot"`
}

func (s *SlackAPI) BotsInfo(bot string) ResponseBot {
	var response ResponseBot
	s.GetRequest(&response, "bots.info", "token", "bot="+s.UsersId(bot))
	return response
}
