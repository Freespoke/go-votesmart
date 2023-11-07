package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) OfficeGetTypes(ctx context.Context) (*votesmarttypes.OfficeGetTypes, error) {
	var resp votesmarttypes.OfficeGetTypes
	if err := c.get(ctx, "Office.getTypes", &url.Values{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficeGetBranches(ctx context.Context) (*votesmarttypes.OfficeGetBranches, error) {
	var resp votesmarttypes.OfficeGetBranches
	if err := c.get(ctx, "Office.getBranches", &url.Values{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficeGetLevels(ctx context.Context) (*votesmarttypes.OfficeGetLevels, error) {
	var resp votesmarttypes.OfficeGetLevels
	if err := c.get(ctx, "Office.getLevels", &url.Values{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficeGetOfficesByType(ctx context.Context, officeTypeID string) (*votesmarttypes.OfficeGetOfficesByType, error) {
	v := url.Values{}
	v.Add("officeTypeId", officeTypeID)
	var resp votesmarttypes.OfficeGetOfficesByType
	if err := c.get(ctx, "Office.getOfficesByType", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficeGetOfficesByLevel(ctx context.Context, levelID string) (*votesmarttypes.OfficeGetOfficesByLevel, error) {
	v := url.Values{}
	v.Add("levelId", levelID)
	var resp votesmarttypes.OfficeGetOfficesByLevel
	if err := c.get(ctx, "Office.getOfficesByLevel", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficeGetOfficesByTypeLevel(ctx context.Context, officeTypeID, officeLevelID string) (*votesmarttypes.OfficeGetOfficesByTypeLevel, error) {
	v := url.Values{}
	v.Add("officeTypeId", officeTypeID)
	v.Add("officeLevelId", officeLevelID)
	var resp votesmarttypes.OfficeGetOfficesByTypeLevel
	if err := c.get(ctx, "Office.getOfficesByTypeLevel", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) OfficeGetOfficesByBranchLevel(ctx context.Context, branchID, levelID string) (*votesmarttypes.OfficeGetOfficesByBranchLevel, error) {
	v := url.Values{}
	v.Add("branchId", branchID)
	v.Add("levelId", levelID)
	var resp votesmarttypes.OfficeGetOfficesByBranchLevel
	if err := c.get(ctx, "Office.getOfficesByBranchLevel", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
