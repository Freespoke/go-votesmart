package internal

import "context"

func (c *Client) LeadershipGetPositions(ctx context.Context, stateID, officeID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) LeadershipGetOfficials(ctx context.Context, leadershipID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}
