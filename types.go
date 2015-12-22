package main

type Users struct {
	Ok      bool   `json:"ok"`
	Members []User `json:"members"`
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
