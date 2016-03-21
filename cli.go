package main

func (s *SlackAPI) PrintEmojiList() {
	response := s.EmojiList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintFilesCommentsAdd(file string, text string) {
	response := s.FilesCommentsAdd(file, text)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintFilesCommentsDelete(file string, textid string) {
	response := s.FilesCommentsDelete(file, textid)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintFilesCommentsEdit(file string, textid string, text string) {
	response := s.FilesCommentsEdit(file, textid, text)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintFilesDelete(file string) {
	response := s.FilesDelete(file)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintFilesUpload(channel string, file string) {
	response := s.FilesUpload(channel, file)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsClose(channel string) {
	response := s.GroupsClose(channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsHistory(channel string, latest string) {
	response := s.GroupsHistory(channel, latest)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsId(query string) {
	response := s.GroupsId(query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsInfo(channel string) {
	response := s.GroupsInfo(channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsList() {
	response := s.GroupsList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsMark(channel string, timestamp string) {
	response := s.GroupsMark(channel, timestamp)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsMyHistory(channel string, latest string) {
	response := s.GroupsMyHistory(channel, latest)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsOpen(channel string) {
	response := s.GroupsOpen(channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsPurgeHistory(channel string, latest string) {
	s.GroupsPurgeHistory(channel, latest, true)
}

func (s *SlackAPI) PrintGroupsSetPurpose(channel string, purpose string) {
	response := s.GroupsSetPurpose(channel, purpose)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintGroupsSetTopic(channel string, topic string) {
	response := s.GroupsSetTopic(channel, topic)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingClose(channel string) {
	response := s.InstantMessagingClose(channel)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingHistory(channel string, latest string) {
	response := s.InstantMessagingHistory(channel, latest)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingList() {
	response := s.InstantMessagingList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingMark(channel string, timestamp string) {
	response := s.InstantMessagingMark(channel, timestamp)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingMyHistory(channel string, latest string) {
	response := s.InstantMessagingMyHistory(channel, latest)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingOpen(userid string) {
	response := s.InstantMessagingOpen(userid)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintInstantMessagingPurgeHistory(channel string, latest string) {
	s.InstantMessagingPurgeHistory(channel, latest, true)
}

func (s *SlackAPI) PrintMultiPartyInstantMessagingList() {
	response := s.MultiPartyInstantMessagingList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintReactionsAdd(name string, channel string, timestamp string) {
	response := s.ReactionsAdd(name, channel, timestamp)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintReactionsGet(resource string, unique string) {
	response := s.ReactionsGet(resource, unique)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintReactionsList(userid string) {
	response := s.ReactionsList(userid)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintReactionsRemove(name string, channel string, timestamp string) {
	response := s.ReactionsRemove(name, channel, timestamp)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintTeamInfo() {
	response := s.TeamInfo()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersGetPresence(query string) {
	response := s.UsersGetPresence(query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersId(query string) {
	response := s.UsersId(query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersInfo(query string) {
	response := s.UsersInfo(query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersList() {
	response := s.UsersList()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersSearch(query string) {
	response := s.UsersSearch(query)
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersSetActive() {
	response := s.UsersSetActive()
	s.PrintAndExit(response)
}

func (s *SlackAPI) PrintUsersSetPresence(value string) {
	response := s.UsersSetPresence(value)
	s.PrintAndExit(response)
}
