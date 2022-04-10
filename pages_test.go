package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestRetrievePage(t *testing.T) {
	const resJson = `{"object":"page","id":"185fdb91-43bd-42ab-9ded-ca84f171afaf","created_time":"2022-03-20T12:09:00.000Z","last_edited_time":"2022-04-10T02:22:00.000Z","created_by":{"object":"user","id":"5a8ba6f6-dda9-4256-ac73-f49b5e7cd446"},"last_edited_by":{"object":"user","id":"fc26f6d3-1963-4097-afa0-7fd8e0d84712"},"cover":null,"icon":{"type":"file","file":{"url":"https://s3.us-west-2.amazonaws.com/secure.notion-static.com/1b2b5a51-044a-4ced-8010-7e0c21c7ab0e/megamoji.gif?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220410%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220410T022242Z&X-Amz-Expires=3600&X-Amz-Signature=2ffd01924dfaebdab634b3e6e5ee9f7c274032b709d234dd5b8e831c833c8339&X-Amz-SignedHeaders=host&x-id=GetObject","expiry_time":"2022-04-10T03:22:42.425Z"}},"parent":{"type":"workspace","workspace":true},"archived":false,"properties":{"title":{"id":"title","type":"title","title":[{"type":"text","text":{"content":"Goでなんか作る","link":null},"annotations":{"bold":false,"italic":false,"strikethrough":false,"underline":false,"code":false,"color":"default"},"plain_text":"Goでなんか作る","href":null}]}},"url":"https://www.notion.so/Go-185fdb9143bd42ab9dedca84f171afaf"}`
	resBytes := []byte(resJson)
	var actual *GetRetrievePageResponse
	err := json.Unmarshal(resBytes, &actual)

	const pageId = "185fdb9143bd42ab9dedca84f171afaf"

	if err != nil {
		return
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.notion.com/v1/pages/185fdb9143bd42ab9dedca84f171afaf",
		httpmock.NewBytesResponder(200, resBytes),
	)

	client, err := NewClient("https://api.notion.com/v1", "token", nil)

	if err != nil {
		panic(err)
	}

	res, err := client.GetRetrievePage(context.Background(), pageId)

	if err != nil {
		panic(err)
	}

	t.Run("Retrive Page API test", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, actual, res)
	})
}
