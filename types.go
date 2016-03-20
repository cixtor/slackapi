package main

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type Session struct {
	Response
	AlreadyOpen bool `json:"already_open"`
	NoOp        bool `json:"no_op"`
	Channel     struct {
		Id string `json:"id"`
	} `json:"channel"`
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

type Message struct {
	Attachments  []Attachment `json:"attachments"`
	BotId        string       `json:"bot_id"`
	Channel      string       `json:"channel"`
	DisplayAsBot bool         `json:"display_as_bot"`
	File         File         `json:"file"`
	Ok           bool         `json:"ok"`
	Subtype      string       `json:"subtype"`
	Text         string       `json:"text"`
	Ts           string       `json:"ts"`
	Type         string       `json:"type"`
	Upload       bool         `json:"upload"`
	User         string       `json:"user"`
	Username     string       `json:"username"`
}

type Attachment struct {
	Fallback    string `json:"fallback"`
	FromUrl     string `json:"from_url"`
	Id          int    `json:"id"`
	ImageBytes  int    `json:"image_bytes"`
	ImageHeight int    `json:"image_height"`
	ImageUrl    string `json:"image_url"`
	ImageWidth  int    `json:"image_width"`
	ServiceName string `json:"service_name"`
	Text        string `json:"text"`
	ThumbHeight int    `json:"thumb_height"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	Title       string `json:"title"`
	TitleLink   string `json:"title_link"`
	Ts          int    `json:"ts"`
}

type DeletedMessage struct {
	Deleted bool
	Text    string
	Ts      string
}

type Owner struct {
	Ok     bool   `json:"ok"`
	Team   string `json:"team"`
	TeamId string `json:"team_id"`
	Url    string `json:"url"`
	User   string `json:"user"`
	UserId string `json:"user_id"`
}

type Rooms struct {
	Response
	Channels []Room `json:"channels"`
}

type Groups struct {
	Response
	Channels []Room `json:"groups"`
}

type ChannelEvent struct {
	Response
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

type Room struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	Created    int      `json:"created"`
	IsArchived bool     `json:"is_archived"`
	IsChannel  bool     `json:"is_channel"`
	IsGeneral  bool     `json:"is_general"`
	IsMember   bool     `json:"is_member"`
	NumMembers int      `json:"num_members"`
	Members    []string `json:"members"`
	Purpose    struct {
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
		Value   string `json:"value"`
	} `json:"purpose"`
	Topic struct {
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
		Value   string `json:"value"`
	} `json:"topic"`
}
