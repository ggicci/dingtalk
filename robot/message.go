package robot

type MessageType string

const (
	TextMessageType     MessageType = "text"
	LinkMessageType     MessageType = "link"
	MarkdownMessageType MessageType = "markdown"
)

type messagePayload struct {
	Type MessageType `json:"msgtype"`

	// for text
	Text struct {
		Content string `json:"content"`
	} `json:"text"`

	At struct {
		Mobiles []string `json:"atMobiles"`
		IsAtAll bool     `json:"isAtAll"`
	} `json:"at"`

	// for link
	Link struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PictureURL string `json:"picUrl"`
		MessageURL string `json:"messageUrl"`
	} `json:"link"`

	// for markdown
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
}
