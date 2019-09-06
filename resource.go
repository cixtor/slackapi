package slackapi

import (
	"fmt"
	"time"
)

// HistoryArgs defines the data to send to the API service.
type HistoryArgs struct {
	Channel   string `json:"channel"`
	Latest    string `json:"latest"`
	Oldest    string `json:"oldest"`
	Count     int    `json:"count"`
	Inclusive bool   `json:"inclusive"`
	Unreads   bool   `json:"unreads"`
}

// ResourceArchive archives a channel.
func (s *SlackAPI) ResourceArchive(action string, channel string) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}

// ResourceHistory fetches history of messages and events from a channel.
func (s *SlackAPI) ResourceHistory(action string, data HistoryArgs) History {
	var response History
	if data.Count == 0 {
		data.Count = 100
	}
	if data.Latest == "" {
		data.Latest = fmt.Sprintf("%d", time.Now().Unix())
	}
	s.getRequest(&response, action, data)
	return response
}

// ResourceInvite invites a user to a channel.
func (s *SlackAPI) ResourceInvite(action string, channel string, user string) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
		User    string `json:"user"`
	}{channel, user})
	return response
}

// ResourceKick removes a user from a channel.
func (s *SlackAPI) ResourceKick(action string, channel string, user string) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
		User    string `json:"user"`
	}{channel, user})
	return response
}

// ResourceLeave leaves a channel.
func (s *SlackAPI) ResourceLeave(action string, channel string) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}

// ResourceMark sets the read cursor in a channel.
func (s *SlackAPI) ResourceMark(action string, channel string, ts string) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
		Ts      string `json:"ts"`
	}{channel, ts})
	return response
}

// ResourceMyHistory displays messages of the current user from a channel.
func (s *SlackAPI) ResourceMyHistory(action string, channel string, latest string) MyHistory {
	owner, err := s.AuthTest()

	if err != nil {
		return MyHistory{}
	}

	response := s.ResourceHistory(action, HistoryArgs{
		Channel: channel,
		Latest:  latest,
	})

	var rhistory MyHistory

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
		rhistory.Latest = response.Messages[0].Timestamp
		rhistory.Oldest = response.Messages[offset].Timestamp
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
			result := s.ChatDelete(MessageArgs{
				Channel: channel,
				Ts:      message.Timestamp,
			})

			delmsg.Text = message.Text
			delmsg.Timestamp = message.Timestamp

			if verbose {
				fmt.Printf("\x20 %s from %s ", message.Timestamp, channel)
			}

			if result.Ok {
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

				if result.Error == "RATELIMIT" || result.Error == "ratelimited" {
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
	s.postRequest(&response, action, struct {
		Name     string `json:"name"`
		Channel  string `json:"channel"`
		Validate bool   `json:"validate"`
	}{name, channel, true})
	return response
}

// ResourceSetPurpose sets the purpose for a channel.
func (s *SlackAPI) ResourceSetPurpose(action string, channel string, purpose string) ChannelPurposeNow {
	var response ChannelPurposeNow
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
		Purpose string `json:"purpose"`
	}{channel, purpose})
	return response
}

// ResourceSetRetention sets the retention time of the messages.
func (s *SlackAPI) ResourceSetRetention(action string, channel string, duration int) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel           string `json:"channel"`
		RetentionType     bool   `json:"retention_type"`
		RetentionDuration int    `json:"retention_duration"`
	}{channel, true, duration})
	return response
}

// ResourceSetTopic sets the topic for a channel.
func (s *SlackAPI) ResourceSetTopic(action string, channel string, topic string) ChannelTopicNow {
	var response ChannelTopicNow
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
		Topic   string `json:"topic"`
	}{channel, topic})
	return response
}

// ResourceUnarchive unarchives a channel.
func (s *SlackAPI) ResourceUnarchive(action string, channel string) Response {
	var response Response
	s.postRequest(&response, action, struct {
		Channel string `json:"channel"`
	}{channel})
	return response
}
