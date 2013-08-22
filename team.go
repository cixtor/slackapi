package main

import (
	"encoding/json"
)

type ResponseTeamAccessLogs struct {
	Response
	AccessLogs []AccessLog `json:"logins"`
	Needed     string      `json:"needed"`
	Provided   string      `json:"provided"`
}

type ResponseTeamInfo struct {
	Response
	Team Team `json:"team"`
}

type AccessLog struct {
	UserId    string      `json:"user_id"`
	Username  string      `json:"username"`
	DateFirst json.Number `json:"date_first"`
	DateLast  json.Number `json:"date_last"`
	Count     int         `json:"count"`
	Ip        string      `json:"ip"`
	UserAgent string      `json:"user_agent"`
	Isp       string      `json:"isp"`
	Country   string      `json:"country"`
	Region    string      `json:"region"`
}

type Team struct {
	Domain      string   `json:"domain"`
	EmailDomain string   `json:"email_domain"`
	Icon        TeamIcon `json:"icon"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
}

type TeamIcon struct {
	Image102      string `json:"image_102"`
	Image132      string `json:"image_132"`
	Image34       string `json:"image_34"`
	Image44       string `json:"image_44"`
	Image68       string `json:"image_68"`
	Image88       string `json:"image_88"`
	ImageOriginal string `json:"image_original"`
}

func (s *SlackAPI) TeamAccessLogs(count string, page string) ResponseTeamAccessLogs {
	var response ResponseTeamAccessLogs
	s.GetRequest(&response,
		"team.accessLogs",
		"token",
		"count="+count,
		"page="+page)
	return response
}

func (s *SlackAPI) TeamInfo() ResponseTeamInfo {
	var response ResponseTeamInfo
	s.GetRequest(&response, "team.info", "token")
	return response
}
