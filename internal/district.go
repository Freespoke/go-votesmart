package internal

import "context"

func (c *Client) DistrictGetByOfficeState(ctx context.Context, officeID, stateID, districtName string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) DistrictGetByZip(ctx context.Context, zip5, zip4 string) (any, error) {
	return nil, ErrNotImplemented
}
