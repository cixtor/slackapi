package main

func (s *SlackAPI) FilesCommentsAdd(file string, text string) {
	var response interface{}
	s.GetRequest(&response,
		"files.comments.add",
		"token",
		"file="+file,
		"comment="+text)
	s.PrintAndExit(response)
}

func (s *SlackAPI) FilesCommentsDelete(file string, textid string) {
	var response interface{}
	s.GetRequest(&response,
		"files.comments.delete",
		"token",
		"file="+file,
		"id="+textid)
	s.PrintAndExit(response)
}

func (s *SlackAPI) FilesCommentsEdit(file string, textid string, text string) {
	var response interface{}
	s.GetRequest(&response,
		"files.comments.edit",
		"token",
		"file="+file,
		"id="+textid,
		"comment="+text)
	s.PrintAndExit(response)
}

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
