package main

import (
	"encoding/json"
)

type Response struct {
	Ok       bool   `json:"ok"`
	Error    string `json:"error,omitempty"`
	Needed   string `json:"needed,omitempty"`
	Provided string `json:"provided,omitempty"`
}

type Session struct {
	Response
	AlreadyOpen bool `json:"already_open"`
	NoOp        bool `json:"no_op"`
	Channel     struct {
		Id string `json:"id"`
	} `json:"channel"`
}

type Owner struct {
	Response
	Team   string `json:"team"`
	TeamId string `json:"team_id"`
	Url    string `json:"url"`
	User   string `json:"user"`
	UserId string `json:"user_id"`
}

type History struct {
	Response
	HasMore            bool      `json:"has_more"`
	Messages           []Message `json:"messages"`
	UnreadCountDisplay int       `json:"unread_count_display"`
}

type MyHistory struct {
	Filtered int
	Latest   string
	Messages []Message
	Oldest   string
	Total    int
	Username string
}

type DeletedHistory struct {
	Deleted    int
	NotDeleted int
	Messages   []DeletedMessage
}

type Post struct {
	Response
	Channel string  `json:"channel"`
	Message Message `json:"message"`
	Text    string  `json:"text"`
	Ts      string  `json:"ts"`
}

type Message struct {
	Attachments  []Attachment `json:"attachments"`
	BotId        string       `json:"bot_id"`
	Channel      string       `json:"channel"`
	DisplayAsBot bool         `json:"display_as_bot"`
	File         File         `json:"file"`
	Subtype      string       `json:"subtype"`
	Text         string       `json:"text"`
	Ts           string       `json:"ts"`
	Type         string       `json:"type"`
	Upload       bool         `json:"upload"`
	User         string       `json:"user"`
	Username     string       `json:"username"`
}

type ModifiedMessage struct {
	Response
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Ts      string `json:"ts"`
}

type Attachment struct {
	Fallback    string      `json:"fallback"`
	FromUrl     string      `json:"from_url"`
	Id          int         `json:"id"`
	ImageBytes  int         `json:"image_bytes"`
	ImageHeight int         `json:"image_height"`
	ImageUrl    string      `json:"image_url"`
	ImageWidth  int         `json:"image_width"`
	ServiceName string      `json:"service_name"`
	Text        string      `json:"text"`
	ThumbHeight int         `json:"thumb_height"`
	ThumbUrl    string      `json:"thumb_url"`
	ThumbWidth  int         `json:"thumb_width"`
	Title       string      `json:"title"`
	TitleLink   string      `json:"title_link"`
	Ts          json.Number `json:"ts"`
}

type DeletedMessage struct {
	Deleted bool
	Text    string
	Ts      string
}

type Channel struct {
	Created            int            `json:"created"`
	Creator            string         `json:"creator"`
	Id                 string         `json:"id"`
	IsArchived         bool           `json:"is_archived"`
	IsChannel          bool           `json:"is_channel"`
	IsGeneral          bool           `json:"is_general"`
	IsGroup            bool           `json:"is_group"`
	IsMember           bool           `json:"is_member"`
	IsMpim             bool           `json:"is_mpim"`
	IsOpen             bool           `json:"is_open"`
	LastRead           string         `json:"last_read"`
	Latest             ChannelLatest  `json:"latest"`
	Members            []string       `json:"members"`
	Name               string         `json:"name"`
	NumMembers         int            `json:"num_members"`
	Purpose            ChannelPurpose `json:"purpose"`
	Topic              ChannelTopic   `json:"topic"`
	UnreadCount        int            `json:"unread_count"`
	UnreadCountDisplay int            `json:"unread_count_display"`
}

type ChannelLatest struct {
	Text string `json:"text"`
	Ts   string `json:"ts"`
	Type string `json:"type"`
	User string `json:"user"`
}

type ChannelPurpose struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

type ChannelRename struct {
	Response
	Channel struct {
		ID        string `json:"id"`
		IsChannel bool   `json:"is_channel"`
		IsGroup   bool   `json:"is_group"`
		Name      string `json:"name"`
		Created   int    `json:"created"`
	} `json:"channel"`
}

type ChannelTopic struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

type ChannelPurposeNow struct {
	Response
	Purpose string `json:"purpose"`
}

type ChannelTopicNow struct {
	Response
	Topic string `json:"topic"`
}

type AppsList struct {
	Response
	Apps    []AppsListApps `json:"apps"`
	CacheTs string         `json:"cache_ts"`
}

type AppsListApps struct {
	Id    string            `json:"id"`
	Name  string            `json:"name"`
	Icons AppsListAppsIcons `json:"icons"`
}

type AppsListAppsIcons struct {
	Image1024 string `json:"image_1024"`
	Image128  string `json:"image_128"`
	Image192  string `json:"image_192"`
	Image32   string `json:"image_32"`
	Image36   string `json:"image_36"`
	Image48   string `json:"image_48"`
	Image512  string `json:"image_512"`
	Image64   string `json:"image_64"`
	Image72   string `json:"image_72"`
	Image96   string `json:"image_96"`
}

type Pagination struct {
	Count int `json:"count"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}
