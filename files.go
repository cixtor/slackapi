package main

// ResponseFilesInfo defines the JSON-encoded output for FilesInfo.
type ResponseFilesInfo struct {
	Response
	File     File          `json:"file"`
	Comments []FileComment `json:"comments"`
	Paging   Pagination    `json:"paging"`
}

// ResponseFilesList defines the JSON-encoded output for FilesList.
type ResponseFilesList struct {
	Response
	Files  []File     `json:"files"`
	Paging Pagination `json:"paging"`
}

// ResponseFilesUpload defines the JSON-encoded output for FilesUpload.
type ResponseFilesUpload struct {
	Response
	File File `json:"file"`
}

// ResponseFilesComments defines the JSON-encoded output for FilesComments.
type ResponseFilesComments struct {
	Response
	Comment FileComment `json:"comment"`
}

// File defines the expected data from the JSON-encoded API response.
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
	ID                 string      `json:"id"`
	ImageExifRotation  int         `json:"image_exif_rotation"`
	InstantMessages    []string    `json:"ims"`
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
	PublicURLShared    bool        `json:"public_url_shared"`
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
	URL                string      `json:"url"`
	URLDownload        string      `json:"url_download"`
	URLPrivate         string      `json:"url_private"`
	URLPrivateDownload string      `json:"url_private_download"`
	User               string      `json:"user"`
	Username           string      `json:"username"`
}

// FileComment defines the expected data from the JSON-encoded API response.
type FileComment struct {
	Comment   string `json:"comment"`
	Created   int    `json:"created"`
	ID        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	User      string `json:"user"`
}

// FilesCommentsAdd add a comment to an existing file.
func (s *SlackAPI) FilesCommentsAdd(file string, text string) ResponseFilesComments {
	var response ResponseFilesComments
	s.GetRequest(&response,
		"files.comments.add",
		"token",
		"file="+file,
		"comment="+text)
	return response
}

// FilesCommentsDelete deletes an existing comment on a file.
func (s *SlackAPI) FilesCommentsDelete(file string, textid string) Response {
	var response Response
	s.GetRequest(&response,
		"files.comments.delete",
		"token",
		"file="+file,
		"id="+textid)
	return response
}

// FilesCommentsEdit edit an existing file comment.
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

// FilesDelete deletes a file.
func (s *SlackAPI) FilesDelete(file string) Response {
	var response Response
	s.GetRequest(&response,
		"files.delete",
		"token",
		"file="+file)
	return response
}

// FilesInfo gets information about a team file.
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

// FilesList lists and filters team files.
// FilesListAfterTime lists and filters team files after this timestamp (inclusive).
// FilesListBeforeTime lists and filters team files before this timestamp (inclusive).
// FilesListByChannel lists and filters team files in a specific channel.
// FilesListByType lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs.
// FilesListByUser lists and filters team files created by a single user.
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

// FilesUpload uploads or creates a file.
func (s *SlackAPI) FilesUpload(channel string, file string) ResponseFilesUpload {
	var response ResponseFilesUpload
	s.PostRequest(&response,
		"files.upload",
		"token",
		"file=@"+file,
		"channels="+channel)
	return response
}
