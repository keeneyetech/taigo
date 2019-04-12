package taigo

type TaskStatusService struct {
	client *Client
}

type TaskStatus struct {
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

type TaskStatusListOptions struct {
	ProjectID *int `url:"project,omitempty"`
}

func (s *TaskStatusService) List(opts *TaskStatusListOptions) ([]TaskStatus, *Response, error) {
	req, err := s.client.NewRequest("GET", "task-statuses", opts, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []TaskStatus
	resp, err := s.client.Do(req, &v)
	return v, resp, err
}
