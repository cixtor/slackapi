package slackapi

// ResponseEmojiList defines the JSON-encoded output for EmojiList.
type ResponseEmojiList struct {
	Response
	CacheTimestamp string            `json:"cache_ts"`
	Emoji          map[string]string `json:"emoji"`
}

// EmojiList lists custom emoji for a team.
func (s *SlackAPI) EmojiList() ResponseEmojiList {
	var response ResponseEmojiList
	s.getRequest(&response, "emoji.list", nil)
	return response
}
