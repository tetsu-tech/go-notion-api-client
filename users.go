package notion

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"path"
)

type Person struct {
	Email string `json:"email"`
}

type Bot struct {
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
}

type User struct {
	Object    string      `json:"object"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	AvatarURL interface{} `json:"avatar_url"`
	Type      string      `json:"type"`
	Bot       *Bot        `json:"bot"`
	Person    *Person     `json:"person"`
}

func (c *Client) GetMe(ctx context.Context) (*User, error) {
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

		var User *User
		if err := json.Unmarshal(bodyBytes, &User); err != nil {
			return nil, err
		}

		return User, nil
	default:
		return nil, errors.New("unexpected error")
	}
}

func (c *Client) RetrieveUser(ctx context.Context, userID string) (res *User, err error) {
	url := path.Join("users", userID)

	err = c.call(ctx, url, http.MethodGet, nil, res)

	if err != nil {
		return nil, err
	}

	return res, err
}
