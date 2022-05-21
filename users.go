package notion

import (
	"context"
	"encoding/json"
	"errors"
	"io"
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
	req, err := c.ConstructReq(ctx, "users/me", http.MethodGet)
	if err != nil {
		return nil, err
	}

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
		return nil, errors.New("unexpected error")
	}
}

func (c *Client) RetrieveUser(ctx context.Context, userID string) (res any, err error) {
	url := path.Join("users", userID)
	req, err := c.ConstructReq(ctx, url, http.MethodGet)
	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		var r io.Reader = response.Body
		json.NewDecoder(r).Decode(&res)

		return res, nil
	default:
		return nil, errors.New("unexpected error")
	}
}
