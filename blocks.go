package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

// TODO: Todo, Paragraph, HeadingなどのいずれかをBlockに持ちたい
type Block struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	CreatedBy      User   `json:"created_by"`
	LastEditedTime string `json:"last_edited_time"`
	LastEditedBy   User   `json:"last_edited_by"`
	HasChildren    bool   `json:"has_children"`
	Type           string `json:"type"`
	Archived       bool   `json:"archived"`
	ToDo           ToDo   `json:"to_do"`
}

type ToDo struct {
	RichText RichText `json:"rich_text"`
	Checked  bool     `json:"checked"`
	Color    string   `json:"color"`
}

type User struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type RichText []struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}

func (c *Client) RetrieveBlock(ctx context.Context, blockID string) (*Block, error) {
	reqURL := *c.URL

	reqURL.Path = path.Join(reqURL.Path, "blocks", blockID)

	url := path.Join("blocks", blockID)

	req, err := c.ConstructReq(ctx, url, http.MethodGet)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Notion-Version", "2021-08-16")
	req = req.WithContext(ctx)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	log.Println(res.StatusCode)

	switch res.StatusCode {
	case http.StatusOK:
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var block *Block
		if err := json.Unmarshal(bodyBytes, &block); err != nil {
			return nil, err
		}

		return block, nil

	default:
		return nil, errors.New("unexpected error")
	}
}
