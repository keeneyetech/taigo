package taigo

import (
	"fmt"
	"net/http"
	"time"
)

type ProjectService struct {
	client *Client
}

type ProjectListEntry struct {
	ID                   int           `json:"id"`
	AnonPermissions      []interface{} `json:"anon_permissions"`
	BlockedCode          interface{}   `json:"blocked_code"`
	CreatedDate          time.Time     `json:"created_date"`
	CreationTemplate     int           `json:"creation_template"`
	DefaultEpicStatus    int           `json:"default_epic_status"`
	DefaultIssueStatus   int           `json:"default_issue_status"`
	DefaultIssueType     int           `json:"default_issue_type"`
	DefaultPoints        int           `json:"default_points"`
	DefaultPriority      int           `json:"default_priority"`
	DefaultSeverity      int           `json:"default_severity"`
	DefaultTaskStatus    int           `json:"default_task_status"`
	DefaultUsStatus      int           `json:"default_us_status"`
	Description          string        `json:"description"`
	IAmAdmin             bool          `json:"i_am_admin"`
	IAmMember            bool          `json:"i_am_member"`
	IAmOwner             bool          `json:"i_am_owner"`
	IsBacklogActivated   bool          `json:"is_backlog_activated"`
	IsContactActivated   bool          `json:"is_contact_activated"`
	IsEpicsActivated     bool          `json:"is_epics_activated"`
	IsFan                bool          `json:"is_fan"`
	IsFeatured           bool          `json:"is_featured"`
	IsIssuesActivated    bool          `json:"is_issues_activated"`
	IsKanbanActivated    bool          `json:"is_kanban_activated"`
	IsLookingForPeople   bool          `json:"is_looking_for_people"`
	IsPrivate            bool          `json:"is_private"`
	IsWatcher            bool          `json:"is_watcher"`
	IsWikiActivated      bool          `json:"is_wiki_activated"`
	LogoBigURL           interface{}   `json:"logo_big_url"`
	LogoSmallURL         interface{}   `json:"logo_small_url"`
	LookingForPeopleNote string        `json:"looking_for_people_note"`
	Members              []int         `json:"members"`
	ModifiedDate         time.Time     `json:"modified_date"`
	MyPermissions        []string      `json:"my_permissions"`
	Name                 string        `json:"name"`
	NotifyLevel          int           `json:"notify_level"`
	Owner                struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"owner"`
	PublicPermissions []interface{} `json:"public_permissions"`
	Slug              string        `json:"slug"`
	Tags              []interface{} `json:"tags"`
	TagsColors        struct {
	} `json:"tags_colors"`
	TotalActivity             int         `json:"total_activity"`
	TotalActivityLastMonth    int         `json:"total_activity_last_month"`
	TotalActivityLastWeek     int         `json:"total_activity_last_week"`
	TotalActivityLastYear     int         `json:"total_activity_last_year"`
	TotalClosedMilestones     int         `json:"total_closed_milestones"`
	TotalFans                 int         `json:"total_fans"`
	TotalFansLastMonth        int         `json:"total_fans_last_month"`
	TotalFansLastWeek         int         `json:"total_fans_last_week"`
	TotalFansLastYear         int         `json:"total_fans_last_year"`
	TotalMilestones           interface{} `json:"total_milestones"`
	TotalStoryPoints          interface{} `json:"total_story_points"`
	TotalWatchers             int         `json:"total_watchers"`
	TotalsUpdatedDatetime     time.Time   `json:"totals_updated_datetime"`
	Videoconferences          interface{} `json:"videoconferences"`
	VideoconferencesExtraData interface{} `json:"videoconferences_extra_data"`
}

type Project struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	CreatedDate          time.Time `json:"created_date"`
	ModifiedDate         time.Time `json:"modified_date"`
	Description          string    `json:"description"`
	MyPermissions        []string  `json:"my_permissions"`
	EpicCustomAttributes []struct {
		CreatedDate  time.Time `json:"created_date"`
		Description  string    `json:"description"`
		ID           int       `json:"id"`
		ModifiedDate time.Time `json:"modified_date"`
		Name         string    `json:"name"`
		Order        int       `json:"order"`
		ProjectID    int       `json:"project_id"`
		Type         string    `json:"type"`
	} `json:"epic_custom_attributes"`
	EpicStatuses []struct {
		Color     string `json:"color"`
		ID        int    `json:"id"`
		IsClosed  bool   `json:"is_closed"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
		ProjectID int    `json:"project_id"`
		Slug      string `json:"slug"`
	} `json:"epic_statuses"`
	IAmAdmin           bool `json:"i_am_admin"`
	IAmMember          bool `json:"i_am_member"`
	IAmOwner           bool `json:"i_am_owner"`
	IsBacklogActivated bool `json:"is_backlog_activated"`
	IsContactActivated bool `json:"is_contact_activated"`
	IsEpicsActivated   bool `json:"is_epics_activated"`
	IsFan              bool `json:"is_fan"`
	IsFeatured         bool `json:"is_featured"`
	IsIssuesActivated  bool `json:"is_issues_activated"`
	IsKanbanActivated  bool `json:"is_kanban_activated"`
	IsLookingForPeople bool `json:"is_looking_for_people"`
	IsOutOfOwnerLimits bool `json:"is_out_of_owner_limits"`
	IsPrivate          bool `json:"is_private"`
	IsPrivateExtraInfo struct {
		CanBeUpdated bool        `json:"can_be_updated"`
		Reason       interface{} `json:"reason"`
	} `json:"is_private_extra_info"`
	IsWatcher             bool `json:"is_watcher"`
	IsWikiActivated       bool `json:"is_wiki_activated"`
	IssueCustomAttributes []struct {
		CreatedDate  time.Time `json:"created_date"`
		Description  string    `json:"description"`
		ID           int       `json:"id"`
		ModifiedDate time.Time `json:"modified_date"`
		Name         string    `json:"name"`
		Order        int       `json:"order"`
		ProjectID    int       `json:"project_id"`
		Type         string    `json:"type"`
	} `json:"issue_custom_attributes"`
	IssueStatuses []struct {
		Color     string `json:"color"`
		ID        int    `json:"id"`
		IsClosed  bool   `json:"is_closed"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
		ProjectID int    `json:"project_id"`
		Slug      string `json:"slug"`
	} `json:"issue_statuses"`
	IssueTypes []struct {
		Color     string `json:"color"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
		ProjectID int    `json:"project_id"`
	} `json:"issue_types"`
	IssuesCsvUUID        interface{} `json:"issues_csv_uuid"`
	LogoBigURL           string      `json:"logo_big_url"`
	LogoSmallURL         string      `json:"logo_small_url"`
	LookingForPeopleNote string      `json:"looking_for_people_note"`
	MaxMemberships       interface{} `json:"max_memberships"`
	Members              []struct {
		Color           string      `json:"color"`
		FullName        string      `json:"full_name"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Role            int         `json:"role"`
		RoleName        string      `json:"role_name"`
		Username        string      `json:"username"`
	} `json:"members"`
	Milestones []struct {
		Closed bool   `json:"closed"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
	} `json:"milestones"`
	NotifyLevel int `json:"notify_level"`
	Owner       struct {
		BigPhoto        interface{} `json:"big_photo"`
		FullNameDisplay string      `json:"full_name_display"`
		GravatarID      string      `json:"gravatar_id"`
		ID              int         `json:"id"`
		IsActive        bool        `json:"is_active"`
		Photo           interface{} `json:"photo"`
		Username        string      `json:"username"`
	} `json:"owner"`
	Points []struct {
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		Order     int         `json:"order"`
		ProjectID int         `json:"project_id"`
		Value     interface{} `json:"value"`
	} `json:"points"`
	Priorities []struct {
		Color     string `json:"color"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
		ProjectID int    `json:"project_id"`
	} `json:"priorities"`
	PublicPermissions []interface{} `json:"public_permissions"`
	Roles             []struct {
		Computable  bool     `json:"computable"`
		ID          int      `json:"id"`
		Name        string   `json:"name"`
		Order       int      `json:"order"`
		Permissions []string `json:"permissions"`
		ProjectID   int      `json:"project_id"`
		Slug        string   `json:"slug"`
	} `json:"roles"`
	Severities []struct {
		Color     string `json:"color"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
		ProjectID int    `json:"project_id"`
	} `json:"severities"`
	Slug       string        `json:"slug"`
	Tags       []interface{} `json:"tags"`
	TagsColors struct {
	} `json:"tags_colors"`
	TaskCustomAttributes []struct {
		CreatedDate  time.Time `json:"created_date"`
		Description  string    `json:"description"`
		ID           int       `json:"id"`
		ModifiedDate time.Time `json:"modified_date"`
		Name         string    `json:"name"`
		Order        int       `json:"order"`
		ProjectID    int       `json:"project_id"`
		Type         string    `json:"type"`
	} `json:"task_custom_attributes"`
	TaskStatuses []struct {
		Color     string `json:"color"`
		ID        int    `json:"id"`
		IsClosed  bool   `json:"is_closed"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
		ProjectID int    `json:"project_id"`
		Slug      string `json:"slug"`
	} `json:"task_statuses"`
	TasksCsvUUID           interface{} `json:"tasks_csv_uuid"`
	TotalActivity          int         `json:"total_activity"`
	TotalActivityLastMonth int         `json:"total_activity_last_month"`
	TotalActivityLastWeek  int         `json:"total_activity_last_week"`
	TotalActivityLastYear  int         `json:"total_activity_last_year"`
	TotalClosedMilestones  int         `json:"total_closed_milestones"`
	TotalFans              int         `json:"total_fans"`
	TotalFansLastMonth     int         `json:"total_fans_last_month"`
	TotalFansLastWeek      int         `json:"total_fans_last_week"`
	TotalFansLastYear      int         `json:"total_fans_last_year"`
	TotalMemberships       int         `json:"total_memberships"`
	TotalMilestones        int         `json:"total_milestones"`
	TotalStoryPoints       float64     `json:"total_story_points"`
	TotalWatchers          int         `json:"total_watchers"`
	TotalsUpdatedDatetime  time.Time   `json:"totals_updated_datetime"`
	TransferToken          string      `json:"transfer_token"`
	UsStatuses             []struct {
		Color      string      `json:"color"`
		ID         int         `json:"id"`
		IsArchived bool        `json:"is_archived"`
		IsClosed   bool        `json:"is_closed"`
		Name       string      `json:"name"`
		Order      int         `json:"order"`
		ProjectID  int         `json:"project_id"`
		Slug       string      `json:"slug"`
		WipLimit   interface{} `json:"wip_limit"`
	} `json:"us_statuses"`
	UserstoriesCsvUUID        interface{} `json:"userstories_csv_uuid"`
	UserstoryCustomAttributes []struct {
		CreatedDate  time.Time `json:"created_date"`
		Description  string    `json:"description"`
		ID           int       `json:"id"`
		ModifiedDate time.Time `json:"modified_date"`
		Name         string    `json:"name"`
		Order        int       `json:"order"`
		ProjectID    int       `json:"project_id"`
		Type         string    `json:"type"`
	} `json:"userstory_custom_attributes"`
	Videoconferences          interface{} `json:"videoconferences"`
	VideoconferencesExtraData interface{} `json:"videoconferences_extra_data"`
}

func (s *ProjectService) List() ([]ProjectListEntry, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "projects", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []ProjectListEntry
	resp, err := s.client.Do(req, &v)
	return v, resp, err
}

func (s *ProjectService) Get(id int) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects/%d", id)
	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v Project
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}

func (s *ProjectService) GetBySlug(slug string) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects/by_slug?slug=%s", slug)
	req, err := s.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var v Project
	resp, err := s.client.Do(req, &v)
	return &v, resp, err
}
