package notion

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
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

type ListAllUserOpts struct {
	StartCursor string `json:"start_cursor"`
	PageSize    int    `json:"page_size"`
}

type ListAllUsersResponse struct {
	Object  string `json:"object"`
	Results []struct {
		Object    string  `json:"object"`
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		AvatarURL string  `json:"avatar_url"`
		Type      string  `json:"type"`
		Person    *Person `json:"person,omitempty"`
		Bot       *Bot    `json:"bot,omitempty"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
	Type       string      `json:"type"`
	User       struct {
	} `json:"user"`
}

func (c *Client) GetMe(ctx context.Context) (res *User, err error) {
	url := path.Join("users", "me")

	err = c.call(ctx, url, http.MethodGet, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) RetrieveUser(ctx context.Context, userID string) (res *User, err error) {
	url := path.Join("users", userID)

	err = c.call(ctx, url, http.MethodGet, nil, nil, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ListAllUsers(ctx context.Context, opts *ListAllUserOpts) (res *ListAllUsersResponse, err error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	err = c.call(ctx, "users", http.MethodGet, v, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
