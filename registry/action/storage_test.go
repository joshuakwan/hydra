package action

import (
	"context"
	"testing"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/storage"
	"github.com/joshuakwan/hydra/utils"
)

func getAction(as *Storage, module, name string) (*models.Action, error) {
	return as.Get(context.Background(), module, name)
}

func TestGetAction(t *testing.T) {
	storage, destroy, err := storage.NewStorage()
	codec := codec.NewCodec("json")
	as := NewActionStorage(storage, codec, destroy)
	defer as.Close()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	a, err := getAction(as, "core", "remote_command")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	utils.AssertEqual(t, a.Module, "core", "")
	utils.AssertEqual(t, a.Name, "remote_command", "")
	utils.AssertEqual(t, a.Enabled, true, "")
}

func TestListActions(t *testing.T) {
	storage, destroy, err := storage.NewStorage()
	codec := codec.NewCodec("json")
	as := NewActionStorage(storage, codec, destroy)
	defer as.Close()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	al, err := as.List(context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(al)
}

func TestCreateAction(t *testing.T) {
	storage, destroy, err := storage.NewStorage()
	codec := codec.NewCodec("json")
	as := NewActionStorage(storage, codec, destroy)
	defer as.Close()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	var action = models.Action{
		Module:      "action_module",
		Name:        "action_name",
		Description: "action description",
		Enabled:     true,
		Parameters: []models.Parameter{
			models.Parameter{Name: "parameter_1", Type: "string", Description: "description 1"},
			models.Parameter{Name: "parameter_2", Type: "int", Description: "description 2"},
		},
	}

	err = as.Create(context.Background(), &action)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	a, err := as.Get(context.Background(), action.Module, action.Name)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(a)

	utils.AssertEqual(t, a.Module, "action_module", "")
	utils.AssertEqual(t, a.Name, "action_name", "")
	utils.AssertEqual(t, a.Enabled, true, "")

	err = as.Delete(context.Background(), action.Module, action.Name)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
