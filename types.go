package main

type Base struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type Session struct {
	Base
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

type Users struct {
	Base
	Members []User `json:"members"`
}

type Rooms struct {
	Base
	Channels []Room `json:"channels"`
}

type Groups struct {
	Base
	Channels []Room `json:"groups"`
}

type User struct {
	Id                string  `json:"id"`
	TeamId            string  `json:"team_id"`
	Name              string  `json:"name"`
	Deleted           bool    `json:"deleted"`
	Status            string  `json:"status"`
	Color             string  `json:"color"`
	RealName          string  `json:"real_name"`
	Tz                string  `json:"tz"`
	TzLabel           string  `json:"tz_label"`
	TzOffset          int     `json:"tz_offset"`
	Profile           Profile `json:"profile"`
	IsAdmin           bool    `json:"is_admin"`
	IsOwner           bool    `json:"is_owner"`
	IsPrimaryOwner    bool    `json:"is_primary_owner"`
	IsRestricted      bool    `json:"is_restricted"`
	IsUltraRestricted bool    `json:"is_ultra_restricted"`
	IsBot             bool    `json:"is_bot"`
	Presence          string  `json:"presence"`
}

type Profile struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Title              string `json:"title"`
	Skype              string `json:"skype"`
	Phone              string `json:"phone"`
	Image24            string `json:"image_24"`
	Image32            string `json:"image_32"`
	Image48            string `json:"image_48"`
	Image72            string `json:"image_72"`
	Image192           string `json:"image_192"`
	ImageOriginal      string `json:"image_original"`
	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`
	Email              string `json:"email"`
}

type ChannelEvent struct {
	Base
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

type Message struct {
	ChannelEvent
	Message struct {
		Subtype  string `json:"subtype"`
		Text     string `json:"text"`
		Ts       string `json:"ts"`
		Type     string `json:"type"`
		Username string `json:"username"`
	} `json:"message"`
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
