package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) LeadershipGetPositions(ctx context.Context, stateID, officeID string) (*votesmarttypes.LeadershipGetPositions, error) {
	v := url.Values{}
	v.Add("stateId", stateID)
	v.Add("officeId", officeID)
	var resp votesmarttypes.LeadershipGetPositions
	if err := c.get(ctx, "Leadership.getPositions", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) LeadershipGetOfficials(ctx context.Context, leadershipID, stateID string) (*votesmarttypes.LeadershipGetOfficials, error) {
	v := url.Values{}
	v.Add("leadershipId", leadershipID)
	v.Add("stateId", stateID)
	var resp votesmarttypes.LeadershipGetOfficials
	if err := c.get(ctx, "Leadership.getOfficials", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
