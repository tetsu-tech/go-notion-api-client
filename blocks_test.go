package notion

import (
	"context"
	"log"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

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

	expectedBytes := []byte(expectedJson)
	expected := Block{
		Object:      "block",
		ID:          "1a1b50f1-2a36-446c-b6d4-88049f05e096",
		CreatedTime: "2022-03-20T12:20:00.000Z",
		CreatedBy: User{
			Object: "user",
			ID:     "test_user",
		},
		LastEditedTime: "2022-03-20T12:20:00.000Z",
		LastEditedBy: User{
			Object: "user",
			ID:     "test_user",
		},
		HasChildren: false,
		Type:        "paragraph",
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.notion.com/v1/bocks"+blockID, httpmock.NewBytesResponder(200, expectedBytes))

	client, err := NewClient("token", nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.RetrieveBlock(context.Background(), blockID)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Retrieve a block endpoint", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, resM)
	})

}
