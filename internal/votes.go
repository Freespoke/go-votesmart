package internal

import "context"

func (c *Client) VotesGetCategories(ctx context.Context, year, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBill(ctx context.Context, billID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillAction(ctx context.Context, actionID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillActionVotes(ctx context.Context, actionID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillActionVoteByOfficial(ctx context.Context, actionID, candidateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetByBillNumber(ctx context.Context, billNumber string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillsByCategoryYearState(ctx context.Context, categoryID, yearID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillByYearState(ctx context.Context, year, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillsByOfficialYearOffice(ctx context.Context, candidateID, year, officeID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillsByOfficialCategoryOffice(ctx context.Context, candidateID, categoryID, officeID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetByOfficial(ctx context.Context, officeID, categoryID, year string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillsBySponsorYear(ctx context.Context, candidateID, year string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillsBySponsorCategory(ctx context.Context, candidateID, categoryID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetBillsByStateRecent(ctx context.Context, amount, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) VotesGetVetoes(ctx context.Context, candidateID string) (any, error) {
	return nil, ErrNotImplemented
}
