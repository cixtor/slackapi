package main

type Base struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

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

type Owner struct {
	Ok     bool   `json:"ok"`
	Team   string `json:"team"`
	TeamId string `json:"team_id"`
	Url    string `json:"url"`
	User   string `json:"user"`
	UserId string `json:"user_id"`
}

type Rooms struct {
	Base
	Channels []Room `json:"channels"`
}

type Groups struct {
	Base
	Channels []Room `json:"groups"`
}

type ChannelEvent struct {
	Base
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

type Message struct {
	ChannelEvent
	Message MessageNode `json:"message"`
}

type History struct {
	Base
	HasMore            bool          `json:"has_more"`
	Messages           []MessageNode `json:"messages"`
	UnreadCountDisplay int           `json:"unread_count_display"`
}

type UserHistory struct {
	Filtered int
	Latest   string
	Messages []MessageNode
	Oldest   string
	Total    int
	Username string
}

type MessageNode struct {
	Subtype  string `json:"subtype"`
	Text     string `json:"text"`
	Ts       string `json:"ts"`
	Type     string `json:"type"`
	User     string `json:"user"`
	Username string `json:"username"`
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
