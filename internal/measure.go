package internal

import (
	"context"
	"net/url"

	"dev.freespoke.com/go-votesmart/votesmarttypes"
)

func (c *Client) MeasureGetMeasuresByYearState(ctx context.Context, year string, stateID string) (*votesmarttypes.MeasureGetMeasuresByYearState, error) {
	v := url.Values{}
	v.Add("year", year)
	v.Add("stateId", stateID)
	var resp votesmarttypes.MeasureGetMeasuresByYearState
	if err := c.get(ctx, "Measure.getMeasuresByYearState", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) MeasureGetMeasure(ctx context.Context, measureID string) (*votesmarttypes.MeasureGetMeasure, error) {
	v := url.Values{}
	v.Add("measureId", measureID)
	var resp votesmarttypes.MeasureGetMeasure
	if err := c.get(ctx, "Measure.getMeasure", &v, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
