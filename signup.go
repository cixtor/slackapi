package slackapi

// ResponseEmailCheck defines the JSON-encoded output for Revocation.
type ResponseEmailCheck struct {
	Response
	Type    string `json:"type"`
	Email   string `json:"email"`
	AuthURL string `json:"auth_url"`
}

// SignupCheckEmail checks if an email address is valid.
func (s *SlackAPI) SignupCheckEmail(email string) ResponseEmailCheck {
	var response ResponseEmailCheck
	s.postRequest(&response, "signup.checkEmail", struct {
		Email   string `json:"email"`
		GetInfo bool   `json:"get_info"`
	}{email, true})
	return response
}

// SignupConfirmEmail confirm an email address for signup.
func (s *SlackAPI) SignupConfirmEmail(email string) Response {
	var response Response
	s.postRequest(&response, "signup.confirmEmail", struct {
		Email string `json:"email"`
	}{email})
	return response
}
