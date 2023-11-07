package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) ElectionGetElection(ctx context.Context, electionID string) (*votesmarttypes.ElectionGetElection, error) {
	v := url.Values{}
	v.Add("electionId", electionID)
	var resp votesmarttypes.ElectionGetElection
	if err := c.get(ctx, "Election.getElection", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ElectionGetElectionByYearState(ctx context.Context, year, stateID string) (*votesmarttypes.ElectionGetElectionByYearState, error) {
	v := url.Values{}
	v.Add("year", year)
	v.Add("stateId", stateID)
	var resp votesmarttypes.ElectionGetElectionByYearState
	if err := c.get(ctx, "Election.getElectionByYearState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ElectionGetElectionByZip(ctx context.Context, zip5, zip4, year string) (*votesmarttypes.ElectionGetElectionByZip, error) {
	v := url.Values{}
	v.Add("year", year)
	v.Add("zip5", zip5)
	v.Add("zip4", zip4)
	var resp votesmarttypes.ElectionGetElectionByZip
	if err := c.get(ctx, "Election.getElectionByZip", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ElectionGetStageCandidates(ctx context.Context, electionID, stageID, major, party, districtID, stateID string) (any, error) {
	return nil, ErrNotImplemented
}
