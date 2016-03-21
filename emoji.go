package main

type ResponseEmojiList struct {
	Response
	CacheTs string            `json:"cache_ts"`
	Emoji   map[string]string `json:"emoji"`
}

func (s *SlackAPI) EmojiList() ResponseEmojiList {
	var response ResponseEmojiList
	s.GetRequest(&response, "emoji.list", "token")
	return response
}
