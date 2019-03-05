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
	return c
}

// NewRequest creates an API request. This method can be used to performs
// API request not implemented in this library. Otherwise it should not be
// be used directly.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method, urlStr string, opt interface{}, body interface{}) (*http.Request, error) {
	rel, err := addOptions(urlStr, opt)
	if err != nil {
		return nil, err
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
	if c := resp.StatusCode; c < 200 || c > 299 {
		return newResponse(resp), fmt.Errorf("Server returns status %d", c)
	}

	if v != nil {
		defer resp.Body.Close()
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return newResponse(resp), err
}

type Response struct {
	*http.Response

	// Paginated indicates if pagination is being used for the request
	Paginated bool
	// PaginatedBy holds the number of results per page
	PaginatedBy int
	// PaginationCount holds total number of results
	PaginationCount int
	// PaginationCurrent holds the current page
	PaginationCurrent int
	// PaginationNext holds the next results
	PaginationNext int
	// PaginationPrev holds the previous results
	PaginationPrev int
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
	xPaginationCount   = "x-pagintation-count"
	xPaginationCurrent = "x-pagintation-current"
	xPaginationNext    = "x-pagintation-next"
	xPaginationPrev    = "x-pagintation-prev"
)

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.
func (r *Response) populatePageValues() {
	r.Paginated = r.Header.Get(xPaginated) == "true"
	r.PaginatedBy, _ = strconv.Atoi(r.Header.Get(xPaginatedBy))
	r.PaginationCount, _ = strconv.Atoi(r.Header.Get(xPaginationCount))
	r.PaginationCurrent, _ = strconv.Atoi(r.Header.Get(xPaginationCurrent))
	r.PaginationNext, _ = strconv.Atoi(r.Header.Get(xPaginationNext))
	r.PaginationPrev, _ = strconv.Atoi(r.Header.Get(xPaginationPrev))
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
