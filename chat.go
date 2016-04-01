package main

func (s *SlackAPI) ChatDelete(channel string, timestamp string) Post {
	var response Post
	s.GetRequest(&response,
		"chat.delete",
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

func (s *SlackAPI) ChatPostMessage(channel string, message string) Post {
	var response Post

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
			"parse=none",
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
			"parse=none",
			"channel="+channel,
			"text="+message,
			"as_user=true",
			"link_names=1")
	}

	return response
}

func (s *SlackAPI) ChatUpdate(channel string, timestamp string, message string) Post {
	var response Post
	s.GetRequest(&response,
		"chat.update",
		"token",
		"parse=none",
		"channel="+channel,
		"text="+message,
		"ts="+timestamp,
		"link_names=1")
	return response
}
