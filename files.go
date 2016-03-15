package main

type ResponseFilesUpload struct {
	Ok   bool `json:"ok"`
	File File `json:"file"`
}

type ResponseFilesComments struct {
	Ok      bool        `json:"ok"`
	Comment FileComment `json:"comment"`
}

type File struct {
	Channels           []string `json:"channels"`
	CommentsCount      int      `json:"comments_count"`
	Created            int      `json:"created"`
	DisplayAsBot       bool     `json:"display_as_bot"`
	Editable           bool     `json:"editable"`
	ExternalType       string   `json:"external_type"`
	Filetype           string   `json:"filetype"`
	Groups             []string `json:"groups"`
	Id                 string   `json:"id"`
	ImageExifRotation  int      `json:"image_exif_rotation"`
	Ims                []string `json:"ims"`
	IsExternal         bool     `json:"is_external"`
	IsPublic           bool     `json:"is_public"`
	Mimetype           string   `json:"mimetype"`
	Mode               string   `json:"mode"`
	Name               string   `json:"name"`
	OriginalH          int      `json:"original_h"`
	OriginalW          int      `json:"original_w"`
	Permalink          string   `json:"permalink"`
	PrettyType         string   `json:"pretty_type"`
	PublicUrlShared    bool     `json:"public_url_shared"`
	Size               int      `json:"size"`
	Thumb160           string   `json:"thumb_160"`
	Thumb360           string   `json:"thumb_360"`
	Thumb360H          int      `json:"thumb_360_h"`
	Thumb360W          int      `json:"thumb_360_w"`
	Thumb480           string   `json:"thumb_480"`
	Thumb480H          int      `json:"thumb_480_h"`
	Thumb480W          int      `json:"thumb_480_w"`
	Thumb64            string   `json:"thumb_64"`
	Thumb80            string   `json:"thumb_80"`
	Timestamp          int      `json:"timestamp"`
	Title              string   `json:"title"`
	UrlPrivate         string   `json:"url_private"`
	UrlPrivateDownload string   `json:"url_private_download"`
	User               string   `json:"user"`
	Username           string   `json:"username"`
}

type FileComment struct {
	Comment   string `json:"comment"`
	Created   int    `json:"created"`
	Id        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	User      string `json:"user"`
}

func (s *SlackAPI) FilesCommentsAdd(file string, text string) ResponseFilesComments {
	var response ResponseFilesComments
	s.GetRequest(&response,
		"files.comments.add",
		"token",
		"file="+file,
		"comment="+text)
	return response
}

func (s *SlackAPI) FilesCommentsDelete(file string, textid string) Response {
	var response Response
	s.GetRequest(&response,
		"files.comments.delete",
		"token",
		"file="+file,
		"id="+textid)
	return response
}

func (s *SlackAPI) FilesCommentsEdit(file string, textid string, text string) ResponseFilesComments {
	var response ResponseFilesComments
	s.GetRequest(&response,
		"files.comments.edit",
		"token",
		"file="+file,
		"id="+textid,
		"comment="+text)
	return response
}

func (s *SlackAPI) FilesDelete(file string) Response {
	var response Response
	s.GetRequest(&response,
		"files.delete",
		"token",
		"file="+file)
	return response
}

func (s *SlackAPI) FilesUpload(channel string, file string) ResponseFilesUpload {
	var response ResponseFilesUpload
	s.PostRequest(&response,
		"files.upload",
		"token",
		"file=@"+file,
		"channels="+channel)
	return response
}
