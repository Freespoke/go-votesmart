package internal

import "context"

func (c *Client) LocalGetCounties(ctx context.Context, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) LocalGetCities(ctx context.Context, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) LocalGetOfficials(ctx context.Context, localID string) (any, error) {
	return nil, ErrNotImplemented
}
