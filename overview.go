package matata

import (
	"context"
	"fmt"
	"net/http"
)

type Overview struct {
	Overtime          string `json:"overtime"`
	OvertimeInSeconds int    `json:"overtime_in_seconds"`
	Vacation          struct {
		RedeemedDays  float64 `json:"redeemed_days"`
		RemainingDays float64 `json:"remaining_days"`
	} `json:"vacation"`
}

func (c *Client) GetOverview(ctx context.Context) (*Overview, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/overview", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := Overview{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
