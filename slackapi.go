package slackapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

const version = "1.5.44"

// TOKEN defines the name for the token HTTP request parameter.
const TOKEN = "token"

// EMOJI defines the name for the emoji HTTP request parameter.
const EMOJI = "emoji"

// ICONURL defines the name for the icon_url HTTP request parameter.
const ICONURL = "icon_url"

// SlackAPI defines the base object. It holds the API token, the information of
// the user account associated to such API token, the information for the robot
// session (if the user decides has activated it), a list of all the available
// public channels, a list of all the accessible private groups, and a list of
// the users registered into the Slack team.
type SlackAPI struct {
	Owner          Owner
	RequestParams  []string
	RobotImage     string
	RobotImageType string
	RobotIsActive  bool
	RobotName      string
	TeamChannels   ResponseChannelsList
	TeamGroups     ResponseGroupsList
	TeamUsers      ResponseUsersList
	Token          string
}

// New instantiates a new object.
func New() *SlackAPI {
	return &SlackAPI{}
}

// Version returns the package version number.
func (s *SlackAPI) Version() string {
	return version
}

// SetToken sets the API token for the session.
func (s *SlackAPI) SetToken(token string) {
	s.Token = token
}

// AutoConfigure sets the API token from an environment variable.
func (s *SlackAPI) AutoConfigure() {
	s.Token = os.Getenv("SLACK_TOKEN")

	s.RobotName = "bender"
	s.RobotImage = ":robot_face:"
	s.RobotImageType = EMOJI
	s.RobotIsActive = false
}

// System executes an external command and returns its output.
func (s *SlackAPI) System(kommand string) []byte {
	var binary string
	var parts []string
	var arguments []string

	if kommand == "" {
		s.ReportError(errors.New("invalid empty command"))
	}

	parts = strings.Fields(kommand)
	binary = parts[0]
	arguments = parts[1:len(parts)]

	response, err := exec.Command(binary, arguments...).Output()

	if err != nil {
		s.ReportError(err)
	}

	return bytes.Trim(response, "\n")
}

// URL builds and returns the URL to send the HTTP requests.
func (s *SlackAPI) URL(action string, params []string) string {
	data := url.Values{}
	var parts []string
	var encoded string
	var url string = fmt.Sprintf("https://slack.com/api/%s", action)

	for _, keyvalue := range params {
		if keyvalue == TOKEN {
			keyvalue += "=" + s.Token
		}

		parts = strings.SplitN(keyvalue, "=", 2)

		if len(parts) == 2 {
			data.Add(parts[0], parts[1])
		}
	}

	encoded = data.Encode()

	if encoded != "" {
		url = fmt.Sprintf("%s?%s", url, encoded)
	}

	return url
}

// HTTPRequest builds an HTTP request object and attaches the action parameters.
func (s *SlackAPI) HTTPRequest(method string, body io.Reader, action string, params []string) *http.Request {
	if len(s.RequestParams) > 0 {
		params = append(params, s.RequestParams...)
		s.RequestParams = s.RequestParams[:0]
	}

	url := s.URL(action, params)
	req, err := http.NewRequest(method, url, body)

	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36")

	if err != nil {
		s.ReportError(err)
	}

	return req
}

// AddRequestParam adds an additional parameter to the HTTP request.
func (s *SlackAPI) AddRequestParam(param string, value string) {
	var parameter string = fmt.Sprintf("%s=%s", param, value)
	s.RequestParams = append(s.RequestParams, parameter)
}

// ExecuteRequest sends the HTTP request and decodes the JSON response.
func (s *SlackAPI) ExecuteRequest(req *http.Request, data interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		s.ReportError(err)
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)
	if err := json.NewDecoder(tee).Decode(&data); err != nil {
		out, _ := ioutil.ReadAll(&buf)
		if strings.Contains(string(out), "too many requests") {
			fake := "{\"ok\":false, \"error\":\"RATELIMIT\"}"
			read := bytes.NewReader([]byte(fake))
			json.NewDecoder(read).Decode(&data)
			return
		}
		s.ReportError(err)
	}
}

// PrintCurlCommand prints the HTTP request object into the console.
func (s *SlackAPI) PrintCurlCommand(req *http.Request, params []string) {
	if os.Getenv("SLACK_VERBOSE") == "true" {
		fmt.Printf("curl -X %s \"%s\"", req.Method, req.URL)

		for header, values := range req.Header {
			fmt.Printf(" \x5c\n-H \"%s: %s\"", header, values[0])
		}

		fmt.Printf(" \x5c\n-H \"Host: %s\"", req.Host)

		if req.Method == "POST" {
			for _, param := range params {
				if param == TOKEN {
					param = TOKEN + "=" + s.Token
				}
				fmt.Printf(" \x5c\n-d \"%s\"", param)
			}
		}

		fmt.Println()
	}
}

// GetRequest sends a HTTP GET request to the API and returns the response.
func (s *SlackAPI) GetRequest(data interface{}, action string, params ...string) {
	req := s.HTTPRequest("GET", nil, action, params)
	s.PrintCurlCommand(req, params)
	s.ExecuteRequest(req, &data)
}

// PostRequest sends a HTTP POST request to the API and returns the response. If
// one of the request parameters is prefixed with an AT symbol the method will
// assume that the user is trying to load a local file, it will proceed to check
// this by locating such file in the disk, then will attach the data into the
// HTTP request object and upload it to the API. Alternatively, if the file does
// not exists, the method will send the parameter with the apparent filename as
// a string value.
func (s *SlackAPI) PostRequest(data interface{}, action string, params ...string) {
	var parts []string
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Append more HTTP request params.
	for _, keyvalue := range params {
		if keyvalue == TOKEN {
			keyvalue += "=" + s.Token
		}

		parts = strings.SplitN(keyvalue, "=", 2)

		/* Ignore empty parameters */
		if len(parts) != 2 {
			continue
		}

		/* Check if the parameter is referencing a file */
		isfile, fpath, fname := s.CheckFileReference(parts[1])

		if !isfile {
			fwriter, _ := writer.CreateFormField(parts[0])
			fwriter.Write([]byte(parts[1]))
			continue
		}

		if len(parts[1]) > 0 && parts[1][0] == '@' {
		}

		/* Read referenced file and attach to the request */
		resource, err := os.Open(fpath)
		if err != nil {
			s.ReportError(err)
		}
		defer resource.Close()
		fwriter, _ := writer.CreateFormFile(parts[0], fpath)
		io.Copy(fwriter, resource) /* attach file data */
		fwriter, _ = writer.CreateFormField("filename")
		fwriter.Write([]byte(fname))
	}

	writer.Close()

	// Now that you have a form, you can submit it to your handler.
	req := s.HTTPRequest("POST", &buffer, action, nil)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	s.PrintCurlCommand(req, params)
	s.ExecuteRequest(req, &data)
}

// CheckFileReference checks if a HTTP request parameter is a file reference.
func (s *SlackAPI) CheckFileReference(text string) (bool, string, string) {
	if len(text) < 2 || text[0] != '@' {
		return false, "", ""
	}

	fpath := text[1:len(text)]

	/* Check if the file actually exists */
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return false, "", ""
	}

	parts := strings.Split(fpath, "/")
	fname := parts[len(parts)-1]

	return true, fpath, fname
}

// PrintFormattedJSON prints the JSON-encoded data as a string.
func (s *SlackAPI) PrintFormattedJSON(data interface{}) {
	response, err := json.MarshalIndent(data, "", "\x20\x20")

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
}

// PrintInlineJSON prints the JSON-encoded data as a formatted string.
func (s *SlackAPI) PrintInlineJSON(data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		s.ReportError(err)
	}

	fmt.Printf("%s\n", response)
}

// PrintAndExit prints the JSON-encoded data as a string and exits.
func (s *SlackAPI) PrintAndExit(data interface{}) {
	s.PrintFormattedJSON(data)
	os.Exit(0)
}

// ReportError prints an error message and exits.
func (s *SlackAPI) ReportError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
