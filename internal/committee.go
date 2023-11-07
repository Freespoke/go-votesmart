package internal

import "context"

func (c *Client) CommitteeGetTypes(ctx context.Context) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CommitteeGetCommitteesByTypeState(ctx context.Context, typeID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CommitteeGetCommittee(ctx context.Context, committeeID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CommitteeGetCommitteeMembers(ctx context.Context, committeeID string) (any, error) {
	return nil, ErrNotImplemented
}
