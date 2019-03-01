package taigo

import (
	"fmt"
	"testing"
)

func TestUserStoryStatusList(t *testing.T) {

	sts, resp, err := client.UserStoryStatus.List(nil)

	assertNoError(t, err, resp)
	for _, st := range sts {
		fmt.Println(st.ID, st.Name, st.Slug, st.Project)
	}

	opts := &UserStoryStatusListOptions{ProjectID: Int(6)}
	sts, resp, err = client.UserStoryStatus.List(opts)

	assertNoError(t, err, resp)
	for _, st := range sts {
		fmt.Println(st.ID, st.Name, st.Slug, st.Project)
	}
}
