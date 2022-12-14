package matata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StartTimerOptions struct {
	TaskID    int    `json:"task_id"`
	StartTime string `json:"start_time,omitempty"`
	ProjectID int    `json:"project_id,omitempty"`
	Note      string `json:"note,omitempty"`
}

type StopTimerOptions struct {
	EndTime string `json:"end_time,omitempty"`
}

type Timer struct {
	Date              string  `json:"date"`
	StartTime         string  `json:"start_time"`
	Duration          string  `json:"duration"`
	DurationInSeconds float64 `json:"duration_in_seconds"`
	Note              string  `json:"note"`
	User              User    `json:"user"`
	Task              Task    `json:"task"`
	Project           Project `json:"project"`
}

func (c *Client) GetTimer(ctx context.Context) (*Timer, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/timer", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := Timer{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) StartTimer(ctx context.Context, options *StartTimerOptions) (*Timer, error) {

	if options == nil {
		log.Fatal("At least a task id is required to start a timer.")
	}

	body, _ := json.Marshal(options)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/timer", c.BaseURL), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	res := Timer{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) StopTimer(ctx context.Context, options *StopTimerOptions) (*Timer, error) {

	body, _ := json.Marshal(options)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/timer", c.BaseURL), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	res := Timer{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) CancelTimer(ctx context.Context) (*Timer, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/timer", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	if err := c.sendRequest(req, nil); err != nil {
		return nil, err
	}

	return nil, nil

}
