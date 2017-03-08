package robot

import "errors"

type linkMessage struct {
	robot   *Robot
	payload messagePayload
}

func (m *linkMessage) String() string {
	return "(" + m.payload.Link.Title + ", " + m.payload.Link.Text + ", " + m.payload.Link.PictureURL + "," + m.payload.Link.MessageURL + ")"
}

func (m *linkMessage) SetTitle(title string) *linkMessage {
	m.payload.Link.Title = title
	return m
}

func (m *linkMessage) SetText(text string) *linkMessage {
	m.payload.Link.Text = text
	return m
}

func (m *linkMessage) SetPictureURL(urlstr string) *linkMessage {
	m.payload.Link.PictureURL = urlstr
	return m
}

func (m *linkMessage) SetMessageURL(urlstr string) *linkMessage {
	m.payload.Link.MessageURL = urlstr
	return m
}

func (m *linkMessage) Send() error {
	if m.robot == nil {
		return errors.New("nil robot")
	}
	m.payload.Type = LinkMessageType
	return m.robot.sendMessagePayload(&m.payload)
}
