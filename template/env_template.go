package template

import (
	"errors"
	"regexp"
	"strings"
)

type EnvTemplate struct {
	*Template
	Env func(string) (string, error)
}

func NewEnvTemplate(name string, path string, env func(string) (string, error)) *EnvTemplate {
	template := NewTemplate(name, path)
	return &EnvTemplate{
		template,
		env,
	}
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
