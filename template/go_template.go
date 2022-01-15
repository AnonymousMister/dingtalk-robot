package template

import (
	"bytes"
	"text/template"
)

type GoTemplate struct {
	Template
	Data interface{}
}

func (t *GoTemplate) GetMessage() (tpl string, err error) {
	tpl, err = t.getTemplate()
	if err != nil {
		return "", err
	}
	tpl, err = t.template(tpl)
	return tpl, err
}

func (t *GoTemplate) template(tpl string) (string, error) {
	tmpl, e := template.New(t.Name).Parse(tpl)
	if e != nil {
		return "", e
	}
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, t.Data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
