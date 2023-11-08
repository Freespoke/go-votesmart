package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) CandidatesGetByOfficeState(ctx context.Context, officeID, stateID, electionYear, stageID string) (*votesmarttypes.CandidatesGetByOfficeState, error) {
	v := url.Values{}
	v.Add("officeId", officeID)
	v.Add("stateId", stateID)
	v.Add("electionYear", electionYear)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByOfficeState
	if err := c.get(ctx, "Candidates.getByOfficeState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidatesGetByOfficeTypeState(ctx context.Context, officeTypeID, stateID, electionYear, stageID string) (*votesmarttypes.CandidatesGetByOfficeTypeState, error) {
	v := url.Values{}
	v.Add("officeTypeId", officeTypeID)
	v.Add("stateId", stateID)
	v.Add("electionYear", electionYear)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByOfficeTypeState
	if err := c.get(ctx, "Candidates.getByOfficeTypeState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidatesGetByLastname(ctx context.Context, lastName, electionYear, stageID string) (*votesmarttypes.CandidatesGetByLastname, error) {
	v := url.Values{}
	v.Add("lastName", lastName)
	v.Add("electionYear", electionYear)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByLastname
	if err := c.get(ctx, "Candidates.getByLastname", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidatesGetByLevenshtein(ctx context.Context, lastName, electionYear, stageID string) (*votesmarttypes.CandidatesGetByLastname, error) {
	v := url.Values{}
	v.Add("lastName", lastName)
	v.Add("electionYear", electionYear)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByLastname
	if err := c.get(ctx, "Candidates.getByLevenshtein", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidatesGetByElection(ctx context.Context, electionID, stageID string) (*votesmarttypes.CandidatesGetByElection, error) {
	v := url.Values{}
	v.Add("electionID", electionID)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByElection
	if err := c.get(ctx, "Candidates.getByElection", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidatesGetByDistrict(ctx context.Context, districtID, electionYear, stageID string) (*votesmarttypes.CandidatesGetByDistrict, error) {
	v := url.Values{}
	v.Add("districtId", districtID)
	v.Add("electionYear", electionYear)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByDistrict
	if err := c.get(ctx, "Candidates.getByDistrict", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CandidatesGetByZip(ctx context.Context, zip5, zip4, electionYear, stageID string) (*votesmarttypes.CandidatesGetByZip, error) {
	v := url.Values{}
	v.Add("zip5", zip5)
	v.Add("zip4", zip4)
	v.Add("electionYear", electionYear)
	v.Add("stageId", stageID)
	var resp votesmarttypes.CandidatesGetByZip
	if err := c.get(ctx, "Candidates.getByZip", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
