package internal

import "context"

func (c *Client) RatingGetCategories(ctx context.Context, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) RatingGetSigList(ctx context.Context, categoryID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) RatingGetSig(ctx context.Context, sigID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) RatingGetSigRatings(ctx context.Context, sigID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) RatingGetCandidateRating(ctx context.Context, candidateID, sigID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) RatingGetRating(ctx context.Context, ratingID string) (any, error) {
	return nil, ErrNotImplemented
}
