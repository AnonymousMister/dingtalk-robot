package dingtalk

import (
	"errors"
	"fmt"
	"github.com/AnonymousMister/dingtalk-robot/client"
	"github.com/AnonymousMister/dingtalk-robot/template"
	"strings"
)

type DingDing struct {
	AccessToken string
	Type        string
	Title       string
	Secret      string
	Mobiles     string
	Link        string
	IsAtALL     bool
	PicURL      string
	Template    template.T
}
type Config struct {
	AccessToken string
	Type        string
	Title       string
	Secret      string
	Mobiles     string
	Link        string
	IsAtALL     bool
}

func (d *DingDing) check() error {
	return nil
}

func (d *DingDing) Exec() error {
	if e := d.check(); e != nil {
		return e
	}
	mes, err := d.Template.GetMessage()
	if err != nil {
		return err
	}
	robotClient := client.New(d.AccessToken)
	if "" != d.Secret {
		robotClient.Secret = d.Secret
	}
	mobiles := strings.Split(d.Mobiles, ",")
	switch strings.ToLower(d.Type) {
	case "markdown":
		err = robotClient.SendMarkdownMsg(d.Title, mes, d.IsAtALL, mobiles...)
	case "text":
		err = robotClient.SendTextMsg(mes, d.IsAtALL, mobiles...)
	case "link":
		err = robotClient.SendLinkMsg(d.Title, mes, d.PicURL, d.Link)
	default:
		msg := "not support message type"
		err = errors.New(msg)
	}

	if err == nil {
		fmt.Println("send message success!")
	}
	return err
}
