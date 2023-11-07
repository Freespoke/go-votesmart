package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) DistrictGetByOfficeState(ctx context.Context, officeID, stateID, districtName string) (*votesmarttypes.DistrictGetByOfficeState, error) {
	v := url.Values{}
	v.Add("officeId", officeID)
	v.Add("stateId", stateID)
	v.Add("districtName", districtName)
	var resp votesmarttypes.DistrictGetByOfficeState
	if err := c.get(ctx, "District.getByOfficeState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) DistrictGetByZip(ctx context.Context, zip5, zip4 string) (*votesmarttypes.DistrictGetByZip, error) {
	v := url.Values{}
	v.Add("zip5", zip5)
	v.Add("zip4", zip4)
	var resp votesmarttypes.DistrictGetByZip
	if err := c.get(ctx, "District.getByZip", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
