package main

func (s *SlackAPI) FilesDelete(file string) {
	var response interface{}
	s.GetRequest(&response,
		"files.delete",
		"token",
		"file="+file)
	s.PrintAndExit(response)
}

func (s *SlackAPI) FilesUpload(channel string, file string) {
	var response interface{}
	s.PostRequest(&response,
		"files.upload",
		"token",
		"file=@"+file,
		"channels="+channel)
	s.PrintAndExit(response)
}
