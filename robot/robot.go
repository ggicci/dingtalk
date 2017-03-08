package robot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	httpClient http.Client
)

type Robot struct {
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
}

func New(name, webhook string) *Robot {
	return &Robot{
		Name:    name,
		Webhook: webhook,
	}
}

func (r *Robot) String() string {
	return fmt.Sprintf("DingtalkRobot#[%s](%s)", r.Name, r.Webhook)
}

func (r *Robot) NewTextMessage() *textMessage {
	return &textMessage{robot: r}
}

func (r *Robot) NewLinkMessage() *linkMessage {
	return &linkMessage{robot: r}
}

func (r *Robot) NewMarkdownMessage() *markdownMessage {
	return &markdownMessage{robot: r}
}

func (r *Robot) sendMessagePayload(payload *messagePayload) error {
	bs, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := httpClient.Post(r.Webhook, "application/json", bytes.NewReader(bs))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var sr serverResponse
	responseTextBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(responseTextBytes, &sr); err != nil {
		return errors.New(err.Error() + " :" + string(responseTextBytes))
	}
	if sr.ErrorCode != 0 {
		return errors.New(string(responseTextBytes))
	}
	return nil
}

type serverResponse struct {
	ErrorCode int `json:"errcode"` // just check error code
}
