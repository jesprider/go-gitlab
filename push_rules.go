package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

type PushRules struct {
	ID                 int        `json:"id,omitempty"`
	PID                int        `json:"project_id,omitempty"`
	CommitMessageRegex *string    `json:"commit_message_regex,omitempty"`
	BranchNameRegex    *string    `json:"branch_name_regex,omitempty"`
	DenyDeleteTag      *bool      `json:"deny_delete_tag,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	MemberCheck        *bool      `json:"member_check,omitempty"`
	PreventSecrets     *bool      `json:"prevent_secrets,omitempty"`
	AuthorEmailRegex   *string    `json:"author_email_regex,omitempty"`
	FileNameRegex      *string    `json:"file_name_regex,omitempty"`
	MaxFileSizeMB      *int       `json:"max_file_size,omitempty"`
}

func (s PushRules) String() string {
	return Stringify(s)
}

// GetPushRules gets a specific project's push rules, identified by project ID or
// NAMESPACE/PROJECT_NAME, which is owned by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/projects.html#get-project-push-rules
func (s *ProjectsService) GetPushRule(pid interface{}, options ...OptionFunc) (*PushRules, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(PushRules)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

type AddPushRulesOptions struct {
	CommitMessageRegex *string    `json:"commit_message_regex,omitempty"`
	BranchNameRegex    *string    `json:"branch_name_regex,omitempty"`
	DenyDeleteTag      *bool      `json:"deny_delete_tag,omitempty"`
	MemberCheck        *bool      `json:"member_check,omitempty"`
	PreventSecrets     *bool      `json:"prevent_secrets,omitempty"`
	AuthorEmailRegex   *string    `json:"author_email_regex,omitempty"`
	FileNameRegex      *string    `json:"file_name_regex,omitempty"`
	MaxFileSizeMB      *int       `json:"max_file_size,omitempty"`
}

// AddPushRule adds a push rule to a specified project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#add-project-push-rule
func (s *ProjectsService) AddPushRule(pid interface{}, opt *AddPushRulesOptions, options ...OptionFunc) (*PushRules, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(PushRules)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

type EditPushRulesOptions AddPushRulesOptions

// EditPushRule edits a push rule for a specified project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#edit-project-push-rule
func (s *ProjectsService) EditPushRule(pid interface{}, opt *EditPushRulesOptions, options ...OptionFunc) (*PushRules, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", url.QueryEscape(project))

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(PushRules)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// DeletePushRule removes a push rule from a project. This is an idempotent method and can be called multiple times. Either the push rule is available or not.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/projects.html#delete-project-push-rule
func (s *ProjectsService) DeletePushRule(pid interface{}, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/push_rule", url.QueryEscape(project))

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
