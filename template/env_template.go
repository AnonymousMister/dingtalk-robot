package template

import (
	"regexp"
)

type EnvTemplate struct {
	*RegexpTemplate
}

func NewEnvTemplate(name string, path string, env Exchange) *EnvTemplate {
	template := NewRegexpTemplate(name, path, regexp.MustCompile(`\[([^\[\]]*)]`), env)
	return &EnvTemplate{
		template,
	}
}
