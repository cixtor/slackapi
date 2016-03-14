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
