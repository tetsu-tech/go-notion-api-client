package notion

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestRetrievePage(pageId string) (res *GetRetrievePageResponse, err error) {
	client, err := NewClient("token", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err = client.GetRetrievePage(context.Background(), pageId)
	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func TestRetrievePage(t *testing.T) {
	pageId := "page1"
	path := "https://api.notion.com/v1/pages/" + pageId

	res, err := os.ReadFile("./testdata/pages/retrieve_page.json")

	if err != nil {
		log.Fatal(err)
	}

	err = registerMock(t, string(res), path, http.MethodGet)
	if err != nil {
		log.Fatal(err)
	}

	var expected *GetRetrievePageResponse
	err = json.Unmarshal([]byte(res), &expected)

	if err != nil {
		log.Fatal(err)
	}

	actual, err := requestRetrievePage(pageId)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Retrueve Page endpoint", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
