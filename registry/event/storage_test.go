package event

import (
	"context"
	"testing"
	"time"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/utils"
)

// Test procedure
// Initial storage client
// Create
// Get
// Update
// Delete
// Multiple create & List

func TestEventStorage(t *testing.T) {
	// Setup code
	es, err := NewEventStorage(codec.NewCodec("json"))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	timestamp := time.Now().Unix()
	// Test data
	var event = models.Event{
		Type:      models.IF,
		Source:    "event_trigger",
		Message:   "this is event 1",
		Timestamp: timestamp,
	}

	var event2 = models.Event{
		Type:      models.THEN,
		Source:    "event_handler",
		Message:   "this is event 2",
		Timestamp: timestamp,
	}

	// Test suite
	t.Run("Create event", func(t *testing.T) {
		err = es.Create(context.Background(), &event)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("Get event", func(t *testing.T) {
		e, err := es.Get(context.Background(), event.Type, event.Source, timestamp)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		utils.AssertEqual(t, e.Type, event.Type, "")
		utils.AssertEqual(t, e.Source, event.Source, "")
		utils.AssertEqual(t, e.Message, event.Message, "")
		utils.AssertEqual(t, e.Timestamp, event.Timestamp, "")
	})

	t.Run("Delete event", func(t *testing.T) {
		err = es.Delete(context.Background(), event.Type, event.Source, timestamp)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("List events", func(t *testing.T) {
		err = es.Create(context.Background(), &event)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		err = es.Create(context.Background(), &event2)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		el, err := es.List(context.Background())
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		utils.AssertEqual(t, len(el), 2, "")

		err = es.Delete(context.Background(), event.Type, event.Source, timestamp)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		err = es.Delete(context.Background(), event2.Type, event2.Source, timestamp)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	// Teardown code
	es.Close()
}
