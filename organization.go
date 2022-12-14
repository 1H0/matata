package matata

import (
	"context"
	"fmt"
	"net/http"
)

type UserStatus struct {
	User                User `json:"user"`
	AbsentFirstHalfDay  bool `json:"absent_first_half_day"`
	AbsentSecondHalfDay bool `json:"absent_second_half_day"`
	HasTimerRunning     bool `json:"has_timer_running"`
}

type OrganizationStatusList []UserStatus

func (c *Client) GetOrganizationStatus(ctx context.Context) (*OrganizationStatusList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organization/status", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := OrganizationStatusList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
