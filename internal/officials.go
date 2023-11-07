package internal

import "context"

func (c *Client) OfficialsGetStatewide(ctx context.Context, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByOfficeState(ctx context.Context, officeID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByOfficeTypeState(ctx context.Context, officeTypeID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByLastname(ctx context.Context, lastName string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByLevenshtein(ctx context.Context, lastName string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByDistrict(ctx context.Context, districtID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByZip(ctx context.Context, zip5, zip4 string) (any, error) {
	return nil, ErrNotImplemented
}
