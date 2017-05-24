package slackapi

import (
	"fmt"
	"time"
)

// ResourceArchive archives a channel.
func (s *SlackAPI) ResourceArchive(action string, channel string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel)
	return response
}

// ResourceHistory fetches history of messages and events from a channel.
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

// ResourceInvite invites a user to a channel.
func (s *SlackAPI) ResourceInvite(action string, channel string, user string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"user="+user)
	return response
}

// ResourceKick removes a user from a channel.
func (s *SlackAPI) ResourceKick(action string, channel string, user string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"user="+user)
	return response
}

// ResourceLeave leaves a channel.
func (s *SlackAPI) ResourceLeave(action string, channel string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel)
	return response
}

// ResourceMark sets the read cursor in a channel.
func (s *SlackAPI) ResourceMark(action string, channel string, timestamp string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"ts="+timestamp)
	return response
}

// ResourceMyHistory displays messages of the current user from a channel.
func (s *SlackAPI) ResourceMyHistory(action string, channel string, latest string) MyHistory {
	var rhistory MyHistory

	owner := s.AuthTest()
	response := s.ResourceHistory(action, channel, latest)

	for _, message := range response.Messages {
		rhistory.Total++

		if message.User == owner.UserID {
			rhistory.Messages = append(rhistory.Messages, message)
			rhistory.Filtered++
		}
	}

	if rhistory.Total > 0 {
		offset := len(response.Messages) - 1

		rhistory.Username = owner.User
		rhistory.Latest = response.Messages[0].Ts
		rhistory.Oldest = response.Messages[offset].Ts
	}

	return rhistory
}

// ResourcePurgeHistory deletes history of messages and events from a channel.
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

				if result.Error == "RATELIMIT" {
					time.Sleep(10 * time.Second)
				}
			}

			delhist.Messages = append(delhist.Messages, delmsg)
		}
	}

	return delhist
}

// ResourceRename renames a channel.
func (s *SlackAPI) ResourceRename(action string, channel string, name string) ChannelRename {
	var response ChannelRename
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"name="+name)
	return response
}

// ResourceSetPurpose sets the purpose for a channel.
func (s *SlackAPI) ResourceSetPurpose(action string, channel string, purpose string) ChannelPurposeNow {
	var response ChannelPurposeNow
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"purpose="+purpose)
	return response
}

// ResourceSetRetention sets the retention time of the messages.
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

// ResourceSetTopic sets the topic for a channel.
func (s *SlackAPI) ResourceSetTopic(action string, channel string, topic string) ChannelTopicNow {
	var response ChannelTopicNow
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel,
		"topic="+topic)
	return response
}

// ResourceUnarchive unarchives a channel.
func (s *SlackAPI) ResourceUnarchive(action string, channel string) Response {
	var response Response
	s.GetRequest(&response,
		action,
		"token",
		"channel="+channel)
	return response
}
