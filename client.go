package notion

import (
	"context"
	"errors"
	"fmt"
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

func (c *Client) ConstructReq(ctx context.Context, url string) (*http.Request, error) {
	reqURL := *c.URL
	reqURL.Path = path.Join(reqURL.Path, url)

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Notion-Version", "2022-02-22")
	req = req.WithContext(ctx)
	return req, nil
}
