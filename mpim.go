package main

type ResponseMPIMList struct {
	Response
	Groups []MPIMGroup `json:"groups"`
}

type MPIMGroup struct {
	Created    int              `json:"created"`
	Creator    string           `json:"creator"`
	Id         string           `json:"id"`
	IsArchived bool             `json:"is_archived"`
	IsGroup    bool             `json:"is_group"`
	IsMpim     bool             `json:"is_mpim"`
	Members    []string         `json:"members"`
	Name       string           `json:"name"`
	Purpose    MPIMGroupPurpose `json:"purpose"`
	Topic      MPIMGroupTopic   `json:"topic"`
}

type MPIMGroupPurpose struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

type MPIMGroupTopic struct {
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
	Value   string `json:"value"`
}

func (s *SlackAPI) MultiPartyInstantMessagingList() ResponseMPIMList {
	var response ResponseMPIMList
	s.GetRequest(&response, "mpim.list", "token")
	return response
}
