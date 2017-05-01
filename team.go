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

type ResponseTeamProfile struct {
	Response
	Profile TeamProfile `json:"profile"`
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

type TeamProfile struct {
	Fields []TeamProfileField `json:"fields"`
}

type TeamProfileField struct {
	ID             string      `json:"id"`
	Ordering       int         `json:"ordering"`
	FieldName      string      `json:"field_name"`
	Label          string      `json:"label"`
	Hint           string      `json:"hint"`
	Type           string      `json:"type"`
	PossibleValues interface{} `json:"possible_values"`
	Options        interface{} `json:"options"`
	IsHidden       bool        `json:"is_hidden"`
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

type BillableInfo struct {
	BillingActive bool `json:"billing_active"`
}

type ResponseBillableInfo struct {
	Response
	BillableInfo map[string]BillableInfo `json:"billable_info"`
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

func (s *SlackAPI) TeamBillableInfo(user string) ResponseBillableInfo {
	var response ResponseBillableInfo

	if user == "" {
		s.GetRequest(&response, "team.billableInfo", "token")
	} else {
		s.GetRequest(&response, "team.billableInfo", "token", "user="+s.UsersId(user))
	}

	return response
}

func (s *SlackAPI) TeamInfo() ResponseTeamInfo {
	var response ResponseTeamInfo
	s.GetRequest(&response, "team.info", "token")
	return response
}

func (s *SlackAPI) TeamProfileGet() ResponseTeamProfile {
	var response ResponseTeamProfile
	s.GetRequest(&response, "team.profile.get", "token")
	return response
}
