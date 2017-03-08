package robot

import "errors"

type markdownMessage struct {
	robot   *Robot
	payload messagePayload
}

func (m *markdownMessage) String() string {
	return "MarkdownMessage(" + m.payload.Markdown.Title + "," + m.payload.Markdown.Text + ")"
}

func (m *markdownMessage) SetTitle(title string) *markdownMessage {
	m.payload.Markdown.Title = title
	return m
}

func (m *markdownMessage) SetMarkdown(markdown string) *markdownMessage {
	m.payload.Markdown.Text = markdown
	return m
}

func (m *markdownMessage) Send() error {
	if m.robot == nil {
		return errors.New("nil robot")
	}
	m.payload.Type = MarkdownMessageType
	return m.robot.sendMessagePayload(&m.payload)
}
