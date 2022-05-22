package notion

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func registerMock(t *testing.T, resJson string, path string, method string) (err error) {
	t.Helper()
	resBytes := []byte(resJson)

	httpmock.RegisterResponder(method, path,
		httpmock.NewBytesResponder(200, resBytes),
	)

	return nil
}
