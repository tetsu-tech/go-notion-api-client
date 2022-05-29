package notion

import (
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	status := m.Run()

	os.Exit(status)
}
