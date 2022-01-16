package template

import "os"

type GitlabTemplate struct {
	*EnvTemplate
	GitlabEnv map[string]interface{}
}

func NewGitlabTemplate(name string, path string) *GitlabTemplate {
	template := NewEnvTemplate(name, path, nil)
	return &GitlabTemplate{
		template,
		nil,
	}
}

func (g *GitlabTemplate) GetMessage() (tpl string, err error) {
	tpl, err = g.getTemplate()
	if err != nil {
		return "", err
	}
	g.GitlabEnv = g.initGitlabEnvs()
	g.Env = func(s string) (string, error) {
		if val := g.GitlabEnv[s]; val != nil {
			return g.GitlabEnv[s].(string), nil
		}
		return "", nil
	}
	tpl, err = g.replace(tpl)
	return tpl, err
}

func (g *GitlabTemplate) initGitlabEnvs() map[string]interface{} {
	envs := make(map[string]interface{})
	envs["CI_PROJECT_TITLE"] = os.Getenv("CI_PROJECT_TITLE")
	envs["CI_COMMIT_BRANCH"] = os.Getenv("CI_COMMIT_BRANCH")
	envs["CI_COMMIT_MESSAGE"] = os.Getenv("CI_COMMIT_MESSAGE")
	envs["CI_JOB_URL"] = os.Getenv("CI_JOB_URL")
	envs["CI_PROJECT_URL"] = os.Getenv("CI_PROJECT_URL")
	envs["CI_PROJECT_TITLE"] = os.Getenv("CI_PROJECT_TITLE")
	envs["GITLAB_USER_NAME"] = os.Getenv("GITLAB_USER_NAME")
	envs["CI_COMMIT_URL"] = os.Getenv("CI_PROJECT_URL") + "/-/commit/" + os.Getenv("CI_COMMIT_SHA")
	ISO8601 := os.Getenv("CI_COMMIT_TIMESTAMP")
	if ISO8601 != "" {
		envs["CI_COMMIT_TIMESTAMP"] = os.Getenv("CI_COMMIT_TIMESTAMP")
	}
	return envs
}
