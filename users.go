package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

type GetMeResponse struct {
	Object    string      `json:"object"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	AvatarURL interface{} `json:"avatar_url"`
	Type      string      `json:"type"`
	Bot       struct {
		Owner struct {
			Type string `json:"type"`
			User struct {
				Object    string      `json:"object"`
				ID        string      `json:"id"`
				Name      string      `json:"name"`
				AvatarURL interface{} `json:"avatar_url"`
				Type      string      `json:"type"`
				Person    struct {
					Email string `json:"email"`
				} `json:"person"`
			} `json:"user"`
		} `json:"owner"`
	} `json:"bot"`
}

func (c *Client) GetMe(ctx context.Context) (*GetMeResponse, error) {
	reqURL := *c.URL

	reqURL.Path = path.Join(reqURL.Path, "users", "me")

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)

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

	switch res.StatusCode {
	case http.StatusOK:
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var getMeResponse *GetMeResponse
		if err := json.Unmarshal(bodyBytes, &getMeResponse); err != nil {
			return nil, err
		}

		return getMeResponse, nil
	default:
		return nil, errors.New("Unexpected error")
	}
}
