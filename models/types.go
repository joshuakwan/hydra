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

// Event represents an event
type Event struct {
	Source    string `json:"source"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

const eventStringTemplate = `Event from {{.Source}} at {{.Timestamp}}: {{.Message}} `

func (e Event) String() string {
	var buffer bytes.Buffer
	t := template.New("event template")
	t, err := t.Parse(eventStringTemplate)
	if err != nil {
		return ""
	}

	err = t.Execute(&buffer, e)
	if err != nil {
		return ""
	}
	return buffer.String()
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

// WorkerStatus defines status of workers
type WorkerStatus string

const (
	// UP means the worker is running
	UP WorkerStatus = "UP"
	// DOWN means the worker is down
	DOWN WorkerStatus = "DOWN"
	// ERROR means the worker is in an error state
	ERROR WorkerStatus = "ERROR"
)

// Worker defines a worker
type Worker struct {
	Name       string       `json:"name" yaml:"name"`
	Address    string       `json:"address" yaml:"address"`
	Status     WorkerStatus `json:"status" yaml:"status"`
	Registered int64        `json:"registeredTimestamp"`
	LastReport int64        `json:"lastReportTimeStamp"`
}

// ObjectType returns the type of Worker
func (w Worker) ObjectType() ObjectType {
	return "Worker"
}
