package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	const resJson = `{"object":"user","id":"721363f4-bc34-4700-85de-c4cca6c423ad","name":"go-notion-api-client","avatar_url":null,"type":"bot","bot":{"owner":{"type":"workspace","workspace":true}}}`
	resBytes := []byte(resJson)
	var actual *User

	err := json.Unmarshal(resBytes, &actual)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.notion.com/v1/users/me",
		httpmock.NewBytesResponder(200, resBytes),
	)

	client, err := NewClient("token", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.GetMe(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Get me endpoint", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, actual, res)
	})
}

func RetrieveUserTestFunc(resJson string, userID string) (expected User, res *User, err error) {
	resBytes := []byte(resJson)

	err = json.Unmarshal(resBytes, &expected)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.notion.com/v1/users/"+userID,
		httpmock.NewBytesResponder(200, resBytes),
	)

	client, err := NewClient("token", nil)
	res, err = client.RetrieveUser(context.Background(), userID)
	return expected, res, err
}

func TestRetrieveUser(t *testing.T) {
	var userID string
	t.Run("Endpoint: Retrieve a user with person type", func(t *testing.T) {
		var expected User
		userID = "9a26468a-dad1-498a-becb-3eb19be24f0b"
		var resJson = fmt.Sprintf(`{
			"object": "user",
			"id": "%s",
			"name": "小池智哉",
			"avatar_url": "https://s3-us-west-2.amazonaws.com/public.notion-static.com/a402c784-af23-43fb-b572-09a9d7533f59/IMG_1984_(1).jpg",
			"type": "person",
			"person": {
				"email": "tommy@p.u-tokyo.ac.jp"
			}
		}`, userID)
		expected, actual, err := RetrieveUserTestFunc(resJson, userID)
		if err != nil {
			log.Fatal(err)
		}
		assert.Nil(t, err)
		assert.Equal(t, &expected, actual)
	})

	// t.Run("Endpoint: Retrieve a user with bot type", func(t *testing.T) {
	// 	var expected User
	// 	userID = "721363f4-bc34-4700-85de-c4cca6c423ad"
	// 	resJson := fmt.Sprintf(`{
	// 		"object": "user",
	// 		"id": "%s",
	// 		"name": "go-notion-api-client",
	// 		"avatar_url": null,
	// 		"type": "bot",
	// 		"bot": {
	// 			"owner": {
	// 				"type": "workspace",
	// 				"workspace": true
	// 			}
	// 		}
	// 	}`, userID)
	// 	expected, actual, err := RetrieveUserTestFunc(resJson, userID)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, expected, actual)
	// })
}
