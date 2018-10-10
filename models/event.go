package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/joshuakwan/hydra/storage"
)

const (
	eventRegistryName = "/events/"
)

// EventType represents the type of an event
type EventType int

const (
	// IF represents a triggered event
	IF EventType = iota
	// THEN represents a generated event by a rule
	THEN
	// FINALLY represents a generated event by an action
	FINALLY
)

func (e EventType) String() string {
	switch e {
	case IF:
		return "IF"
	case THEN:
		return "THEN"
	case FINALLY:
		return "FINALLY"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

// Event represents an event
type Event struct {
	Type      EventType `json:"type"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
	Timestamp int64     `json:"timestamp"`
}

// CreateIFEvent creates an IF event from a trigger
func CreateIFEvent(source, message string, timestamp int64) (*Event, error) {
	return createEvent(IF, source, message, timestamp)
}

// CreateTHENEvent creates an THEN event from a trigger
func CreateTHENEvent(source, message string, timestamp int64) (*Event, error) {
	return createEvent(THEN, source, message, timestamp)
}

// CreateFINALLYEvent creates an FINALLY event from a trigger
func CreateFINALLYEvent(source, message string, timestamp int64) (*Event, error) {
	return createEvent(FINALLY, source, message, timestamp)
}

func createEvent(evtType EventType, source, message string, timestamp int64) (*Event, error) {
	event := Event{Type: evtType, Source: source, Message: message, Timestamp: timestamp}

	client, destroyFunc, err := storage.CreateClient()
	if err != nil {
		return nil, err
	}
	defer destroyFunc()
	timeOutContext, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = client.CreateObject(
		timeOutContext,
		fmt.Sprintf("%s%s/%s/%s", eventRegistryName, event.Type, event.Source, strconv.FormatInt(event.Timestamp, 10)),
		string(data))
	if err != nil {
		return nil, err
	}

	return &event, nil
}
