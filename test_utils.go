package notion

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
)

func registerMock(t *testing.T, resJson string, path string, method string) (expected *User, err error) {
	t.Helper()
	resBytes := []byte(resJson)

	err = json.Unmarshal(resBytes, &expected)
	if err != nil {
		return nil, err
	}

	httpmock.RegisterResponder(method, path,
		httpmock.NewBytesResponder(200, resBytes),
	)

	return expected, nil
}
