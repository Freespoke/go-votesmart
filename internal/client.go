package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	HTTP    *http.Client
	BaseURL *url.URL
	APIKey  string
}

var ErrNotImplemented = errors.New("not implemented")

func (c *Client) get(ctx context.Context, path string, values *url.Values, dst any) error {
	values.Add("key", c.APIKey)
	values.Add("o", "JSON")
	requestURL := *c.BaseURL
	requestURL.Path = path
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
