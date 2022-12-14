package matata

import (
	"context"
	"fmt"
	"net/http"
)

type AbsenceType struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Archived       bool   `json:"archived"`
	GrantsWorkTime bool   `json:"grants_work_time"`
	IsVacation     bool   `json:"is_vacation"`
}

// AbsenceTypeList represents a list of AbsenceType objects.
type AbsenceTypeList []AbsenceType

func (c *Client) GetAbsenceTypes(ctx context.Context) (*AbsenceTypeList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/absence_types", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := AbsenceTypeList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
