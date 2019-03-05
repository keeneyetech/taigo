package taigo

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var (
	client = NewClient(os.Getenv("TAIGA_URL"), os.Getenv("TAIGA_TOKEN"))
)

func assertNoError(t *testing.T, err error, resp *Response) {
	t.Helper()
	if err != nil {
		var b bytes.Buffer
		io.Copy(&b, resp.Body)
		t.Fatalf("error %+v : %s", err, b.String())
	}
}
