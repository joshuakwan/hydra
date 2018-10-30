package models

import (
	"bytes"
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
type EventType string

const (
	// IF represents a triggered event
	IF EventType = "IF"
	// THEN represents a generated event by a rule
	THEN EventType = "THEN"
	// FINALLY represents a generated event by an action
	FINALLY EventType = "FINALLY"
)

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

// ObjectType returns the type of ActionList
func (al ActionList) ObjectType() ObjectType {
	return "ActionList"
}

// ObjectType returns the type of Event
func (e Event) ObjectType() ObjectType {
	return "Event"
}

// Now comes to the core RULE part

// Then defines a successful evaluation action
type Then struct {
	Run        string            `json:"run" yaml:"run"`
	Parameters map[string]string `json:"parameters" yaml:"parameters"`
}

// Rule defines a rule
type Rule struct {
	Module      string `json:"module" yaml:"module"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	If          string `json:"if" yaml:"if"`
	Then        *Then  `json:"then" yaml:"then"`
}

// ObjectType returns the type of Then
func (t Then) ObjectType() ObjectType {
	return "Then"
}

// ObjectType returns the type of Rule
func (r Rule) ObjectType() ObjectType {
	return "Rule"
}
