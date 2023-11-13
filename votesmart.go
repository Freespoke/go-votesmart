package votesmart

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://api.votesmart.org/"

func New(apiKey string, options ...Option) (Client, error) {
	var opts clientOptions
	for _, opt := range options {
		opts = opt(opts)
	}

	if opts.baseURL == "" {
		opts.baseURL = defaultBaseURL
	}

	if opts.client == nil {
		opts.client = http.DefaultClient
	}

	u, err := url.Parse(opts.baseURL)
	if err != nil {
		return nil, err
	}

	return &client{
		HTTP:    opts.client,
		BaseURL: u,
		APIKey:  apiKey,
	}, nil
}

type Client interface {
	Invoke(ctx context.Context, params *url.Values, dst Method) error
}

type Method interface {
	Method() string
}

type clientOptions struct {
	client  *http.Client
	baseURL string
}

type Option func(clientOptions) clientOptions

func WithClient(v *http.Client) Option {
	return func(o clientOptions) clientOptions {
		o.client = v
		return o
	}
}

func WithBaseURL(v string) Option {
	return func(o clientOptions) clientOptions {
		o.baseURL = v
		return o
	}
}

type client struct {
	HTTP    *http.Client
	BaseURL *url.URL
	APIKey  string
}

func (c *client) Invoke(ctx context.Context, values *url.Values, dst Method) error {
	values.Add("key", c.APIKey)
	values.Add("o", "JSON")
	requestURL := *c.BaseURL
	requestURL.Path = dst.Method()
	requestURL.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return err
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("unexpected response status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var errorval struct {
		Error struct {
			ErrorMessage string `json:"errorMessage"`
		} `json:"error"`
	}

	if err := json.Unmarshal(b, &errorval); err != nil {
		return err
	}

	if errorval.Error.ErrorMessage != "" {
		return errors.New(errorval.Error.ErrorMessage)
	}

	return json.Unmarshal(b, dst)
}
