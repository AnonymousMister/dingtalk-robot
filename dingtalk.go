package dingtalk

import (
	"errors"
	"fmt"
	"github.com/AnonymousMister/dingtalk-robot/client"
	"github.com/AnonymousMister/dingtalk-robot/template"
	"strings"
)

type (
	DingDing struct {
		Config   *Config
		Template *template.Template
		Gitlab   *Gitlab
		Data     interface{}
	}
	Config struct {
		AccessToken string
		Type        string
		Title       string
		Secret      string
		Mobiles     string
		Link        string
		IsAtALL     bool
	}
	Gitlab struct {
		Status      string
		Avatar      string
		ProjectName string
		Branch      string
		CiEnv       map[string]interface{}
	}
)

func (d *DingDing) Title() {
	if d.Config.Title != "" {
		return
	}
	d.Config.Title = d.Gitlab.Status + "--" + d.Gitlab.ProjectName + "--" + d.Gitlab.Branch
}

func (d *DingDing) Exec() error {
	var err error
	if "" == d.Config.AccessToken {
		msg := "missing DingTalk access token"
		return errors.New(msg)
	}
	mes, err := d.Template.GetMessage()
	if err != nil {
		return err
	}
	d.Title()
	newWebHook := client.New(d.Config.AccessToken)
	if "" != d.Config.Secret {
		newWebHook.Secret = d.Config.Secret
	}
	mobiles := strings.Split(d.Config.Mobiles, ",")
	switch strings.ToLower(d.Config.Type) {
	case "markdown":
		err = newWebHook.SendMarkdownMsg(d.Config.Title, mes, d.Config.IsAtALL, mobiles...)
	case "text":
		err = newWebHook.SendTextMsg(mes, d.Config.IsAtALL, mobiles...)
	case "link":
		err = newWebHook.SendLinkMsg(d.Gitlab.Status, mes, d.Gitlab.Avatar, d.Config.Link)
	default:
		msg := "not support message type"
		err = errors.New(msg)
	}

	if err == nil {
		fmt.Println("send message success!")
	}
	return err
}
