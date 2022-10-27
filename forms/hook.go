package forms

import "strings"

//从gitlab的触发事件接收到的请求参数表单
type GitlabHookForm struct {
	Eventname string       `json:"event_name"`
	Ref       string       `json:"ref"`
	Commit    string       `json:"checkout_sha"`
	ProjectID int64        `json:"project_id"`
	Project   *ProjectForm `json:"project"`
}

type ProjectForm struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	GitSSHURL  string `json:"git_ssh_url"`
	GitHTTPURL string `json:"git_http_url"`
}

func (f *GitlabHookForm) TagName() string {
	return strings.TrimPrefix(f.Ref, "refs/tags/")
}
