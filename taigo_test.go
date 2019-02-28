package taigo

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
)

var (
	client = NewClient(os.Getenv("TAIGA_TOKEN"), os.Getenv("TAIGA_URL"))
)

func assertNoError(t *testing.T, err error, resp *http.Response) {
	t.Helper()
	if err != nil {
		var b bytes.Buffer
		io.Copy(&b, resp.Body)
		t.Fatalf("error %+v : %s", err, b.String())
	}
}
