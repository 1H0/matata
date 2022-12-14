package matata

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type AbsenceOptions struct {
	Year string
}

type Absence struct {
	ID                   int         `json:"id"`
	StartDate            string      `json:"start_date"`
	EndDate              string      `json:"end_date"`
	FirstHalfDay         bool        `json:"first_half_day"`
	SecondHalfDay        bool        `json:"second_half_day"`
	IsRecurring          bool        `json:"is_recurring"`
	WeeklyRepeatInterval interface{} `json:"weekly_repeat_interval"`
	User                 User        `json:"user"`
	Type                 AbsenceType `json:"absence_type"`
}

type AbsenceList []Absence

func (c *Client) GetAbsences(ctx context.Context, options *AbsenceOptions) (*AbsenceList, error) {

	if options.Year == "" {
		options.Year = strconv.Itoa(time.Now().Year())
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/absences?year=%s", c.BaseURL, options.Year), nil)

	if err != nil {
		return nil, err
	}

	res := AbsenceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
