package slackapi

// ResponseStarsList defines the JSON-encoded output for StarsList.
type ResponseStarsList struct {
	Response
	Items  []StarsListItem `json:"items"`
	Paging Pagination      `json:"paging"`
}

// StarsListItem defines the expected data from the JSON-encoded API response.
type StarsListItem struct {
	Type    string  `json:"type"`
	Channel string  `json:"channel"`
	Message Message `json:"message"`
	File    File    `json:"file"`
	Comment Comment `json:"comment"`
}

// StarsList lists stars for a user.
func (s *SlackAPI) StarsList(count int, page int) ResponseStarsList {
	var response ResponseStarsList
	s.GetRequest(&response, "stars.list", struct {
		Count int `json:"count"`
		Page  int `json:"page"`
	}{count, page})
	return response
}
