package main

type ResponseGroupsInfo struct {
	Response
	Group Group `json:"group"`
}

type ResponseGroupsList struct {
	Response
	Groups []Group `json:"groups"`
}

type ResponsePurpose struct {
	Response
	Purpose string `json:"purpose"`
}

type ResponseTopic struct {
	Response
	Topic string `json:"topic"`
}

type Group struct {
	Created            int          `json:"created"`
	Creator            string       `json:"creator"`
	Id                 string       `json:"id"`
	IsArchived         bool         `json:"is_archived"`
	IsGroup            bool         `json:"is_group"`
	IsMpim             bool         `json:"is_mpim"`
	IsOpen             bool         `json:"is_open"`
	LastRead           string       `json:"last_read"`
	Latest             GroupLatest  `json:"latest"`
	Members            []string     `json:"members"`
	Name               string       `json:"name"`
	Purpose            GroupPurpose `json:"purpose"`
	Topic              GroupTopic   `json:"topic"`
	UnreadCount        int          `json:"unread_count"`
	UnreadCountDisplay int          `json:"unread_count_display"`
}

type GroupLatest struct {
	Text string `json:"text"`
	Ts   string `json:"ts"`
	Type string `json:"type"`
	User string `json:"user"`
}

type GroupPurpose struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

type GroupTopic struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

func (s *SlackAPI) GroupsClose(channel string) Response {
	var response Response
	s.GetRequest(&response, "groups.close", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsHistory(channel string, latest string) History {
	return s.ResourceHistory("groups.history", channel, latest)
}

func (s *SlackAPI) GroupsId(query string) string {
	response := s.GroupsList()

	if response.Ok {
		for _, room := range response.Groups {
			if room.Name == query {
				return room.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) GroupsInfo(channel string) ResponseGroupsInfo {
	var response ResponseGroupsInfo
	channel = s.GroupsId(channel)
	s.GetRequest(&response, "groups.info", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsList() ResponseGroupsList {
	if s.TeamGroups.Ok == true {
		return s.TeamGroups
	}

	var response ResponseGroupsList
	s.GetRequest(&response, "groups.list", "token", "exclude_archived=0")
	s.TeamGroups = response

	return response
}

func (s *SlackAPI) GroupsMark(channel string, timestamp string) Response {
	return s.ResourceMark("groups.mark", channel, timestamp)
}

func (s *SlackAPI) GroupsMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("groups.history", channel, latest)
}

func (s *SlackAPI) GroupsOpen(channel string) Session {
	var response Session
	channel = s.GroupsId(channel)
	s.GetRequest(&response, "groups.open", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) GroupsPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("groups.history", channel, latest, verbose)
}

func (s *SlackAPI) GroupsSetPurpose(channel string, purpose string) ResponsePurpose {
	var response ResponsePurpose
	s.GetRequest(&response,
		"groups.setPurpose",
		"token",
		"channel="+channel,
		"purpose="+purpose)
	return response
}

func (s *SlackAPI) GroupsSetTopic(channel string, topic string) ResponseTopic {
	var response ResponseTopic
	s.GetRequest(&response,
		"groups.setTopic",
		"token",
		"channel="+channel,
		"topic="+topic)
	return response
}
