package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Token string

	Logger *log.Logger
}

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

func NewClient(urlStr, token string, logger *log.Logger) (*Client, error) {
	baseURI, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return nil, err
	}

	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	if logger == nil {
		logger = log.New(os.Stderr, "[LOG]", log.LstdFlags)
	}

	return &Client{
		URL:        baseURI,
		HTTPClient: http.DefaultClient,
		Token:      token,
		Logger:     logger,
	}, nil
}

func (c *Client) GetMe(ctx context.Context) (*GetMeResponse, error) {
	reqURL := *c.URL

	reqURL.Path = path.Join(reqURL.Path, "users", "me")

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization: Bearer %s", c.Token)
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
