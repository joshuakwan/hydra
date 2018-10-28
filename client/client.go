package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/joshuakwan/hydra/models"
)

// Client is used to talk to the API server
type Client struct {
	ServerURL string
}

// NewClient initializes a new REST client to the API server
func NewClient(serverURL string) Client {
	return Client{ServerURL: serverURL}
}

func (c *Client) ListEvents() ([]*models.Event, error) {
	resp, err := resty.R().Get(fmt.Sprintf("%s%s", c.ServerURL, "/events"))
	if err != nil {
		return nil, err
	}
	var events []*models.Event
	if err = json.Unmarshal(resp.Body(), &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (c *Client) CreateEvent(event *models.Event) error {
	_, err := resty.R().SetBody(event).Post(fmt.Sprintf("%s%s", c.ServerURL, "/events"))
	return err
}
