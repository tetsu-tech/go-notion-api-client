package notion

import (
	"context"
	"net/http"
	"path"
	"time"
)

type GetRetrievePageResponse struct {
	Object         string    `json:"object"`
	ID             string    `json:"id"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"last_edited_by"`
	Cover interface{} `json:"cover"`
	Icon  struct {
		Type string `json:"type"`
		File struct {
			URL        string    `json:"url"`
			ExpiryTime time.Time `json:"expiry_time"`
		} `json:"file"`
	} `json:"icon"`
	Parent struct {
		Type      string `json:"type"`
		Workspace bool   `json:"workspace"`
	} `json:"parent"`
	Archived   bool `json:"archived"`
	Properties struct {
		Title struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Title []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"title"`
		} `json:"title"`
	} `json:"properties"`
	URL string `json:"url"`
}

func (c *Client) GetRetrievePage(ctx context.Context, pageId string) (res *GetRetrievePageResponse, err error) {
	url := path.Join("pages", pageId)

	err = c.call(ctx, url, http.MethodGet, nil, &res)

	if err != nil {
		return nil, err
	}

	return res, err
}
