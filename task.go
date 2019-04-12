package taigo

import (
	"fmt"
	"time"
)

type TaskService struct {
	client *Client
}

type TaskListEntry struct {
	AssignedTo        int             `json:"assigned_to"`
	Attachments       []interface{}   `json:"attachments"`
	BlockedNote       string          `json:"blocked_note"`
	CreatedDate       time.Time       `json:"created_date"`
	DueDate           interface{}     `json:"due_date"`
	DueDateReason     string          `json:"due_date_reason"`
	DueDateStatus     string          `json:"due_date_status"`
	ExternalReference interface{}     `json:"external_reference"`
	FinishedDate      time.Time       `json:"finished_date"`
	ID                int             `json:"id"`
	IsBlocked         bool            `json:"is_blocked"`
	IsClosed          bool            `json:"is_closed"`
	IsIocaine         bool            `json:"is_iocaine"`
	IsVoter           bool            `json:"is_voter"`
	IsWatcher         bool            `json:"is_watcher"`
	Milestone         int             `json:"milestone"`
	MilestoneSlug     string          `json:"milestone_slug"`
	ModifiedDate      time.Time       `json:"modified_date"`
	Owner             int             `json:"owner"`
	Project           int             `json:"project"`
	Ref               int             `json:"ref"`
	Status            int             `json:"status"`
	Subject           string          `json:"subject"`
	Tags              [][]interface{} `json:"tags"`
	TaskboardOrder    int64           `json:"taskboard_order"`
	TotalComments     int             `json:"total_comments"`
	TotalVoters       int             `json:"total_voters"`
	TotalWatchers     int             `json:"total_watchers"`
	UsOrder           int64           `json:"us_order"`
	UserStory         int             `json:"user_story"`
	Version           int             `json:"version"`
	Watchers          []int           `json:"watchers"`
}

type Task struct {
	AssignedTo           int             `json:"assigned_to"`
	Attachments          []interface{}   `json:"attachments"`
	BlockedNote          string          `json:"blocked_note"`
	BlockedNoteHTML      string          `json:"blocked_note_html"`
	Comment              string          `json:"comment"`
	CreatedDate          time.Time       `json:"created_date"`
	Description          string          `json:"description"`
	DescriptionHTML      string          `json:"description_html"`
	DueDate              interface{}     `json:"due_date"`
	DueDateReason        string          `json:"due_date_reason"`
	DueDateStatus        string          `json:"due_date_status"`
	ExternalReference    interface{}     `json:"external_reference"`
	FinishedDate         time.Time       `json:"finished_date"`
	GeneratedUserStories interface{}     `json:"generated_user_stories"`
	ID                   int             `json:"id"`
	IsBlocked            bool            `json:"is_blocked"`
	IsClosed             bool            `json:"is_closed"`
	IsIocaine            bool            `json:"is_iocaine"`
	IsVoter              bool            `json:"is_voter"`
	IsWatcher            bool            `json:"is_watcher"`
	Milestone            int             `json:"milestone"`
	MilestoneSlug        string          `json:"milestone_slug"`
	ModifiedDate         time.Time       `json:"modified_date"`
	Owner                int             `json:"owner"`
	Project              int             `json:"project"`
	Ref                  int             `json:"ref"`
	Status               int             `json:"status"`
	Subject              string          `json:"subject"`
	Tags                 [][]interface{} `json:"tags"`
	TaskboardOrder       int64           `json:"taskboard_order"`
	TotalComments        int             `json:"total_comments"`
	TotalVoters          int             `json:"total_voters"`
	TotalWatchers        int             `json:"total_watchers"`
	UsOrder              int64           `json:"us_order"`
	UserStory            int             `json:"user_story"`
	Version              int             `json:"version"`
	Watchers             []int           `json:"watchers"`
}

type TaskRequest struct {
	Version           int          `json:"version"`
	BlockedNote       *string      `json:"blocked_note,omitempty"`
	Description       *string      `json:"description,omitempty"`
	ExternalReference *interface{} `json:"external_reference,omitempty"`
	IsBlocked         *bool        `json:"is_blocked,omitempty"`
	IsClosed          *bool        `json:"is_closed,omitempty"`
	IsIocaine         *bool        `json:"is_iocaine,omitempty"`
	Milestone         *int         `json:"milestone,omitempty"`
	Project           *int         `json:"project,omitempty"`
	Status            *int         `json:"status,omitempty"`
	Subject           *string      `json:"subject,omitempty"`
	Tags              *[]string    `json:"tags,omitempty"`
	TaskboardOrder    *int         `json:"taskboard_order,omitempty"`
	UsOrder           *int         `json:"us_order,omitempty"`
	UserStory         *int         `json:"user_story,omitempty"`
}

type TaskListOptions struct {
	Project        *int  `url:"project,omitempty"`
	Status         *int  `url:"status,omitempty"`
	Owner          *int  `json:"owner"`
	Type           *int  `json:"type"`
	Watchers       *int  `url:"watchers,omitempty"`
	AssignedTo     *int  `url:"assigned_to,omitempty"`
	StatusIsClosed *bool `url:"status__is_closed,omitempty"`
}

func (s *TaskService) List(opts *TaskListOptions) ([]TaskListEntry, *Response, error) {
	req, err := s.client.NewRequest("GET", "tasks", opts, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []TaskListEntry
	resp, err := s.client.Do(req, &v)
	return v, resp, err
}

func (s *TaskService) Get(ID int) (*Task, *Response, error) {
	u := fmt.Sprintf("tasks/%d", ID)
	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v Task
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *TaskService) GetByRef(project interface{}, ref int) (*Task, *Response, error) {
	req, err := s.client.NewRequest("GET", "tasks/by_ref", getByRefOptions(project, ref), nil)
	if err != nil {
		return nil, nil, err
	}

	var v Task
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *TaskService) Create(body *TaskRequest) (*Task, *Response, error) {
	req, err := s.client.NewRequest("POST", "tasks", nil, body)
	if err != nil {
		return nil, nil, err
	}

	var v Task
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *TaskService) Edit(ID int, body *TaskRequest) (*Task, *Response, error) {
	u := fmt.Sprintf("tasks/%d", ID)
	req, err := s.client.NewRequest("PATCH", u, nil, body)
	if err != nil {
		return nil, nil, err
	}

	var v Task
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
