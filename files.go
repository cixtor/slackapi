package main

type ResponseFilesInfo struct {
	Response
	File     File          `json:"file"`
	Comments []FileComment `json:"comments"`
	Paging   Pagination    `json:"paging"`
}

type ResponseFilesList struct {
	Response
	Files  []File     `json:"files"`
	Paging Pagination `json:"paging"`
}

type ResponseFilesUpload struct {
	Response
	File File `json:"file"`
}

type ResponseFilesComments struct {
	Response
	Comment FileComment `json:"comment"`
}

type File struct {
	Channels           []string    `json:"channels"`
	CommentsCount      int         `json:"comments_count"`
	Created            int         `json:"created"`
	DisplayAsBot       bool        `json:"display_as_bot"`
	Editable           bool        `json:"editable"`
	EditLink           string      `json:"edit_link"`
	ExternalType       string      `json:"external_type"`
	Filetype           string      `json:"filetype"`
	Groups             []string    `json:"groups"`
	Id                 string      `json:"id"`
	ImageExifRotation  int         `json:"image_exif_rotation"`
	Ims                []string    `json:"ims"`
	InitialComment     FileComment `json:"initial_comment"`
	IsExternal         bool        `json:"is_external"`
	IsPublic           bool        `json:"is_public"`
	IsStarred          bool        `json:"is_starred"`
	Lines              int         `json:"lines"`
	LinesMore          int         `json:"lines_more"`
	Mimetype           string      `json:"mimetype"`
	Mode               string      `json:"mode"`
	Name               string      `json:"name"`
	NumStars           int         `json:"num_stars"`
	OriginalH          int         `json:"original_h"`
	OriginalW          int         `json:"original_w"`
	Permalink          string      `json:"permalink"`
	PrettyType         string      `json:"pretty_type"`
	Preview            string      `json:"preview"`
	PreviewHighlight   string      `json:"preview_highlight"`
	PublicUrlShared    bool        `json:"public_url_shared"`
	Size               int         `json:"size"`
	Thumb160           string      `json:"thumb_160"`
	Thumb360           string      `json:"thumb_360"`
	Thumb360Gif        string      `json:"thumb_360_gif"`
	Thumb360H          int         `json:"thumb_360_h"`
	Thumb360W          int         `json:"thumb_360_w"`
	Thumb480           string      `json:"thumb_480"`
	Thumb480H          int         `json:"thumb_480_h"`
	Thumb480W          int         `json:"thumb_480_w"`
	Thumb64            string      `json:"thumb_64"`
	Thumb80            string      `json:"thumb_80"`
	Timestamp          int         `json:"timestamp"`
	Title              string      `json:"title"`
	Url                string      `json:"url"`
	UrlDownload        string      `json:"url_download"`
	UrlPrivate         string      `json:"url_private"`
	UrlPrivateDownload string      `json:"url_private_download"`
	User               string      `json:"user"`
	Username           string      `json:"username"`
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

func (s *SlackAPI) FilesInfo(file string, count string, page string) ResponseFilesInfo {
	var response ResponseFilesInfo
	s.GetRequest(&response,
		"files.info",
		"token",
		"file="+file,
		"count="+count,
		"page="+page)
	return response
}

func (s *SlackAPI) FilesList(action string, filter string, count string, page string) ResponseFilesList {
	if action != "" && action != "none" {
		s.AddRequestParam(action, filter)
	}

	if count == "" {
		count = "100"
	}

	var response ResponseFilesList
	s.GetRequest(&response,
		"files.list",
		"token",
		"count="+count,
		"page="+page)
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
