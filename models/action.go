package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/joshuakwan/hydra/storage"
)

const (
	actionRegistryName = "/actions/"
)

const actionStringTemplate = `Action {{.Module}}.{{.Name}}
- enabled: {{.Enabled}}
- description: {{.Description}}`

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

// CreateAction creates a db entry for an action object
func CreateAction(action Action) error {
	client, destroyFunc, err := storage.CreateClient()
	if err != nil {
		return err
	}
	defer destroyFunc()
	timeOutContext, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	data, err := json.Marshal(action)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s%s/%s", actionRegistryName, action.Module, action.Name)
	err = client.CreateObject(timeOutContext, path, string(data))
	if err != nil {
		return err
	}
	return nil
}

// GetAction retrieves a db entry for an action object
func GetAction(module, name string) (*Action, error) {
	client, destroyFunc, err := storage.CreateClient()
	if err != nil {
		return nil, err
	}
	defer destroyFunc()
	timeOutContext, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	path := fmt.Sprintf("%s%s/%s", actionRegistryName, module, name)
	data, err := client.GetObject(timeOutContext, path)
	if err != nil {
		return nil, err
	}

	var a Action
	err = json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}
	return &a, err
}

// DeleteAction deletes a db entry for an action object
func DeleteAction(module, name string) error {
	client, destroyFunc, err := storage.CreateClient()
	if err != nil {
		return err
	}
	defer destroyFunc()
	timeOutContext, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	path := fmt.Sprintf("%s%s/%s", actionRegistryName, module, name)
	count, err := client.DeleteObject(timeOutContext, path)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("key %s not exist, nothing deleted", path)
	}
	return nil
}

// UpdateAction updates a db entry for an action object
func UpdateAction(action *Action) error {
	client, destroyFunc, err := storage.CreateClient()
	if err != nil {
		return err
	}
	defer destroyFunc()
	timeOutContext, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	path := fmt.Sprintf("%s%s/%s", actionRegistryName, action.Module, action.Name)
	data, err := json.Marshal(action)
	if err != nil {
		return err
	}
	return client.UpdateObject(timeOutContext, path, string(data))
}
