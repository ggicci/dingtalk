package robot

import "errors"

type textMessage struct {
	robot   *Robot
	payload messagePayload
}

func (m *textMessage) String() string {
	return m.payload.Text.Content
}

func (m *textMessage) SetText(text string) *textMessage {
	m.payload.Text.Content = text
	return m
}

func (m *textMessage) At(mobile ...string) *textMessage {
	m.payload.At.Mobiles = mobile
	return m
}

func (m *textMessage) AtAll(isAtAll bool) *textMessage {
	m.payload.At.IsAtAll = isAtAll
	return m
}

func (m *textMessage) Send() error {
	if m.robot == nil {
		return errors.New("nil robot")
	}
	m.payload.Type = TextMessageType
	return m.robot.sendMessagePayload(&m.payload)
}
