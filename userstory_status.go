package taigo

import (
	"net/http"
)

type UserStoryStatusService struct {
	client *Client
}

type UserStoryStatus struct {
	Color      string      `json:"color"`
	ID         int         `json:"id"`
	IsArchived bool        `json:"is_archived"`
	IsClosed   bool        `json:"is_closed"`
	Name       string      `json:"name"`
	Order      int         `json:"order"`
	Project    int         `json:"project"`
	Slug       string      `json:"slug"`
	WipLimit   interface{} `json:"wip_limit"`
}

type UserStoryStatusListOptions struct {
	ProjectID *int `url:"project,omitempty"`
}

func (s *UserStoryStatusService) List(opts *UserStoryStatusListOptions) ([]UserStoryStatus, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "userstory-statuses", opts, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []UserStoryStatus
	resp, err := s.client.Do(req, &v)
	return v, resp, err
}
