package main

func (s *SlackAPI) FilesDelete(file string) {
	var response interface{}
	s.GetRequest(&response,
		"files.delete",
		"token",
		"file="+file)
	s.PrintAndExit(response)
}
