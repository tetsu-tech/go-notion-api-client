package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	const resJson = `{"object":"user","id":"721363f4-bc34-4700-85de-c4cca6c423ad","name":"go-notion-api-client","avatar_url":null,"type":"bot","bot":{"owner":{"type":"workspace","workspace":true}}}`
	resBytes := []byte(resJson)
	var actual *GetMeResponse
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

	client, err := NewClient("https://api.notion.com/v1", "token", nil)

	if err != nil {
		panic(err)
	}

	res, err := client.GetMe(context.Background())

	if err != nil {
		panic(err)
	}

	t.Run("sample test", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, actual, res)
	})
}
