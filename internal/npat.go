package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) NpatGetNpat(ctx context.Context, candidateID string) (*votesmarttypes.NpatGetNpat, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	var resp votesmarttypes.NpatGetNpat
	if err := c.get(ctx, "Npat.getNpat", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
