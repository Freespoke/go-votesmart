package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) StateGetStateIDs(ctx context.Context) (*votesmarttypes.StateGetStateIDs, error) {
	var resp votesmarttypes.StateGetStateIDs
	if err := c.get(ctx, "State.getStateIDs", &url.Values{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) StateGetState(ctx context.Context, stateID string) (*votesmarttypes.StateGetState, error) {
	v := url.Values{}
	v.Add("stateId", stateID)
	var resp votesmarttypes.StateGetState
	if err := c.get(ctx, "State.getState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
