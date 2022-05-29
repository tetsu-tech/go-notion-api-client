package notion

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestRetrieveBlock(blockID string) (res *Block, err error) {
	client, err := NewClient("token", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err = client.RetrieveBlock(context.Background(), blockID)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func TestRetrieveBlock(t *testing.T) {
	const blockID = "1a1b50f1-2a36-446c-b6d4-88049f05e096"
	const expectedJson = `{
		"object": "block",
		"id": "1a1b50f1-2a36-446c-b6d4-88049f05e096",
		"created_time": "2022-03-20T12:20:00.000Z",
		"last_edited_time": "2022-03-20T12:20:00.000Z",
		"created_by": {
			"object": "user",
			"id": "5a8ba6f6-dda9-4256-ac73-f49b5e7cd446"
		},
		"last_edited_by": {
			"object": "user",
			"id": "5a8ba6f6-dda9-4256-ac73-f49b5e7cd446"
		},
		"has_children": false,
		"archived": false,
		"type": "paragraph",
		"paragraph": {
			"rich_text": [
				{
					"type": "text",
					"text": {
						"content": "なにしたらいいのか？",
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
					"plain_text": "なにしたらいいのか？",
					"href": null
				}
			],
			"color": "default"
		}
	}
	`
	path := "https://api.notion.com/v1/blocks/" + blockID
	err := registerMock(t, expectedJson, path, http.MethodGet)

	if err != nil {
		log.Fatal(err)
	}

	var expected *Block
	err = json.Unmarshal([]byte(expectedJson), &expected)

	if err != nil {
		log.Fatal(err)
	}

	actual, err := requestRetrieveBlock(blockID)

	t.Run("Retrieve a block endpoint", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

}
