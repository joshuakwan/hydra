package client

import (
	"testing"
	"time"

	"github.com/joshuakwan/hydra/utils"

	"github.com/joshuakwan/hydra/models"
)

func TestEvents(t *testing.T) {
	client := NewClient("http://127.0.0.1:8080")

	eventType := models.IF
	source := "test_source"
	msg := "test message"
	timestamp := time.Now().Unix()
	event := models.Event{
		Type:      eventType,
		Source:    source,
		Message:   msg,
		Timestamp: timestamp,
	}

	t.Run("Create event", func(t *testing.T) {
		err := client.CreateEvent(&event)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("List to find out the created event", func(t *testing.T) {
		events, err := client.ListEvents()
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		found := false
		for _, e := range events {
			if e.Type == eventType && e.Source == source && e.Message == msg && e.Timestamp == timestamp {
				found = true
				break
			}
		}
		utils.AssertEqual(t, found, true, "")
	})
}
