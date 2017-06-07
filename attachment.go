package slackapi

import "encoding/json"

// Attachment defines the expected data from the JSON-encoded API response.
type Attachment struct {
	Actions       []AttachmentAction `json:"actions,omitempty"`
	AuthorIcon    string             `json:"author_icon,omitempty"`
	AuthorLink    string             `json:"author_link,omitempty"`
	AuthorName    string             `json:"author_name,omitempty"`
	AuthorSubname string             `json:"author_subname,omitempty"`
	CallbackID    string             `json:"callback_id,omitempty"`
	Color         string             `json:"color,omitempty"`
	Fallback      string             `json:"fallback"`
	Fields        []AttachmentField  `json:"fields,omitempty"`
	Footer        string             `json:"footer,omitempty"`
	FooterIcon    string             `json:"footer_icon,omitempty"`
	FromURL       string             `json:"from_url"`
	ID            int                `json:"id"`
	ImageBytes    int                `json:"image_bytes"`
	ImageHeight   int                `json:"image_height"`
	ImageURL      string             `json:"image_url,omitempty"`
	ImageWidth    int                `json:"image_width"`
	MarkdownIn    []string           `json:"mrkdwn_in,omitempty"`
	Pretext       string             `json:"pretext,omitempty"`
	ServiceName   string             `json:"service_name"`
	Text          string             `json:"text"`
	ThumbHeight   int                `json:"thumb_height"`
	ThumbURL      string             `json:"thumb_url,omitempty"`
	ThumbWidth    int                `json:"thumb_width"`
	Timestamp     json.Number        `json:"ts,omitempty"`
	Title         string             `json:"title,omitempty"`
	TitleLink     string             `json:"title_link,omitempty"`
}

// AttachmentField defines one single message attachment field.
type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// AttachmentAction defines a data type to append a button or select box into a
// message that users can interact with. A maximum of five actions may be added
// per attachment. Notice that the properties without the "omitempty" attribute
// are required, the rest are optional.
//
// Valid values for Type are: "button", "select".
// Valid values for Style are: "default", "primary", "danger".
// Default value for the MinQueryLength property is: 1 (one).
// The Options property has a limit of one hundred entries.
type AttachmentAction struct {
	Name            string                        `json:"name"`
	Text            string                        `json:"text"`
	Type            string                        `json:"type"`
	Style           string                        `json:"style,omitempty"`
	Value           string                        `json:"value,omitempty"`
	DataSource      string                        `json:"data_source,omitempty"`
	MinQueryLength  int                           `json:"min_query_length,omitempty"`
	Options         []AttachmentActionOption      `json:"options,omitempty"`
	SelectedOptions []AttachmentActionOption      `json:"selected_options,omitempty"`
	OptionGroups    []AttachmentActionOptionGroup `json:"option_groups,omitempty"`
	Confirm         *ConfirmationField            `json:"confirm,omitempty"`
}

// AttachmentActionOption defines the structure of a single button option.
// Notice that both Text and Value are required properties, and there is a limit
// for the number of characters in the Description which is thirty (30).
type AttachmentActionOption struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
}

// AttachmentActionOptionGroup defines the structure of a single select option.
type AttachmentActionOptionGroup struct {
	Text    string                   `json:"text"`
	Options []AttachmentActionOption `json:"options"`
}

// ConfirmationField defines the information for the prompt boxes.
type ConfirmationField struct {
	Text        string `json:"text"`
	Title       string `json:"title,omitempty"`
	OkText      string `json:"ok_text,omitempty"`
	DismissText string `json:"dismiss_text,omitempty"`
}

// AttachmentActionCallback defines the structure of the data sent by Slack to
// your webhook when a user or bot clicks a button or selects an option from the
// select box in interactive messages.
type AttachmentActionCallback struct {
	Token            string             `json:"token"`
	CallbackID       string             `json:"callback_id"`
	ResponseURL      string             `json:"response_url"`
	AttachmentID     string             `json:"attachment_id"`
	ActionTimestamp  string             `json:"action_ts"`
	MessageTimestamp string             `json:"message_ts"`
	Actions          []AttachmentAction `json:"actions"`
	OriginalMessage  Message            `json:"original_message"`
	Channel          Channel            `json:"channel"`
	Team             Team               `json:"team"`
	User             User               `json:"user"`
}
