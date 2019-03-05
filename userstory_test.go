package taigo

import (
	"fmt"
	"testing"
)

func TestUserStory(t *testing.T) {

	opts := &UserStoryListOptions{}

	uss, resp, err := client.UserStory.List(opts)

	assertNoError(t, err, resp)
	for _, us := range uss {
		fmt.Println(us.ID, us.Ref, us.Project, us.Subject, us.Status, us.StatusExtraInfo.Name)
	}

	opts.StatusIsClosed = Bool(true)

	uss, resp, err = client.UserStory.List(opts)

	assertNoError(t, err, resp)
	for _, us := range uss {
		fmt.Println(us.ID, us.Ref, us.Project, us.Subject, us.Status, us.StatusExtraInfo.Name)
	}

	opts.StatusIsClosed = Bool(false)
	opts.Status = Int(91)

	uss, resp, err = client.UserStory.List(opts)

	assertNoError(t, err, resp)
	for _, us := range uss {
		fmt.Println(us.ID, us.Ref, us.Project, us.Subject, us.Status, us.StatusExtraInfo.Name)
	}

	us, resp, err := client.UserStory.Get(uss[0].ID)

	assertNoError(t, err, resp)
	fmt.Println(uss[0].Subject, us.Subject)

	us, resp, err = client.UserStory.GetByRef(uss[0].Project, uss[0].Ref)

	assertNoError(t, err, resp)
	fmt.Println(uss[0].Subject, us.Subject)

	us, resp, err = client.UserStory.GetByRef(uss[0].ProjectExtraInfo.Slug, uss[0].Ref)

	assertNoError(t, err, resp)
	fmt.Println(uss[0].Subject, us.Subject)
}
