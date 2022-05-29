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

<<<<<<< HEAD
func (c *Client) GetMe(ctx context.Context) (*GetMeResponse, error) {
=======
func (c *Client) GetMe(ctx context.Context) (*User, error) {
>>>>>>> 141b935f9aaf119d74c645a7def0034e97bdf6a7
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

<<<<<<< HEAD
func (c *Client) RetrieveUser(ctx context.Context, userId string) (res any, err error) {
	url := path.Join("users", userId)
	req, err := c.ConstructReq(ctx, url, http.MethodGet)
	if err != nil {
		return nil, err
	}
=======
func (c *Client) RetrieveUser(ctx context.Context, userID string) (res *User, err error) {
	url := path.Join("users", userID)

	err = c.call(ctx, url, http.MethodGet, nil, &res)
>>>>>>> 141b935f9aaf119d74c645a7def0034e97bdf6a7

	if err != nil {
		return nil, err
	}

	return res, err
}
