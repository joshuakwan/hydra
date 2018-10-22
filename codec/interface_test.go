package codec

import (
	"testing"

	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/utils"
)

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

var actionJSON = `
{
	"module": "action_module",
	"name": "action_name",
	"description": "action description",
	"enabled": true,
	"parameters": [
		{"name": "parameter_1", "type":"string", "description": "description 1"},
		{"name": "parameter_2", "type":"int", "description": "description 2"}
	]
}
`

func TestActionCodec(t *testing.T) {
	var a models.Action
	codec := NewCodec("json")
	err := codec.Decode([]byte(actionJSON), &a)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	utils.AssertEqual(t, a.Module, "action_module", "")
	utils.AssertEqual(t, a.Name, "action_name", "")
	utils.AssertEqual(t, a.Description, "action description", "")
	utils.AssertEqual(t, a.Enabled, true, "")
	for _, para := range a.Parameters {
		if para.Name == "parameter_1" {
			utils.AssertEqual(t, para.Type, "string", "")
			utils.AssertEqual(t, para.Description, "description 1", "")
		}
		if para.Name == "parameter_2" {
			utils.AssertEqual(t, para.Type, "int", "")
			utils.AssertEqual(t, para.Description, "description 2", "")
		}
		if para.Name != "parameter_1" && para.Name != "parameter_2" {
			t.Fail()
		}
	}

	data, err := codec.Encode(&action)
	var newAction models.Action
	err = codec.Decode(data, &newAction)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	utils.AssertEqual(t, newAction.Module, "action_module", "")
	utils.AssertEqual(t, newAction.Name, "action_name", "")
	utils.AssertEqual(t, newAction.Description, "action description", "")
	utils.AssertEqual(t, newAction.Enabled, true, "")
	for _, para := range newAction.Parameters {
		if para.Name == "parameter_1" {
			utils.AssertEqual(t, para.Type, "string", "")
			utils.AssertEqual(t, para.Description, "description 1", "")
		}
		if para.Name == "parameter_2" {
			utils.AssertEqual(t, para.Type, "int", "")
			utils.AssertEqual(t, para.Description, "description 2", "")
		}
		if para.Name != "parameter_1" && para.Name != "parameter_2" {
			t.Fail()
		}
	}
}
