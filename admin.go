package slackapi

import (
	"fmt"
	"net/url"
)

type ChannelAnalyticsInput struct {
	Count int
	// Example: 30d or YYYY-MM
	DateInterval string
	// Date in YYYY-MM-dd format.
	StartDate string
	// Date in YYYY-MM-dd format.
	EndDate string
	// Either public or private.
	Privacy string
	Query   string
	// Default: name
	SortColumn string
	// Either asc or desc.
	SortDirection string
}

type ChannelAnalyticsResponse struct {
	Response
	NumFound       int                `json:"num_found"`
	AllAnalytics   []ChannelAnalytics `json:"channel_analytics"`
	Channels       []Channel          `json:"channels"`
	NextCursorMark string             `json:"next_cursor_mark"`
}

type ChannelAnalytics struct {
	ChannelID            string `json:"channel_id"`
	TeamID               string `json:"team_id"`
	DateCreate           int    `json:"date_create"`
	IsShared             int    `json:"is_shared"`
	Name                 string `json:"name"`
	Topic                string `json:"topic"`
	Purpose              string `json:"purpose"`
	TotalMembersCount    int    `json:"total_members_count"`
	FullMembersCount     int    `json:"full_members_count"`
	GuestMembersCount    int    `json:"guest_members_count"`
	ChatsCount           int    `json:"chats_count"`
	MessagesCount        int    `json:"messages_count"`
	ReactionsCount       int    `json:"reactions_count"`
	UsersWhoReactedCount int    `json:"users_who_reacted_count"`
	WritersCount         int    `json:"writers_count"`
	IsPrivate            bool   `json:"is_private"`
	LastMessagePosted    int    `json:"last_message_posted"`
	ReadersCount         int    `json:"readers_count"`
}

// ChannelAnalytics is https://api.slack.com/methods/admin.analytics.getChannelAnalytics
func (s *SlackAPI) ChannelAnalytics(input ChannelAnalyticsInput) ChannelAnalyticsResponse {
	in := url.Values{}
	if input.Count > 0 {
		in.Add("count", fmt.Sprintf("%d", input.Count))
	}
	if input.DateInterval == "" {
		in.Add("date_interval", "30d")
	} else {
		in.Add("date_interval", input.DateInterval)
	}
	if input.StartDate != "" {
		in.Add("start_date", input.StartDate)
	}
	if input.EndDate != "" {
		in.Add("end_date", input.EndDate)
	}
	if input.Privacy != "" {
		in.Add("privacy", input.Privacy)
	}
	if input.Query != "" {
		in.Add("query", input.Query)
	}
	if input.SortColumn == "" {
		in.Add("sort_column", "name")
	} else {
		in.Add("sort_column", input.SortColumn)
	}
	if input.SortDirection == "desc" {
		in.Add("sort_direction", "desc")
	} else {
		in.Add("sort_direction", "asc")
	}
	var out ChannelAnalyticsResponse
	if err := s.baseFormPOST("/api/admin.analytics.getChannelAnalytics", in, &out); err != nil {
		return ChannelAnalyticsResponse{Response: Response{Error: err.Error()}}
	}
	return out
}
