package action

import (
	"context"
	"testing"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/storage"
	"github.com/joshuakwan/hydra/utils"
)

// Test procedure
// Initial storage client
// Create
// Get
// Update
// Delete
// Multiple create & List

func TestActionStorage(t *testing.T) {
	// Setup code
	storage, destroy, err := storage.NewStorage()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	codec := codec.NewCodec("json")
	as := NewActionStorage(storage, codec, destroy)

	// Test data
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

	var action2 = models.Action{
		Module:      "action_module",
		Name:        "action_name_2",
		Description: "action description",
		Enabled:     true,
		Parameters: []models.Parameter{
			models.Parameter{Name: "parameter_1", Type: "string", Description: "description 1"},
			models.Parameter{Name: "parameter_2", Type: "int", Description: "description 2"},
		},
	}

	// Test suite
	t.Run("Create action", func(t *testing.T) {
		err = as.Create(context.Background(), &action)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("Get action", func(t *testing.T) {
		a, err := as.Get(context.Background(), action.Module, action.Name)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		utils.AssertEqual(t, a.Module, action.Module, "")
		utils.AssertEqual(t, a.Name, action.Name, "")
		utils.AssertEqual(t, a.Description, action.Description, "")
		utils.AssertEqual(t, a.Enabled, action.Enabled, "")
		utils.AssertEqual(t, len(a.Parameters), len(action.Parameters), "")
	})

	t.Run("Update action", func(t *testing.T) {
		a, err := as.Get(context.Background(), action.Module, action.Name)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		a.Description = "action description updated"
		a.Enabled = false

		err = as.Update(context.Background(), a)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		aUpdated, err := as.Get(context.Background(), a.Module, a.Name)

		utils.AssertEqual(t, a.Module, aUpdated.Module, "")
		utils.AssertEqual(t, a.Name, aUpdated.Name, "")
		utils.AssertEqual(t, aUpdated.Description, "action description updated", "")
		utils.AssertEqual(t, aUpdated.Enabled, false, "")
	})

	t.Run("Delete action", func(t *testing.T) {
		err = as.Delete(context.Background(), action.Module, action.Name)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("List action", func(t *testing.T) {
		err = as.Create(context.Background(), &action)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		err = as.Create(context.Background(), &action2)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		al, err := as.List(context.Background())
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		utils.AssertEqual(t, len(al), 2, "")

		err = as.Delete(context.Background(), action.Module, action.Name)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		err = as.Delete(context.Background(), action2.Module, action2.Name)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	// Teardown code
	as.Close()
}
