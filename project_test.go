package taigo

import (
	"fmt"
	"testing"
)

func TestProjectList(t *testing.T) {
	fmt.Println("CLIENT:", client)

	projects, resp, err := client.Project.List()

	assertNoError(t, err, resp)
	fmt.Println(projects)
}
