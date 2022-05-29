package notion

import (
	"context"
	"fmt"
	"log"
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
	fmt.Println(pageId)
	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func TestRetrievePage(t *testing.T) {
	// pageId := "page1"
	pageId := "762e592d92ca49abac6903a3c2f4c9d5"
	// path := "https://api.notion.com/v1/pages/" + pageId

	// err := registerMock(t, resJson, path, http.MethodGet)
	// res, err := os.ReadFile("./testdata/pages/retrieve_page.json")

	resJson, err := os.ReadFile("./testdata/pages/retrieve_page.json")

	if err != nil {
		log.Fatal(err)
	}

	actual, err := requestRetrievePage(pageId)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Retrueve Page endpoint", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, resJson, actual)
	})

}
