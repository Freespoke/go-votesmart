package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) CandidateBioGetBio(ctx context.Context, candidateID string) (*votesmarttypes.CandidateBioGetBio, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	var resp votesmarttypes.CandidateBioGetBio
	if err := c.get(ctx, "CandidateBio.getBio", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidateBioGetDetailedBio(ctx context.Context, candidateID string) (*votesmarttypes.CandidateBioGetDetailedBio, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	var resp votesmarttypes.CandidateBioGetDetailedBio
	if err := c.get(ctx, "CandidateBio.getDetailedBio", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidateBioGetAdditionalBio(ctx context.Context, candidateID string) (*votesmarttypes.CandidateBioGetAdditionalBio, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	var resp votesmarttypes.CandidateBioGetAdditionalBio
	if err := c.get(ctx, "CandidateBio.getAddlBio", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
