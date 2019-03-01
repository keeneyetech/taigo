package taigo

import (
	"fmt"
	"testing"
)

func TestProjectList(t *testing.T) {
	fmt.Println("CLIENT:", client)

	projects, resp, err := client.Project.List()

	assertNoError(t, err, resp)
	for _, p := range projects {
		pp, resp, err := client.Project.Get(p.ID)

		assertNoError(t, err, resp)
		fmt.Println(p.ID, pp.ID)

		pp, resp, err = client.Project.GetBySlug(p.Slug)

		assertNoError(t, err, resp)
		fmt.Println(p.Slug, pp.Slug)
	}
}
