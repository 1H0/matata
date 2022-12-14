package matata

import (
	"context"
	"fmt"
	"net/http"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	Default  bool   `json:"default"`
}

type TaskList []Task

func (c *Client) GetTasks(ctx context.Context) (*TaskList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tasks", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := TaskList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
