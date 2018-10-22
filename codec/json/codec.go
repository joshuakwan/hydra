package json

import (
	"encoding/json"
	"fmt"

	"github.com/joshuakwan/hydra/models"
)

type JSONCodec struct {
}

func NewJSONCodec() *JSONCodec {
	return &JSONCodec{}
}

func (c *JSONCodec) Encode(obj models.Object) ([]byte, error) {
	switch obj.ObjectType() {
	case "Parameter":
		return json.Marshal(obj.(*models.Parameter))
	case "Action":
		return json.Marshal(obj.(*models.Action))
	case "ActionList":
		return json.Marshal(obj.(*models.ActionList))
	case "Event":
		return json.Marshal(obj.(*models.Event))
	default:
		return nil, fmt.Errorf("invalid type")
	}
}

func (c *JSONCodec) Decode(data []byte, objRef models.Object) error {
	switch objRef.ObjectType() {
	case "Parameter":
		return json.Unmarshal(data, objRef.(*models.Parameter))
	case "Action":
		return json.Unmarshal(data, objRef.(*models.Action))
	case "ActionList":
		return json.Unmarshal(data, objRef.(*models.ActionList))
	case "Event":
		return json.Unmarshal(data, objRef.(*models.Event))
	default:
		return fmt.Errorf("invalid type")
	}
}
