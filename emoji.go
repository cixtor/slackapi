package slackapi

// Emoji contains the data associated to a single emoji.
type Emoji struct {
	AliasFor        string   `json:"alias_for"`
	AvatarHash      string   `json:"avatar_hash"`
	CanDelete       bool     `json:"can_delete"`
	Created         int      `json:"created"`
	IsAlias         int      `json:"is_alias"`
	IsBad           bool     `json:"is_bad"`
	Name            string   `json:"name"`
	Synonyms        []string `json:"synonyms"`
	TeamID          string   `json:"team_id"`
	URL             string   `json:"url"`
	UserDisplayName string   `json:"user_display_name"`
	UserID          string   `json:"user_id"`
}

type EmojiListInput struct {
	IncludeCategories bool `json:"include_categories"`
}

type EmojiListResponse struct {
	Response
	CacheTimestamp string            `json:"cache_ts"`
	Emoji          map[string]string `json:"emoji"`
}

// EmojiList lists custom emoji for a team.
func (s *SlackAPI) EmojiList() EmojiListResponse {
	var out EmojiListResponse
	if err := s.baseGET("/api/emoji.list", nil, &out); err != nil {
		return EmojiListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
