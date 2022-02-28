package template

import (
	"errors"
	"regexp"
	"strings"
)

type Exchange func(string) (string, error)

type RegexpTemplate struct {
	*Template
	Regexp *regexp.Regexp
	Env    Exchange
}

func NewRegexpTemplate(name string, path string, Regexp *regexp.Regexp, env Exchange) *RegexpTemplate {
	template := NewTemplate(name, path)
	return &RegexpTemplate{
		template,
		Regexp,
		env,
	}
}
func (t *RegexpTemplate) GetMessage() (tpl string, err error) {
	tpl, err = t.getTemplate()
	if err != nil {
		return "", err
	}
	tpl, err = t.replace(tpl)
	return tpl, err
}

func (t *RegexpTemplate) replace(tpl string) (string, error) {
	if t.Env == nil {
		return "", errors.New("not Env ")
	}
	match := t.Regexp.FindAllStringSubmatch(tpl, -1)
	for _, m := range match {
		if ok, _ := t.Env(m[1]); ok != "" {
			tpl = strings.ReplaceAll(tpl, m[0], ok)
		}
	}
	return tpl, nil
}
