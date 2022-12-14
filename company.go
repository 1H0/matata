package matata

import (
	"context"
	"fmt"
	"net/http"
)

type Company struct {
	CompanyName            string `json:"company_name"`
	DurationFormat         string `json:"duration_format"`
	AbsenceRequestsEnabled bool   `json:"absence_requests_enabled"`
	ProjectsEnabled        bool   `json:"projects_enabled"`
	GroupsEnabled          bool   `json:"groups_enabled"`
}

func (c *Client) GetCompany(ctx context.Context) (*Company, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/company", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := Company{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
