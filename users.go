package matata

import (
	"context"
	"fmt"
	"net/http"
)

type User struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Status string   `json:"status"`
	Groups []string `json:"groups"`
}

type UserList []User

func (c *Client) GetUsers(ctx context.Context) (*UserList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := UserList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/me", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	res := User{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}
