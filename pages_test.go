package notion

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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
	const resJson = `
		{
			"object": "page",
			"id": "762e592d-92ca-49ab-ac69-03a3c2f4c9d5",
			"created_time": "2022-04-09T11:29:00.000Z",
			"last_edited_time": "2022-05-29T08:09:00.000Z",
			"created_by": {
					"object": "user",
					"id": "5a8ba6f6-dda9-4256-ac73-f49b5e7cd446"
			},
			"last_edited_by": {
					"object": "user",
					"id": "fc26f6d3-1963-4097-afa0-7fd8e0d84712"
			},
			"cover": null,
			"icon": null,
			"parent": {
					"type": "workspace",
					"workspace": true
			},
			"archived": false,
			"properties": {
					"title": {
							"id": "title",
							"type": "title",
							"title": [
									{
											"type": "text",
											"text": {
													"content": "分担＆ファイル構成",
													"link": null
											},
											"annotations": {
													"bold": false,
													"italic": false,
													"strikethrough": false,
													"underline": false,
													"code": false,
													"color": "default"
											},
											"plain_text": "分担＆ファイル構成",
											"href": null
									}
							]
					}
			},
			"url": "https://www.notion.so/762e592d92ca49abac6903a3c2f4c9d5"
		}
	`
	pageId := "page1"
	// pageId := "762e592d92ca49abac6903a3c2f4c9d5"
	path := "https://api.notion.com/v1/pages/" + pageId

	// err := registerMock(t, resJson, path, http.MethodGet)
	// res, err := os.ReadFile("./testdata/pages/retrieve_page.json")

	// resJson, err := os.ReadFile("./testdata/pages/retrieve_page.json")
	err := registerMock(t, resJson, path, http.MethodGet)
	if err != nil {
		log.Fatal(err)
	}

	var expected *GetRetrievePageResponse
	err = json.Unmarshal([]byte(resJson), &expected)

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
