package main

func (s *SlackAPI) ChatDelete(channel string, timestamp string) ChannelEvent {
	var response ChannelEvent
	s.GetRequest(&response,
		"chat.delete",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

func (s *SlackAPI) ChatDeleteVerbose(channel string, timestamp string) {
	response := s.ChatDelete(channel, timestamp)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChatPostMessage(channel string, message string) Message {
	var response Message

	if s.RobotIsActive == true {
		var imageType string

		if s.RobotImageType == "emoji" {
			imageType = "icon_emoji"
		} else {
			imageType = "icon_url"
		}

		s.GetRequest(&response,
			"chat.postMessage",
			"token",
			"channel="+channel,
			"text="+message,
			"as_user=false",
			"link_names=1",
			"username="+s.RobotName,
			imageType+"="+s.RobotImage)
	} else {
		s.GetRequest(&response,
			"chat.postMessage",
			"token",
			"channel="+channel,
			"text="+message,
			"as_user=true",
			"link_names=1")
	}

	return response
}

func (s *SlackAPI) ChatPostMessageVerbose(channel string, message string) {
	response := s.ChatPostMessage(channel, message)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ChatUpdate(channel string, timestamp string, message string) Message {
	var response Message
	s.GetRequest(&response,
		"chat.update",
		"token",
		"channel="+channel,
		"text="+message,
		"ts="+timestamp,
		"link_names=1")
	return response
}

func (s *SlackAPI) ChatUpdateVerbose(channel string, timestamp string, message string) {
	response := s.ChatUpdate(channel, timestamp, message)
	s.PrintAndExit(response)
}
