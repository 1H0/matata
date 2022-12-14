package matata

import (
	"context"
	"fmt"
	"net/http"
)

type Project struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Archived        bool        `json:"archived"`
	Groups          interface{} `json:"groups"`
	Client          string      `json:"client"`
	Tasks           []Task      `json:"tasks"`
	Budget          string      `json:"budget"`
	BudgetInSeconds int         `json:"budget_in_seconds"`
	BudgetIsMonthly bool        `json:"budget_is_monthly"`
}

type ProjectList []Project

func (c *Client) GetProjects(ctx context.Context) (*ProjectList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := ProjectList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
