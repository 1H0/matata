package matata

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Ping struct {
	Pong time.Time `json:"pong"`
}

func (c *Client) Ping(ctx context.Context) (*Ping, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ping", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := Ping{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
