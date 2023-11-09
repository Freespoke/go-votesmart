package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) AddressGetCampaign(ctx context.Context, candidateID string) (*votesmarttypes.AddressGetCampaign, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	var resp votesmarttypes.AddressGetCampaign
	if err := c.get(ctx, "Address.getCampaign", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) AddressGetCampaignWebAddress(ctx context.Context, candidateID string) (*votesmarttypes.AddressGetCampaignWebAddress, error) {
	v := url.Values{}
	v.Add("candidateId", candidateID)
	var resp votesmarttypes.AddressGetCampaignWebAddress
	if err := c.get(ctx, "Address.getCampaignWebAddress", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) AddressGetCampaignByElection(ctx context.Context, electionID string) (*votesmarttypes.AddressGetCampaignByElection, error) {
	v := url.Values{}
	v.Add("electionId", electionID)
	var resp votesmarttypes.AddressGetCampaignByElection
	if err := c.get(ctx, "Address.getCampaignByElection", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) AddressGetOffice(ctx context.Context, candidateID string) (any, error) {
	// Address.getOffice
	return nil, ErrNotImplemented
}

func (c *Client) AddressGetOfficeWebAddress(ctx context.Context, candidateID string) (any, error) {
	// Address.getOfficeWebAddress
	return nil, ErrNotImplemented
}

func (c *Client) AddressGetOfficeByOfficeState(ctx context.Context, officeID string, stateID string) (any, error) {
	// Address.getOfficeByOfficeState
	return nil, ErrNotImplemented
}
