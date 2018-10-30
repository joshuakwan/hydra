package cmd

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	action_storage "github.com/joshuakwan/hydra/registry/action"
	rule_storage "github.com/joshuakwan/hydra/registry/rule"

	"github.com/spf13/cobra"
)

func handleCreate(cmd *cobra.Command, args []string) {
	data, err := ioutil.ReadFile(filename)
	checkError(err)

	kind, err := checkYAMLKind(data)
	checkError(err)

	err = createObject(kind, data)
	checkError(err)

	fmt.Printf("%s created\n", kind)
}

func createObject(kind string, data []byte) error {
	c := codec.NewCodec("yaml")
	switch kind {
	default:
		return fmt.Errorf("invalid kind")
	case "rule":
		var rule models.Rule
		if err := c.Decode(data, &rule); err != nil {
			return err
		}
		storage, err := rule_storage.NewRuleStorage(codec.NewCodec("json"))
		if err != nil {
			return err
		}
		return storage.Create(context.Background(), &rule)
	case "action":
		var action models.Action
		if err := c.Decode(data, &action); err != nil {
			return err
		}
		storage, err := action_storage.NewActionStorage(codec.NewCodec("json"))
		if err != nil {
			return err
		}
		return storage.Create(context.Background(), &action)
	}
}
