package main

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
