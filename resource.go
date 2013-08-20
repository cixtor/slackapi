package main

import (
	"fmt"
)

func (s *SlackAPI) ResourceHistory(action string, channel string, latest string) History {
	var response History
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"latest="+latest,
		"inclusive=1",
		"count=1000",
		"unreads=1")
	return response
}

func (s *SlackAPI) ResourceMark(action string, channel string, timestamp string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

func (s *SlackAPI) ResourceMyHistory(action string, channel string, latest string) MyHistory {
	var owner Owner = s.AuthTest()
	var rhistory MyHistory

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

func (s *SlackAPI) ResourcePurgeHistory(action string, channel string, latest string, verbose bool) DeletedHistory {
	var delhist DeletedHistory
	var delmsg DeletedMessage

	response := s.ResourceMyHistory(action, channel, latest)

	if response.Filtered > 0 {
		if verbose {
			fmt.Printf("@ Deleting %d messages\n", response.Filtered)
		}

		for _, message := range response.Messages {
			result := s.ChatDelete(channel, message.Ts)
			delmsg.Text = message.Text
			delmsg.Ts = message.Ts

			if verbose {
				fmt.Printf("\x20 %s from %s ", message.Ts, channel)
			}

			if result.Ok == true {
				delhist.Deleted++
				delmsg.Deleted = true

				if verbose {
					fmt.Println("\u2714")
				}
			} else {
				delhist.NotDeleted++
				delmsg.Deleted = false

				if verbose {
					fmt.Printf("\u2718 %s\n", result.Error)
				}
			}

			delhist.Messages = append(delhist.Messages, delmsg)
		}
	}

	return delhist
}

func (s *SlackAPI) ResourceSetPurpose(action string, channel string, purpose string) ChannelPurposeNow {
	var response ChannelPurposeNow
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"purpose="+purpose)
	return response
}

func (s *SlackAPI) ResourceSetRetention(action string, channel string, duration string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"retention_type=1",
		"retention_duration="+duration)
	return response
}

func (s *SlackAPI) ResourceSetTopic(action string, channel string, topic string) ChannelTopicNow {
	var response ChannelTopicNow
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"topic="+topic)
	return response
}
