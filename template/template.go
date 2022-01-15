package template

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Template struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

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
