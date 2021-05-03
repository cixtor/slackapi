package slackapi

// ResponseStarsList defines the JSON-encoded output for StarsList.
type ResponseStarsList struct {
	Response
	Items  []StarsListItem `json:"items"`
	Paging Paging          `json:"paging"`
}

// StarsListItem defines the expected data from the JSON-encoded API response.
type StarsListItem struct {
	Type    string  `json:"type"`
	Channel string  `json:"channel"`
	Message Message `json:"message"`
	File    File    `json:"file"`
	Comment Comment `json:"comment"`
}

// StarsAdd adds a star to an item.
func (s *SlackAPI) StarsAdd(channel string, itemid string) Response {
	var response Response

	if len(itemid) >= 3 && itemid[0:2] == "Fc" {
		/* remove pinned file comment */
		s.postRequest(&response, "stars.add", struct {
			Channel     string `json:"channel"`
			FileComment string `json:"file_comment"`
		}{
			Channel:     channel,
			FileComment: itemid,
		})
	} else if len(itemid) >= 2 && itemid[0] == 'F' {
		/* remove pinned file */
		s.postRequest(&response, "stars.add", struct {
			Channel string `json:"channel"`
			File    string `json:"file"`
		}{
			Channel: channel,
			File:    itemid,
		})
	} else {
		/* remove pinned message */
		s.postRequest(&response, "stars.add", struct {
			Channel   string `json:"channel"`
			Timestamp string `json:"timestamp"`
		}{
			Channel:   channel,
			Timestamp: itemid,
		})
	}

	return response
}

// StarsList lists stars for a user.
func (s *SlackAPI) StarsList(count int, page int) ResponseStarsList {
	var response ResponseStarsList
	s.getRequest(&response, "stars.list", struct {
		Count int `json:"count"`
		Page  int `json:"page"`
	}{
		Count: count,
		Page:  page,
	})
	return response
}

// StarsRemove removes a star from an item.
func (s *SlackAPI) StarsRemove(channel string, itemid string) Response {
	var response Response

	if len(itemid) >= 3 && itemid[0:2] == "Fc" {
		/* remove pinned file comment */
		s.postRequest(&response, "stars.remove", struct {
			Channel     string `json:"channel"`
			FileComment string `json:"file_comment"`
		}{
			Channel:     channel,
			FileComment: itemid,
		})
	} else if len(itemid) >= 2 && itemid[0] == 'F' {
		/* remove pinned file */
		s.postRequest(&response, "stars.remove", struct {
			Channel string `json:"channel"`
			File    string `json:"file"`
		}{
			Channel: channel,
			File:    itemid,
		})
	} else {
		/* remove pinned message */
		s.postRequest(&response, "stars.remove", struct {
			Channel   string `json:"channel"`
			Timestamp string `json:"timestamp"`
		}{
			Channel:   channel,
			Timestamp: itemid,
		})
	}

	return response
}
