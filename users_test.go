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

func requestGetMe() (res *User, err error) {
	client, err := NewClient("token", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err = client.GetMe(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func TestGetMe(t *testing.T) {
	path := "https://api.notion.com/v1/users/me"
	resJson, err := os.ReadFile("./testdata/users/get_me.json")
	if err != nil {
		log.Fatal(err)
	}

	err = registerMock(t, string(resJson), path, http.MethodGet)

	if err != nil {
		log.Fatal(err)
	}

	var expected *User
	err = json.Unmarshal([]byte(resJson), &expected)
	if err != nil {
		log.Fatal(err)
	}

	actual, err := requestGetMe()
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Get me endpoint", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

func requestRetrieveUser(userID string) (res *User, err error) {
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

func TestRetrieveUser(t *testing.T) {
	var userID string
	t.Run("Endpoint: Retrieve a user with person type", func(t *testing.T) {
		userID = "user1"
		path := "https://api.notion.com/v1/users/" + userID

		resJson, err := os.ReadFile("./testdata/users/retrieveUser/user.json")
		if err != nil {
			log.Fatal(err)
		}

		err = registerMock(t, string(resJson), path, http.MethodGet)
		if err != nil {
			log.Fatal(err)
		}

		var expected *User
		err = json.Unmarshal([]byte(resJson), &expected)

		if err != nil {
			log.Fatal(err)
		}

		actual, err := requestRetrieveUser(userID)
		if err != nil {
			log.Fatal(err)
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Endpoint: Retrieve a user with bot type", func(t *testing.T) {
		userID = "bot_user"
		path := "https://api.notion.com/v1/users/" + userID

		resJson, err := os.ReadFile("./testdata/users/retrieveUser/bot.json")
		err = registerMock(t, string(resJson), path, http.MethodGet)
		if err != nil {
			log.Fatal(err)
		}

		var expected *User
		err = json.Unmarshal([]byte(resJson), &expected)
		if err != nil {
			log.Fatal(err)
		}

		actual, err := requestRetrieveUser(userID)
		if err != nil {
			log.Fatal(err)
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
