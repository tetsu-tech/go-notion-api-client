package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func NewClient(token string, logger *log.Logger) (*Client, error) {
	baseURI, err := url.ParseRequestURI("https://api.notion.com/v1")

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

func (c *Client) call(ctx context.Context, apiPath string, method string, queryParams url.Values, postBody interface{}, res interface{}) error {
	var (
		contentType string
		body        io.Reader
	)

	contentType = "application/json"
	jsonParams, err := json.Marshal(postBody)

	if err != nil {
		return err
	}
	body = bytes.NewBuffer(jsonParams)

	req, err := c.newRequest(ctx, apiPath, method, contentType, queryParams, body)
	if err != nil {
		return err
	}

	return c.do(ctx, req, res)
}

func (c *Client) do(
	ctx context.Context,
	req *http.Request,
	res interface{},
) error {
	httpClient := http.DefaultClient
	response, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case http.StatusOK:
		var r io.Reader = response.Body
		json.NewDecoder(r).Decode(&res)

		return nil
	default:
		return errors.New("unexpected error")
	}
}

func (c *Client) newRequest(ctx context.Context, apiPath string, method string, contentType string, queryParams url.Values, body io.Reader) (*http.Request, error) {
	const (
		baseURL    = "https://api.notion.com"
		apiVersion = "v1"
	)
	u, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, apiVersion, apiPath)
	u.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Notion-Version", "2022-02-22")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return req, nil
}
