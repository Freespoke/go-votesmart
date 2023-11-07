package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) CommitteeGetTypes(ctx context.Context) (*votesmarttypes.CommitteeGetTypes, error) {
	var resp votesmarttypes.CommitteeGetTypes
	if err := c.get(ctx, "Committee.getTypes", &url.Values{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CommitteeGetCommitteesByTypeState(ctx context.Context, typeID, stateID string) (*votesmarttypes.CommitteeGetCommitteesByTypeState, error) {
	v := url.Values{}
	v.Add("typeId", typeID)
	v.Add("stateId", stateID)
	var resp votesmarttypes.CommitteeGetCommitteesByTypeState
	if err := c.get(ctx, "Committee.getCommitteesByTypeState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CommitteeGetCommittee(ctx context.Context, committeeID string) (*votesmarttypes.CommitteeGetCommittee, error) {
	v := url.Values{}
	v.Add("committeeId", committeeID)
	var resp votesmarttypes.CommitteeGetCommittee
	if err := c.get(ctx, "Committee.getCommittee", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) CommitteeGetCommitteeMembers(ctx context.Context, committeeID string) (*votesmarttypes.CommitteeGetCommitteeMembers, error) {
	v := url.Values{}
	v.Add("committeeId", committeeID)
	var resp votesmarttypes.CommitteeGetCommitteeMembers
	if err := c.get(ctx, "Committee.getCommitteeMembers", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
