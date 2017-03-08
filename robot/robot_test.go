package robot_test

import (
	"testing"

	"github.com/ggicci/dingtalk/robot"
)

var rb = robot.New("test", "Add webhook here to text!")

func TestSendTextMessage(t *testing.T) {
	m := rb.NewTextMessage()
	m.SetText("喂~~~服务器扛不住啦，再不来看看我我就离家出走啦~")
	m.AtMobiles("18600000001")
	m.AtAll(true)
	err := m.Send()
	if err != nil {
		t.Error(err)
	}
}

func TestSendLinkMessage(t *testing.T) {
	err := rb.NewLinkMessage().
		SetTitle("谢谢你长那么帅还关注我").
		SetText("人漂亮了就说整容了，那么长得丑就毁过容吗？").
		SetPictureURL("https://ggicci.me/content/images/2016/03/wechat.jpg").
		SetMessageURL("http://ggicci.me").
		Send()
	if err != nil {
		t.Error(err)
	}
}

func TestSendMarkdownMessage(t *testing.T) {
	down := `
## 弗兰兹·卡夫卡

那是饼干吗？🍪

> 不，不是。那是？

- 小
- 饼
- 干
  `
	err := rb.NewMarkdownMessage().
		SetTitle("前前前世").
		SetMarkdown(down).
		Send()
	if err != nil {
		t.Error(err)
	}
}
