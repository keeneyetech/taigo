package taigo

import (
	"fmt"
	"net/http"
	"time"
)

type UserStoryService struct {
	client *Client
}

type UserStoryListEntry struct {
	AssignedTo          interface{}   `json:"assigned_to"`
	AssignedToExtraInfo interface{}   `json:"assigned_to_extra_info"`
	Attachments         []interface{} `json:"attachments"`
	BacklogOrder        int           `json:"backlog_order"`
	BlockedNote         string        `json:"blocked_note"`
	ClientRequirement   bool          `json:"client_requirement"`
	Comment             string        `json:"comment"`
	CreatedDate         time.Time     `json:"created_date"`
	EpicOrder           interface{}   `json:"epic_order"`
	Epics               interface{}   `json:"epics"`
	ExternalReference   interface{}   `json:"external_reference"`
	FinishDate          interface{}   `json:"finish_date"`
	GeneratedFromIssue  interface{}   `json:"generated_from_issue"`
	ID                  int           `json:"id"`
	IsBlocked           bool          `json:"is_blocked"`
	IsClosed            bool          `json:"is_closed"`
	IsVoter             bool          `json:"is_voter"`
	IsWatcher           bool          `json:"is_watcher"`
	KanbanOrder         int           `json:"kanban_order"`
	Milestone           interface{}   `json:"milestone"`
	MilestoneName       interface{}   `json:"milestone_name"`
	MilestoneSlug       interface{}   `json:"milestone_slug"`
	ModifiedDate        time.Time     `json:"modified_date"`
	OriginIssue         interface{}   `json:"origin_issue"`
	Owner               int           `json:"owner"`
	OwnerExtraInfo      struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"owner_extra_info"`
	Points struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
	} `json:"points"`
	Project          int `json:"project"`
	ProjectExtraInfo struct {
		ID           int         `json:"id"`
		LogoSmallURL interface{} `json:"logo_small_url"`
		Name         string      `json:"name"`
		Slug         string      `json:"slug"`
	} `json:"project_extra_info"`
	Ref             int `json:"ref"`
	SprintOrder     int `json:"sprint_order"`
	Status          int `json:"status"`
	StatusExtraInfo struct {
		Color    string `json:"color"`
		IsClosed bool   `json:"is_closed"`
		Name     string `json:"name"`
	} `json:"status_extra_info"`
	Subject         string          `json:"subject"`
	Tags            [][]interface{} `json:"tags"`
	Tasks           []interface{}   `json:"tasks"`
	TeamRequirement bool            `json:"team_requirement"`
	TotalComments   int             `json:"total_comments"`
	TotalPoints     float64         `json:"total_points"`
	TotalVoters     int             `json:"total_voters"`
	TotalWatchers   int             `json:"total_watchers"`
	TribeGig        interface{}     `json:"tribe_gig"`
	Version         int             `json:"version"`
	Watchers        []interface{}   `json:"watchers"`
}

type UserStory struct {
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
	Attachments       []interface{} `json:"attachments"`
	BacklogOrder      int64         `json:"backlog_order"`
	BlockedNote       string        `json:"blocked_note"`
	BlockedNoteHTML   string        `json:"blocked_note_html"`
	ClientRequirement bool          `json:"client_requirement"`
	Comment           string        `json:"comment"`
	CreatedDate       time.Time     `json:"created_date"`
	Description       string        `json:"description"`
	DescriptionHTML   string        `json:"description_html"`
	EpicOrder         interface{}   `json:"epic_order"`
	Epics             []struct {
		Color   string `json:"color"`
		ID      int    `json:"id"`
		Project struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"project"`
		Ref     int    `json:"ref"`
		Subject string `json:"subject"`
	} `json:"epics"`
	ExternalReference  interface{} `json:"external_reference"`
	FinishDate         time.Time   `json:"finish_date"`
	GeneratedFromIssue interface{} `json:"generated_from_issue"`
	ID                 int         `json:"id"`
	IsBlocked          bool        `json:"is_blocked"`
	IsClosed           bool        `json:"is_closed"`
	IsVoter            bool        `json:"is_voter"`
	IsWatcher          bool        `json:"is_watcher"`
	KanbanOrder        int64       `json:"kanban_order"`
	Milestone          int         `json:"milestone"`
	MilestoneName      string      `json:"milestone_name"`
	MilestoneSlug      string      `json:"milestone_slug"`
	ModifiedDate       time.Time   `json:"modified_date"`
	Neighbors          struct {
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
	OriginIssue    interface{} `json:"origin_issue"`
	Owner          int         `json:"owner"`
	OwnerExtraInfo struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"owner_extra_info"`
	Points struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
	} `json:"points"`
	Project          int `json:"project"`
	ProjectExtraInfo struct {
		ID           int         `json:"id"`
		LogoSmallURL interface{} `json:"logo_small_url"`
		Name         string      `json:"name"`
		Slug         string      `json:"slug"`
	} `json:"project_extra_info"`
	Ref             int `json:"ref"`
	SprintOrder     int `json:"sprint_order"`
	Status          int `json:"status"`
	StatusExtraInfo struct {
		Color    string `json:"color"`
		IsClosed bool   `json:"is_closed"`
		Name     string `json:"name"`
	} `json:"status_extra_info"`
	Subject         string        `json:"subject"`
	Tags            [][]string    `json:"tags"`
	Tasks           []interface{} `json:"tasks"`
	TeamRequirement bool          `json:"team_requirement"`
	TotalComments   int           `json:"total_comments"`
	TotalPoints     float64       `json:"total_points"`
	TotalVoters     int           `json:"total_voters"`
	TotalWatchers   int           `json:"total_watchers"`
	TribeGig        interface{}   `json:"tribe_gig"`
	Version         int           `json:"version"`
	Watchers        []interface{} `json:"watchers"`
}

type UserStoryListOptions struct {
	ProjectID        *int  `url:"project,omitempty"`
	MilestoneID      *int  `url:"milestone,omitempty"`
	MilestoneIsNull  *bool `url:"milestone__isnull,omitempty"`
	StatusId         *int  `url:"status,omitempty"`
	StatusIsArchived *bool `url:"status__is_archived,omitempty"`
	Warchers         *int  `url:"watchers,omitempty"`
	AssignedTo       *int  `url:"assigned_to,omitempty"`
	StatusIsClosed   *bool `url:"status__is_closed,omitempty"`
}

func (s *UserStoryService) List(opts *UserStoryListOptions) ([]UserStoryListEntry, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "userstories", opts, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []UserStoryListEntry
	resp, err := s.client.Do(req, &v)
	return v, resp, err
}

func (s *UserStoryService) Get(ID int) (*UserStory, *http.Response, error) {
	u := fmt.Sprintf("userstories/%d", ID)
	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v UserStory
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *UserStoryService) GetByRef(ref, projectID int) (*UserStory, *http.Response, error) {
	u := fmt.Sprintf("userstories/by_ref?ref=%d&project=%d", ref, projectID)
	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v UserStory
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
