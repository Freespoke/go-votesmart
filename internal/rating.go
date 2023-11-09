package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) RatingGetCategories(ctx context.Context, stateID string) (*votesmarttypes.RatingGetCategories, error) {
	v := url.Values{}
	v.Add("stateId", stateID)
	var resp votesmarttypes.RatingGetCategories
	if err := c.get(ctx, "Rating.getCategories", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) RatingGetSigList(ctx context.Context, categoryID, stateID string) (*votesmarttypes.RatingGetSigList, error) {
	v := url.Values{}
	v.Add("categoryId", categoryID)
	v.Add("stateId", stateID)
	var resp votesmarttypes.RatingGetSigList
	if err := c.get(ctx, "Rating.getSigList", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) RatingGetSig(ctx context.Context, sigID string) (*votesmarttypes.RatingGetSig, error) {
	v := url.Values{}
	v.Add("sigId", sigID)
	var resp votesmarttypes.RatingGetSig
	if err := c.get(ctx, "Rating.getSig", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) RatingGetSigRatings(ctx context.Context, sigID string) (*votesmarttypes.RatingGetSigRatings, error) {
	v := url.Values{}
	v.Add("sigId", sigID)
	var resp votesmarttypes.RatingGetSigRatings
	if err := c.get(ctx, "Rating.getSigRatings", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) RatingGetCandidateRating(ctx context.Context, candidateID, sigID string) (*votesmarttypes.RatingGetCandidateRating, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	v.Add("sigId", sigID)
	var resp votesmarttypes.RatingGetCandidateRating
	if err := c.get(ctx, "Rating.getCandidateRating", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) RatingGetRating(ctx context.Context, ratingID string) (*votesmarttypes.RatingGetRating, error) {
	v := url.Values{}
	v.Add("ratingId", ratingID)
	var resp votesmarttypes.RatingGetRating
	if err := c.get(ctx, "Rating.getRating", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
