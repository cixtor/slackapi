package main

type ResponseChannelsInfo struct {
	Response
	Channel Channel `json:"channel"`
}

type ResponseChannelsJoin struct {
	Response
	AlreadyInChannel bool    `json:"already_in_channel"`
	Channel          Channel `json:"channel"`
}

type ResponseChannelsList struct {
	Response
	Channels []Channel `json:"channels"`
}

func (s *SlackAPI) ChannelsArchive(channel string) Response {
	return s.ResourceArchive("channels.archive", s.ChannelsId(channel))
}

func (s *SlackAPI) ChannelsCreate(channel string) ResponseChannelsInfo {
	var response ResponseChannelsInfo
	s.GetRequest(&response, "channels.create", "token", "name="+channel)
	return response
}

func (s *SlackAPI) ChannelsHistory(channel string, latest string) History {
	return s.ResourceHistory("channels.history", channel, latest)
}

func (s *SlackAPI) ChannelsId(query string) string {
	response := s.ChannelsList()

	if response.Ok {
		for _, room := range response.Channels {
			if room.Name == query {
				return room.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) ChannelsInfo(channel string) ResponseChannelsInfo {
	var response ResponseChannelsInfo
	channel = s.ChannelsId(channel)
	s.GetRequest(&response, "channels.info", "token", "channel="+channel)
	return response
}

func (s *SlackAPI) ChannelsInvite(channel string, user string) Response {
	return s.ResourceInvite("channels.invite", s.ChannelsId(channel), s.UsersId(user))
}

func (s *SlackAPI) ChannelsJoin(channel string) ResponseChannelsJoin {
	var response ResponseChannelsJoin
	s.GetRequest(&response, "channels.join", "token", "name="+channel)
	return response
}

func (s *SlackAPI) ChannelsKick(channel string, user string) Response {
	return s.ResourceKick("channels.kick", s.ChannelsId(channel), s.UsersId(user))
}

func (s *SlackAPI) ChannelsLeave(channel string) Response {
	return s.ResourceLeave("channels.leave", s.ChannelsId(channel))
}

func (s *SlackAPI) ChannelsList() ResponseChannelsList {
	if s.TeamChannels.Ok == true {
		return s.TeamChannels
	}

	var response ResponseChannelsList
	s.GetRequest(&response, "channels.list", "token", "exclude_archived=0")
	s.TeamChannels = response

	return response
}

func (s *SlackAPI) ChannelsMark(channel string, timestamp string) Response {
	return s.ResourceMark("channels.mark", channel, timestamp)
}

func (s *SlackAPI) ChannelsMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("channels.history", channel, latest)
}

func (s *SlackAPI) ChannelsPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("channels.history", channel, latest, verbose)
}

func (s *SlackAPI) ChannelsRename(channel string, name string) ChannelRename {
	return s.ResourceRename("channels.rename", s.ChannelsId(channel), name)
}

func (s *SlackAPI) ChannelsSetPurpose(channel string, purpose string) ChannelPurposeNow {
	return s.ResourceSetPurpose("channels.setPurpose", channel, purpose)
}

func (s *SlackAPI) ChannelsSetRetention(channel string, duration string) Response {
	return s.ResourceSetRetention("channels.setRetention", channel, duration)
}

func (s *SlackAPI) ChannelsSetTopic(channel string, topic string) ChannelTopicNow {
	return s.ResourceSetTopic("channels.setTopic", channel, topic)
}

func (s *SlackAPI) ChannelsUnarchive(channel string) Response {
	return s.ResourceUnarchive("channels.unarchive", s.ChannelsId(channel))
}
