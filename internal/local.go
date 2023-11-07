package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) LocalGetCounties(ctx context.Context, stateID string) (*votesmarttypes.LocalGetCounties, error) {
	v := url.Values{}
	v.Add("stateId", stateID)
	var resp votesmarttypes.LocalGetCounties
	if err := c.get(ctx, "Local.getCounties", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) LocalGetCities(ctx context.Context, stateID string) (*votesmarttypes.LocalGetCities, error) {
	v := url.Values{}
	v.Add("stateId", stateID)
	var resp votesmarttypes.LocalGetCities
	if err := c.get(ctx, "Local.getCities", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) LocalGetOfficials(ctx context.Context, localID string) (*votesmarttypes.LocalGetOfficials, error) {
	v := url.Values{}
	v.Add("localId", localID)
	var resp votesmarttypes.LocalGetOfficials
	if err := c.get(ctx, "Local.getOfficials", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
