package matata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeEntriesOptions struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// TODO: Omit emtys are not ideal here
// TODO: Maybe rename to general name, since it's used for creation and update... maybe use general time entry?

type TimeEntryOptions struct {
	Date      string `json:"date,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	TaskID    int    `json:"task_id,omitempty"`
	ProjectID int    `json:"project_id,omitempty"`
	Note      string `json:"note,omitempty"`
}

type TimeEntry struct {
	ID                int     `json:"id"`
	Date              string  `json:"date"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	Duration          string  `json:"duration"`
	DurationInSeconds float64 `json:"duration_in_seconds"`
	Note              string  `json:"note"`
	User              User    `json:"user"`
	Task              Task    `json:"task"`
	Project           Project `json:"project"`
}

type TimeEntryList []TimeEntry

func (c *Client) GetTimeEntries(ctx context.Context, options *TimeEntriesOptions) (*TimeEntryList, error) {

	if options.StartDate == "" {
		options.StartDate = time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	}

	if options.EndDate == "" {
		options.EndDate = time.Now().Format("2006-01-02")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/time_entries?start_date=%s&end_date=%s", c.BaseURL, options.StartDate, options.EndDate), nil)

	if err != nil {
		return nil, err
	}

	res := TimeEntryList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) GetTimeEntry(ctx context.Context, id int) (*TimeEntry, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/time_entries/%d", c.BaseURL, id), nil)

	if err != nil {
		return nil, err
	}

	res := TimeEntry{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) CreateTimeEntry(ctx context.Context, options *TimeEntryOptions) (*TimeEntry, error) {

	body, _ := json.Marshal(options)

	fmt.Println(string(body))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/time_entries", c.BaseURL), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	res := TimeEntry{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateTimeEntry(ctx context.Context, id int, options *TimeEntryOptions) (*TimeEntry, error) {

	body, _ := json.Marshal(options)

	fmt.Println(string(body))

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/time_entries/%d", c.BaseURL, id), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	res := TimeEntry{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteTimeEntry(ctx context.Context, id int) (*TimeEntry, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/time_entries/%d", c.BaseURL, id), nil)

	if err != nil {
		return nil, err
	}

	res := TimeEntry{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
