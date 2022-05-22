package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func requestUser(userID string) (res *User, err error) {
	client, err := NewClient("token", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err = client.RetrieveUser(context.Background(), userID)
	if err != nil {
		log.Fatal(err)
	}
	return res, err
}

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	status := m.Run()

	os.Exit(status)
}

func TestRetrieveUser(t *testing.T) {
	var userID string
	t.Run("Endpoint: Retrieve a user with person type", func(t *testing.T) {
		userID = "user1"
		var resJson = fmt.Sprintf(`{
			"object": "user",
			"id": "%s",
			"name": "user1",
			"avatar_url": "user1_avatar",
			"type": "person",
			"person": {
				"email": "user1@example.com"
			}
		}`, userID)

		path := "https://api.notion.com/v1/users/" + userID
		expected, err := registerMock(t, resJson, path, http.MethodGet)

		if err != nil {
			log.Fatal(err)
		}

		actual, err := requestUser(userID)
		if err != nil {
			log.Fatal(err)
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Endpoint: Retrieve a user with bot type", func(t *testing.T) {
		userID = "bot_user"
		resJson := fmt.Sprintf(`{
			"object": "user",
			"id": "%s",
			"name": "bot user",
			"avatar_url": null,
			"type": "bot",
			"bot": {
				"owner": {
					"type": "workspace",
					"workspace": true
				}
			}
		}`, userID)
		path := "https://api.notion.com/v1/users/" + userID
		expected, err := registerMock(t, resJson, path, http.MethodGet)

		if err != nil {
			log.Fatal(err)
		}

		actual, err := requestUser(userID)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
