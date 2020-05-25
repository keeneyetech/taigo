package taigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

type Client struct {
	authToken string
	URL       *url.URL
	client    *http.Client

	Project         *ProjectService
	UserStory       *UserStoryService
	UserStoryStatus *UserStoryStatusService
	Issue           *IssueService
	IssueStatus     *IssueStatusService
	Task            *TaskService
	TaskStatus      *TaskStatusService
}

func NewClient(URL, authToken string) *Client {
	c := &Client{
		authToken: authToken,
		client:    http.DefaultClient,
	}
	if !strings.HasSuffix(URL, "/") {
		URL += "/"
	}

	var err error
	c.URL, err = url.Parse(URL)
	if err != nil {
		panic(err)
	}

	c.Project = &ProjectService{c}
	c.UserStory = &UserStoryService{c}
	c.UserStoryStatus = &UserStoryStatusService{c}
	c.Issue = &IssueService{c}
	c.IssueStatus = &IssueStatusService{c}
	c.Task = &TaskService{c}
	c.TaskStatus = &TaskStatusService{c}
	return c
}

// NewRequest creates an API request. This method can be used to performs
// API request not implemented in this library. Otherwise it should not be
// be used directly.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method, urlStr string, opt interface{}, body interface{}) (*http.Request, error) {

	addLast := false
	disablePagination := true

	if opt != nil {
		rv := reflect.ValueOf(opt)
		if rv.Type().Kind() != reflect.Ptr {
			// Create a new type of Iface's Type,
			// so we have a pointer to work with
			rv = reflect.New(reflect.TypeOf(opt))
		}
		f := rv.Elem().FieldByName("Page")
		if f.IsValid() && !f.IsNil() {
			// if we have some page parameter
			// transform it for easier usage
			p := f.Elem().Int()
			if p != 0 {
				// for convenience setting *page to 0 also
				// disables pagination
				disablePagination = false
				if p == -1 {
					addLast = true
				}
			}
			if p <= 0 {
				// we want to translate *page = -1 into "last"
				// so we nil the parameter here to replace
				// it later with the string "last" for page
				var x *int
				f.Set(reflect.ValueOf(x))
			}
		}
	}

	rel, err := addOptions(urlStr, opt)
	if err != nil {
		return nil, err
	}

	if addLast {
		// we already removed page -1 and replace it now with "last"
		q := rel.Query()
		q.Set("page", "last")
		rel.RawQuery = q.Encode()
	}

	u := c.URL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if disablePagination {
		// we disable pagination if there is no "page"
		// option for the request or the page was 0
		req.Header.Add("x-disable-pagination", "true")
	}

	req.Header.Add("Authorization", "Bearer "+c.authToken)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// Do performs the request, the json received in the response is decoded
// and stored in the value pointed by v.
// Do can be used to perform the request created with NewRequest, which
// should be used only for API requests not implemented in this library.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		defer resp.Body.Close()
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return response, err
}

type ResponseError struct {
	Status  int
	Message string `json:"_error_message"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("Server returned status:%d message:'%s'", e.Status, e.Message)
}

func IsErrNotFound(err error) bool {
	re, ok := err.(*ResponseError)
	return ok && re.Status == http.StatusNotFound
}

func checkResponse(resp *http.Response) error {
	c := resp.StatusCode
	if c >= http.StatusOK && c <= 299 {
		return nil
	}
	re := &ResponseError{Status: c}
	err := json.NewDecoder(resp.Body).Decode(re)
	if err != nil {
		return err
	}
	return re
}

type Response struct {
	*http.Response

	// Paginated indicates if pagination is being used for the request
	Paginated bool
	// PaginatedBy holds the number of results per page
	PaginatedBy int
	// PaginationCount holds total number of results (not pages)
	PaginationCount int
	// PaginationCurrent holds the current page (starting at 1)
	PaginationCurrent int
	// PaginationNext holds an URL for the next results
	PaginationNext string
	// PaginationPrev holds an URL for the previous results
	PaginationPrev string
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.populatePageValues()
	return response
}

const (
	xPaginated         = "x-paginated"
	xPaginatedBy       = "x-paginated-by"
	xPaginationCount   = "x-pagination-count"
	xPaginationCurrent = "x-pagination-current"
	xPaginationNext    = "x-pagination-next"
	xPaginationPrev    = "x-pagination-prev"
)

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.
func (r *Response) populatePageValues() {
	r.Paginated = r.Header.Get(xPaginated) == "true"
	r.PaginatedBy, _ = strconv.Atoi(r.Header.Get(xPaginatedBy))
	r.PaginationCount, _ = strconv.Atoi(r.Header.Get(xPaginationCount))
	r.PaginationCurrent, _ = strconv.Atoi(r.Header.Get(xPaginationCurrent))
	r.PaginationNext = r.Header.Get(xPaginationNext)
	r.PaginationPrev = r.Header.Get(xPaginationPrev)
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (*url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	if opt == nil {
		return u, nil
	}

	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		// No query string to add
		return u, nil
	}

	qs, err := query.Values(opt)
	if err != nil {
		return nil, err
	}

	u.RawQuery = qs.Encode()
	return u, nil
}

func String(v string) *string { return &v }

func Bool(v bool) *bool { return &v }

func Int(v int) *int { return &v }

func getByRefOptions(project interface{}, ref int) interface{} {
	opts := struct {
		Ref         int     `url:"ref"`
		ProjectID   *int    `url:"project,omitempty"`
		ProjectSlug *string `url:"project__slug,omitempty"`
	}{Ref: ref}
	switch p := project.(type) {
	case int:
		opts.ProjectID = &p
	case string:
		opts.ProjectSlug = &p
	}
	return opts
}
