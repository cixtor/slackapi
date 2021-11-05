package slackapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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

// EmojiAdd uploads and registers a new custom emoji.
func (s *SlackAPI) EmojiAdd(emoji string, filename string) Response {
	in := url.Values{
		"name":   {emoji},
		"mode":   {"data"},
		"@image": {filename},
	}
	var out Response
	if err := s.baseFilePOST("/api/emoji.add", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// EmojiAddAlias creates an alias for an existing emoji.
func (s *SlackAPI) EmojiAddAlias(emoji string, alias string) Response {
	in := url.Values{
		"alias_for": {emoji},
		"name":      {alias},
		"mode":      {"alias"},
	}
	var out Response
	if err := s.baseFormPOST("/api/emoji.add", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
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

type EmojiListWithCategoriesResponse struct {
	EmojiListResponse
	CategoriesVersion string          `json:"categories_version"`
	Categories        []EmojiCategory `json:"categories"`
}

type EmojiCategory struct {
	Name       string   `json:"name"`
	EmojiNames []string `json:"emoji_names"`
}

// EmojiListWithCategories lists custom emoji for a team with categories.
func (s *SlackAPI) EmojiListWithCategories() EmojiListWithCategoriesResponse {
	in := url.Values{"include_categories": {"true"}}
	var out EmojiListWithCategoriesResponse
	if err := s.baseGET("/api/emoji.list", in, &out); err != nil {
		out = EmojiListWithCategoriesResponse{}
		out.Error = err.Error()
		return out
	}
	return out
}

type EmojiAdminListInput struct {
	Page    int
	Count   int
	Queries []string
	UserIDs []string
	SortBy  string
	SortDir string
}

type EmojiAdminListResponse struct {
	Response
	CustomEmojiTotalCount int     `json:"custom_emoji_total_count"`
	DisabledEmoji         []Emoji `json:"disabled_emoji"`
	Emoji                 []Emoji `json:"emoji"`
	Paging                Paging  `json:"paging"`
}

// EmojiAdminList lists custom emoji for a team as an administrator.
func (s *SlackAPI) EmojiAdminList(input EmojiAdminListInput) EmojiAdminListResponse {
	in := url.Values{}
	if input.Page > 0 {
		in.Add("page", fmt.Sprintf("%d", input.Page))
	}
	if input.Count > 0 {
		in.Add("count", fmt.Sprintf("%d", input.Count))
	}
	if len(input.Queries) > 0 {
		b, _ := json.Marshal(input.Queries)
		in.Add("queries", string(b))
	}
	if len(input.UserIDs) > 0 {
		b, _ := json.Marshal(input.UserIDs)
		in.Add("user_ids", string(b))
	}
	if input.SortBy != "" {
		in.Add("sort_by", input.SortBy)
	}
	if input.SortDir != "" {
		in.Add("sort_dir", input.SortDir)
	}

	var out EmojiAdminListResponse
	if err := s.baseGET("/api/emoji.adminList", in, &out); err != nil {
		return EmojiAdminListResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// EmojiRemove deletes a custom emoji for a team as an administrator.
func (s *SlackAPI) EmojiRemove(name string) Response {
	in := url.Values{"name": {name}}
	var out Response
	if err := s.baseFormPOST("/api/emoji.remove", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

type EmojiInfoResponse struct {
	Response
	Emoji
}

// EmojiGetInfo returns information about an existing emoji.
func (s *SlackAPI) EmojiGetInfo(name string) EmojiInfoResponse {
	in := url.Values{"name": {name}}
	var out EmojiInfoResponse
	if err := s.baseFormPOST("/api/emoji.getInfo", in, &out); err != nil {
		return EmojiInfoResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
