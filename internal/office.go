package internal

import "context"

func (c *Client) OfficeGetTypes(ctx context.Context) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficeGetBranches(ctx context.Context) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficeGetLevels(ctx context.Context) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficeGetOfficesByType(ctx context.Context, officeTypeID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficeGetOfficesByLevel(ctx context.Context, levelID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficeGetOfficesByTypeLevel(ctx context.Context, officeTypeID, officeLevelID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficeGetOfficesByBranchLevel(ctx context.Context, branchID, levelID string) (any, error) {
	return nil, ErrNotImplemented
}
