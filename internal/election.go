package internal

import "context"

func (c *Client) ElectionGetElection(ctx context.Context, electionID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) ElectionGetElectionByYearState(ctx context.Context, year, stateID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) ElectionGetElectionByZip(ctx context.Context, zip5, zip4, year string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) ElectionGetStageCandidates(ctx context.Context, electionID, stageID, major, party, districtID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}
