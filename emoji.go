package main

// ResponseEmojiList defines the JSON-encoded output for EmojiList.
type ResponseEmojiList struct {
	Response
	CacheTs string            `json:"cache_ts"`
	Emoji   map[string]string `json:"emoji"`
}

// EmojiList lists custom emoji for a team.
func (s *SlackAPI) EmojiList() ResponseEmojiList {
	var response ResponseEmojiList
	s.GetRequest(&response, "emoji.list", "token")
	return response
}
