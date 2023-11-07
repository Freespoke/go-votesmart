package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) OfficialsGetStatewide(ctx context.Context, stateID string) (*votesmarttypes.OfficialsGetStatewide, error) {
	v := url.Values{}
	v.Add("stateId", stateID)
	var resp votesmarttypes.OfficialsGetStatewide
	if err := c.get(ctx, "Officials.getStatewide", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficialsGetByOfficeState(ctx context.Context, officeID, stateID string) (*votesmarttypes.OfficialsGetByOfficeState, error) {
	v := url.Values{}
	v.Add("officeId", officeID)
	v.Add("stateId", stateID)
	var resp votesmarttypes.OfficialsGetByOfficeState
	if err := c.get(ctx, "Officials.getByOfficeState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficialsGetByOfficeTypeState(ctx context.Context, officeTypeID, stateID string) (*votesmarttypes.OfficialsGetByOfficeTypeState, error) {
	v := url.Values{}
	v.Add("officeTypeId", officeTypeID)
	v.Add("stateId", stateID)
	var resp votesmarttypes.OfficialsGetByOfficeTypeState
	if err := c.get(ctx, "Officials.getByOfficeTypeState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficialsGetByLastname(ctx context.Context, lastName string) (*votesmarttypes.OfficialsGetByLastname, error) {
	v := url.Values{}
	v.Add("lastName", lastName)
	var resp votesmarttypes.OfficialsGetByLastname
	if err := c.get(ctx, "Officials.getByLastname", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficialsGetByLevenshtein(ctx context.Context, lastName string) (*votesmarttypes.OfficialsGetByLevenshtein, error) {
	v := url.Values{}
	v.Add("lastName", lastName)
	var resp votesmarttypes.OfficialsGetByLevenshtein
	if err := c.get(ctx, "Officials.getByLevenshtein", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficialsGetByDistrict(ctx context.Context, districtID string) (any, error) {
	return nil, ErrNotImplemented
}

func (c *Client) OfficialsGetByZip(ctx context.Context, zip5, zip4 string) (*votesmarttypes.OfficialsGetByZip, error) {
	v := url.Values{}
	v.Add("zip5", zip5)
	v.Add("zip4", zip4)
	var resp votesmarttypes.OfficialsGetByZip
	if err := c.get(ctx, "Officials.getByZip", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
