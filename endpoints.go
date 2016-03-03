package main

import (
	"fmt"
)

func (s *SlackAPI) ApiTest() {
	var response interface{}
	s.GetRequest(&response, "api.test")
	s.PrintAndExit(response)
}

func (s *SlackAPI) AuthTest() Owner {
	if s.Owner.Ok == true {
		return s.Owner
	}

	var response Owner
	s.GetRequest(&response, "auth.test", "token")
	s.Owner = response

	return response
}

func (s *SlackAPI) AuthTestVerbose() {
	response := s.AuthTest()
	s.PrintAndExit(response)
}

func (s *SlackAPI) EmojiList() {
	var response interface{}
	s.GetRequest(&response, "emoji.list", "token")
	s.PrintAndExit(response)
}

func (s *SlackAPI) ResourceHistory(action string, channel string, latest string) History {
	var response History

	if latest == "" {
		s.GetRequest(&response,
			action,
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"unreads=1")
	} else {
		s.GetRequest(&response,
			action,
			"token",
			"channel="+channel,
			"inclusive=1",
			"count=1000",
			"latest="+latest,
			"unreads=1")
	}

	return response
}

func (s *SlackAPI) ResourceHistoryVerbose(action string, channel string, latest string) {
	response := s.ResourceHistory(action, channel, latest)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ResourceMark(action string, channel string, timestamp string) {
	var response interface{}
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"ts="+timestamp)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ResourceMyHistory(action string, channel string, latest string) UserHistory {
	var owner Owner = s.AuthTest()
	var rhistory UserHistory

	response := s.ResourceHistory(action, channel, latest)

	for _, message := range response.Messages {
		rhistory.Total += 1

		if message.User == owner.UserId {
			rhistory.Messages = append(rhistory.Messages, message)
			rhistory.Filtered += 1
		}
	}

	if rhistory.Total > 0 {
		var offset int = len(response.Messages) - 1

		rhistory.Username = owner.User
		rhistory.Latest = response.Messages[0].Ts
		rhistory.Oldest = response.Messages[offset].Ts
	}

	return rhistory
}

func (s *SlackAPI) ResourceMyHistoryVerbose(action string, channel string, latest string) {
	response := s.ResourceMyHistory(action, channel, latest)
	s.PrintAndExit(response)
}

func (s *SlackAPI) ResourcePurgeHistory(action string, channel string, latest string) {
	response := s.ResourceMyHistory(action, channel, latest)

	if response.Filtered > 0 {
		fmt.Printf("@ Deleting %d messages\n", response.Filtered)

		for _, message := range response.Messages {
			fmt.Printf("\x20 %s from %s ", message.Ts, channel)
			result := s.ChatDelete(channel, message.Ts)

			if result.Ok == true {
				fmt.Println("\u2714")
			} else {
				fmt.Printf("\u2718 %s\n", result.Error)
			}
		}
	}
}

func (s *SlackAPI) TeamInfo() {
	var response interface{}
	s.GetRequest(&response, "team.info", "token")
	s.PrintAndExit(response)
}
