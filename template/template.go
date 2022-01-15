package template

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type (
	Template struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}
	EnvTemplate struct {
		Template
		Env func(string) (string, error)
	}
	GoTemplate struct {
		Template
		Data interface{}
	}
)

func (t *Template) GetMessage() (tpl string, err error) {
	tpl, err = t.getTemplate()
	if err != nil {
		return "", err
	}
	return tpl, nil
}

func (t *Template) getTemplate() (tpl string, err error) {
	template := fmt.Sprintf("%s/%s", t.Path, strings.ToLower(t.Name))
	if !strings.HasSuffix(t.Name, "md") {
		template = fmt.Sprintf("%s/%s.tpl", t.Path, strings.ToLower(t.Name))
	}
	if !fileExists(template) {
		return "", errors.New("tpl file not exists")
	}
	tplStr, err := ioutil.ReadFile(template)
	if err != nil {
		return "", err
	}
	tpl = string(tplStr)
	return tpl, nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (t *EnvTemplate) GetMessage() (tpl string, err error) {
	tpl, err = t.getTemplate()
	if err != nil {
		return "", err
	}
	tpl, err = t.replace(tpl)
	return tpl, err
}

func (t *EnvTemplate) replace(tpl string) (string, error) {
	if t.Env == nil {
		return "", errors.New("not Env ")
	}
	reg := regexp.MustCompile(`\[([^\[\]]*)]`)
	match := reg.FindAllStringSubmatch(tpl, -1)
	for _, m := range match {
		if ok, _ := t.Env(m[1]); ok != "" {
			tpl = strings.ReplaceAll(tpl, m[0], ok)
		}
	}
	return tpl, nil
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
