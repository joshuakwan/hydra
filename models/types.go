package models

import (
	"bytes"
	"fmt"
	"html/template"
)

// ObjectType represents the type of an object
type ObjectType string

// Object is the base object
type Object interface {
	ObjectType() ObjectType
}

// Parameter defines the parameter
type Parameter struct {
	Name        string `json:"name" yaml:"name"`
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

// Action defines the data model of an action
type Action struct {
	Module      string      `json:"module" yaml:"module"`
	Name        string      `json:"name" yaml:"name"`
	Description string      `json:"description" yaml:"description"`
	Enabled     bool        `json:"enabled" yaml:"enabled"`
	Parameters  []Parameter `json:"parameters" yaml:"parameters"`
}

// ActionList defines a list of Actions
type ActionList struct {
	Actions []Action `json:"actions" yaml:"actions"`
}

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

// ObjectType returns the type of Parameter
func (p Parameter) ObjectType() ObjectType {
	return "Parameter"
}

// ObjectType returns the type of Action
func (a Action) ObjectType() ObjectType {
	return "Action"
}

const actionStringTemplate = `Action {{.Module}}.{{.Name}}
- enabled: {{.Enabled}}
- description: {{.Description}}`

func (a Action) String() string {
	var buffer bytes.Buffer
	t := template.New("action template")
	t, err := t.Parse(actionStringTemplate)
	if err != nil {
		return ""
	}

	err = t.Execute(&buffer, a)
	if err != nil {
		return ""
	}
	return buffer.String()
}

func (al ActionList) ObjectType() ObjectType {
	return "ActionList"
}

// ObjectType returns the type of Event
func (e Event) ObjectType() ObjectType {
	return "Event"
}
