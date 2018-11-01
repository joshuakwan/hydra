package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty"
	"github.com/joshuakwan/hydra/models"
)

// Client is used to talk to the API server
type Client struct {
	ServerURL string
}

// NewClient initializes a new REST client to the API server
func NewClient(serverURL string) *Client {
	return &Client{ServerURL: serverURL}
}

// ListEvents lists all events
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

// CreateEvent creates a new Event
func (c *Client) CreateEvent(event *models.Event) error {
	_, err := resty.R().SetBody(event).Post(fmt.Sprintf("%s%s", c.ServerURL, "/events"))
	return err
}

// WatchEvents setup a HTTP channel to receive event stream
func (c *Client) WatchEvents() <-chan models.Event {
	resp, _ := http.Get(fmt.Sprintf("%s%s", c.ServerURL, "/events/watch"))

	// TODO buffer size
	eventCh := make(chan models.Event, 10)

	go func() {
		for {
			var event models.Event
			err := json.NewDecoder(resp.Body).Decode(&event)
			// TODO send an error event
			if err == nil {
				eventCh <- event
			}
		}
	}()
	return eventCh
}

// ListRules lists all rules
func (c *Client) ListRules() ([]*models.Rule, error) {
	resp, err := resty.R().Get(fmt.Sprintf("%s%s", c.ServerURL, "/rules"))
	if err != nil {
		return nil, err
	}
	var rules []*models.Rule
	if err = json.Unmarshal(resp.Body(), &rules); err != nil {
		return nil, err
	}

	return rules, nil
}
