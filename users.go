package main

import (
	"strings"
)

type ResponseUsersInfo struct {
	Response
	User UserData `json:"user"`
}

type ResponseUsersGetPresence struct {
	Response
	UserPresence
}

type ResponseUsersList struct {
	Response
	Members []UserData `json:"members"`
}

type UserPresence struct {
	AutoAway        bool   `json:"auto_away"`
	ConnectionCount int    `json:"connection_count"`
	LastActivity    int    `json:"last_activity"`
	ManualAway      bool   `json:"manual_away"`
	Online          bool   `json:"online"`
	Presence        string `json:"presence"`
}

type UserData struct {
	Color             string      `json:"color"`
	Deleted           bool        `json:"deleted"`
	Has2fa            bool        `json:"has_2fa"`
	Id                string      `json:"id"`
	IsAdmin           bool        `json:"is_admin"`
	IsBot             bool        `json:"is_bot"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
	Name              string      `json:"name"`
	Presence          string      `json:"presence"`
	Profile           UserProfile `json:"profile"`
	RealName          string      `json:"real_name"`
	Status            string      `json:"status"`
	TeamId            string      `json:"team_id"`
	TwoFactorType     string      `json:"two_factor_type"`
	Tz                string      `json:"tz"`
	TzLabel           string      `json:"tz_label"`
	TzOffset          int         `json:"tz_offset"`
}

type UserProfile struct {
	ApiAppID           string      `json:"api_app_id"`
	BotID              string      `json:"bot_id"`
	AvatarHash         string      `json:"avatar_hash"`
	Email              string      `json:"email"`
	Fields             interface{} `json:"fields"`
	FirstName          string      `json:"first_name"`
	Image1024          string      `json:"image_1024"`
	Image192           string      `json:"image_192"`
	Image24            string      `json:"image_24"`
	Image32            string      `json:"image_32"`
	Image48            string      `json:"image_48"`
	Image512           string      `json:"image_512"`
	Image72            string      `json:"image_72"`
	ImageOriginal      string      `json:"image_original"`
	LastName           string      `json:"last_name"`
	Phone              string      `json:"phone"`
	RealName           string      `json:"real_name"`
	RealNameNormalized string      `json:"real_name_normalized"`
	StatusText         string      `json:"status_text"`
	StatusEmoji        string      `json:"status_emoji"`
	Skype              string      `json:"skype"`
	Title              string      `json:"title"`
}

type ResponseUserIdentity struct {
	Response
	Profile UserProfile `json:"profile"`
}

type ResponseUsersCounts struct {
	Response
	Channels []CountChannel `json:"channels"`
	Groups   []CountGroup   `json:"groups"`
	Ims      []CountIm      `json:"ims"`
}

type CountChannel struct {
	ID                  string `json:"id"`
	IsArchived          bool   `json:"is_archived"`
	IsGeneral           bool   `json:"is_general"`
	IsMember            bool   `json:"is_member"`
	IsMuted             bool   `json:"is_muted"`
	IsStarred           bool   `json:"is_starred"`
	MentionCount        int    `json:"mention_count"`
	MentionCountDisplay int    `json:"mention_count_display"`
	Name                string `json:"name"`
	NameNormalized      string `json:"name_normalized"`
	UnreadCount         int    `json:"unread_count"`
	UnreadCountDisplay  int    `json:"unread_count_display"`
}

type CountGroup struct {
	ID                  string `json:"id"`
	IsArchived          bool   `json:"is_archived"`
	IsMpim              bool   `json:"is_mpim"`
	IsMuted             bool   `json:"is_muted"`
	IsOpen              bool   `json:"is_open"`
	IsStarred           bool   `json:"is_starred"`
	MentionCount        int    `json:"mention_count"`
	MentionCountDisplay int    `json:"mention_count_display"`
	Name                string `json:"name"`
	NameNormalized      string `json:"name_normalized"`
	UnreadCount         int    `json:"unread_count"`
	UnreadCountDisplay  int    `json:"unread_count_display"`
}

type CountIm struct {
	DmCount   int    `json:"dm_count"`
	ID        string `json:"id"`
	IsOpen    bool   `json:"is_open"`
	IsStarred bool   `json:"is_starred"`
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
}

func (s *SlackAPI) UsersCounts() ResponseUsersCounts {
	var response ResponseUsersCounts
	s.GetRequest(&response, "users.counts", "token")
	return response
}

func (s *SlackAPI) UsersGetPresence(query string) ResponseUsersGetPresence {
	var response ResponseUsersGetPresence
	s.GetRequest(&response, "users.getPresence", "token", "user="+query)
	return response
}

func (s *SlackAPI) UsersId(query string) string {
	response := s.UsersList()

	if response.Ok {
		for _, user := range response.Members {
			if user.Name == query {
				return user.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) UsersInfo(query string) ResponseUsersInfo {
	query = s.UsersId(query)
	var response ResponseUsersInfo
	s.GetRequest(&response, "users.info", "token", "user="+query)
	return response
}

func (s *SlackAPI) UsersList() ResponseUsersList {
	if s.TeamUsers.Ok == true {
		return s.TeamUsers
	}

	var response ResponseUsersList
	s.GetRequest(&response, "users.list", "token", "presence=1")
	s.TeamUsers = response

	return response
}

func (s *SlackAPI) UsersSearch(query string) []UserData {
	var matches []UserData
	response := s.UsersList()

	if response.Ok {
		for _, user := range response.Members {
			if strings.Contains(user.Name, query) ||
				strings.Contains(user.RealName, query) ||
				strings.Contains(user.Profile.Email, query) {
				matches = append(matches, user)
			}
		}
	}

	return matches
}

func (s *SlackAPI) UsersSetActive() Response {
	var response Response
	s.GetRequest(&response, "users.setActive", "token")
	return response
}

func (s *SlackAPI) UsersSetPresence(value string) Response {
	var response Response
	s.GetRequest(&response, "users.setPresence", "token", "presence="+value)
	return response
}

func (s *SlackAPI) UsersProfileGet(user string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.GetRequest(&response,
		"users.profile.get",
		"token",
		"user="+s.UsersId(user),
		"include_labels=1")
	return response
}

func (s *SlackAPI) UsersProfileSet(user string, name string, value string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.GetRequest(&response,
		"users.profile.set",
		"token",
		"user="+s.UsersId(user),
		"name="+name,
		"value="+value)
	return response
}
