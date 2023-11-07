package internal

import "context"

func (c *Client) CandidatesGetByOfficeState(ctx context.Context, officeID, stateID, electionYear, stageID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CandidatesGetByOfficeTypeState(ctx context.Context, officeTypeID, stateID, electionYear, stageID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CandidatesGetByLastname(ctx context.Context, lastName, electionYear, stageID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CandidatesGetByLevenshtein(ctx context.Context, lastName, electionyear, stageID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CandidatesGetByElection(ctx context.Context, electionID, stageID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CandidatesGetByDistrict(ctx context.Context, districtID, electionYear, stageID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) CandidatesGetByZip(ctx context.Context, zip5, zip4, electionYear, stageID string) (any, error) {
	return nil, ErrNotImplemented
}
