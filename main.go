package matata

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	baseURL = "hakuna.ch/api/v1"
)

type Client struct {
	BaseURL string
	Tennant string
	token   string
	Client  *http.Client
}

// TODO: Is a bit weird, since somethines you get a Full error object with status code and message, and other times you just get a json object with at error attribute containing a string
type errorResponse struct {
	HTTPStatus int
	Status     int    `json:"status"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func NewClient(tennant string, token string) *Client {

	if tennant == "" {
		log.Fatal("a tennant is required")
	}

	if token == "" {
		log.Fatal("a token is required")

	}

	return &Client{
		BaseURL: "https://" + tennant + "." + baseURL,
		Tennant: tennant,
		token:   token,
		Client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Token", c.token)

	res, err := c.Client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {

		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {

			if errRes.Message == "" {
				return fmt.Errorf(fmt.Sprintf("%d - %s", res.StatusCode, http.StatusText(res.StatusCode)))
			}

			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil

}
