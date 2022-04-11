package notion

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Token string

	Logger *log.Logger
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
