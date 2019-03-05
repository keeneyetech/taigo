package taigo

import (
	"fmt"
	"time"
)

type IssueService struct {
	client *Client
}

type IssueListEntry struct {
	AssignedTo          interface{} `json:"assigned_to"`
	AssignedToExtraInfo interface{} `json:"assigned_to_extra_info"`
	BlockedNote         string      `json:"blocked_note"`
	CreatedDate         time.Time   `json:"created_date"`
	ExternalReference   interface{} `json:"external_reference"`
	FinishedDate        interface{} `json:"finished_date"`
	ID                  int         `json:"id"`
	IsBlocked           bool        `json:"is_blocked"`
	IsClosed            bool        `json:"is_closed"`
	IsVoter             bool        `json:"is_voter"`
	IsWatcher           bool        `json:"is_watcher"`
	Milestone           interface{} `json:"milestone"`
	ModifiedDate        time.Time   `json:"modified_date"`
	Owner               int         `json:"owner"`
	OwnerExtraInfo      struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"owner_extra_info"`
	Priority         int `json:"priority"`
	Project          int `json:"project"`
	ProjectExtraInfo struct {
		ID           int         `json:"id"`
		LogoSmallURL interface{} `json:"logo_small_url"`
		Name         string      `json:"name"`
		Slug         string      `json:"slug"`
	} `json:"project_extra_info"`
	Ref             int `json:"ref"`
	Severity        int `json:"severity"`
	Status          int `json:"status"`
	StatusExtraInfo struct {
		Color    string `json:"color"`
		IsClosed bool   `json:"is_closed"`
		Name     string `json:"name"`
	} `json:"status_extra_info"`
	Subject       string          `json:"subject"`
	Tags          [][]interface{} `json:"tags"`
	TotalVoters   int             `json:"total_voters"`
	TotalWatchers int             `json:"total_watchers"`
	Type          int             `json:"type"`
	Version       int             `json:"version"`
	Watchers      []int           `json:"watchers"`
}

type Issue struct {
	AssignedTo          int `json:"assigned_to"`
	AssignedToExtraInfo struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"assigned_to_extra_info"`
	BlockedNote          string      `json:"blocked_note"`
	BlockedNoteHTML      string      `json:"blocked_note_html"`
	Comment              string      `json:"comment"`
	CreatedDate          time.Time   `json:"created_date"`
	Description          string      `json:"description"`
	DescriptionHTML      string      `json:"description_html"`
	ExternalReference    interface{} `json:"external_reference"`
	FinishedDate         interface{} `json:"finished_date"`
	GeneratedUserStories interface{} `json:"generated_user_stories"`
	ID                   int         `json:"id"`
	IsBlocked            bool        `json:"is_blocked"`
	IsClosed             bool        `json:"is_closed"`
	IsVoter              bool        `json:"is_voter"`
	IsWatcher            bool        `json:"is_watcher"`
	Milestone            interface{} `json:"milestone"`
	ModifiedDate         time.Time   `json:"modified_date"`
	Neighbors            struct {
		Next struct {
			ID      int    `json:"id"`
			Ref     int    `json:"ref"`
			Subject string `json:"subject"`
		} `json:"next"`
		Previous struct {
			ID      int    `json:"id"`
			Ref     int    `json:"ref"`
			Subject string `json:"subject"`
		} `json:"previous"`
	} `json:"neighbors"`
	Owner          int `json:"owner"`
	OwnerExtraInfo struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"owner_extra_info"`
	Priority         int `json:"priority"`
	Project          int `json:"project"`
	ProjectExtraInfo struct {
		ID           int         `json:"id"`
		LogoSmallURL interface{} `json:"logo_small_url"`
		Name         string      `json:"name"`
		Slug         string      `json:"slug"`
	} `json:"project_extra_info"`
	Ref             int `json:"ref"`
	Severity        int `json:"severity"`
	Status          int `json:"status"`
	StatusExtraInfo struct {
		Color    string `json:"color"`
		IsClosed bool   `json:"is_closed"`
		Name     string `json:"name"`
	} `json:"status_extra_info"`
	Subject       string          `json:"subject"`
	Tags          [][]interface{} `json:"tags"`
	TotalVoters   int             `json:"total_voters"`
	TotalWatchers int             `json:"total_watchers"`
	Type          int             `json:"type"`
	Version       int             `json:"version"`
	Watchers      []int           `json:"watchers"`
}

type IssueRequest struct {
	Version         int       `json:"version"`
	Status          *int      `json:"status,omitempty"`
	Project         *int      `json:"project,omitempty"`
	Subject         *string   `json:"subject,omitempty"`
	Description     *string   `json:"description,omitempty"`
	AssignedTo      *int      `json:"assigned_to,omitempty"`
	BlockedNote     *string   `json:"blocked_note,omitempty"`
	IsBlocked       *bool     `json:"is_blocked,omitempty"`
	IsClosed        *bool     `json:"is_closed,omitempty"`
	KanbanOrder     *int      `json:"kanban_order,omitempty"`
	Milestone       *int      `json:"milestone,omitempty"`
	Severity        *int      `json:"severity,omitempty"`
	Priority        *int      `json:"priority,omitempty"`
	Type            *int      `json:"type,omitempty"`
	Tags            *[]string `json:"tags,omitempty"`
	TeamRequirement *bool     `json:"team_requirement,omitempty"`
	Watchers        *[]int    `json:"watchers,omitempty"`
}

type IssueListOptions struct {
	Project        *int  `url:"project,omitempty"`
	Status         *int  `url:"status,omitempty"`
	Severity       *int  `json:"severity"`
	Priority       *int  `json:"priority"`
	Owner          *int  `json:"owner"`
	Type           *int  `json:"type"`
	Watchers       *int  `url:"watchers,omitempty"`
	AssignedTo     *int  `url:"assigned_to,omitempty"`
	StatusIsClosed *bool `url:"status__is_closed,omitempty"`
}

func (s *IssueService) List(opts *IssueListOptions) ([]IssueListEntry, *Response, error) {
	req, err := s.client.NewRequest("GET", "issues", opts, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []IssueListEntry
	resp, err := s.client.Do(req, &v)
	return v, resp, err
}

func (s *IssueService) Get(ID int) (*Issue, *Response, error) {
	u := fmt.Sprintf("issues/%d", ID)
	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v Issue
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *IssueService) GetByRef(project interface{}, ref int) (*Issue, *Response, error) {
	req, err := s.client.NewRequest("GET", "issues/by_ref", getByRefOptions(project, ref), nil)
	if err != nil {
		return nil, nil, err
	}

	var v Issue
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *IssueService) Create(body *IssueRequest) (*Issue, *Response, error) {
	req, err := s.client.NewRequest("POST", "issues", nil, body)
	if err != nil {
		return nil, nil, err
	}

	var v Issue
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *IssueService) Edit(ID int, body *IssueRequest) (*Issue, *Response, error) {
	u := fmt.Sprintf("issues/%d", ID)
	req, err := s.client.NewRequest("PATCH", u, nil, body)
	if err != nil {
		return nil, nil, err
	}

	var v Issue
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
