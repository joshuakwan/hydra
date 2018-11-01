package json

import (
	"encoding/json"
	"fmt"

	"github.com/joshuakwan/hydra/models"
)

// Codec encode/decode objects against JSON
type Codec struct {
}

// NewJSONCodec creates a new JSON codec
func NewJSONCodec() *Codec {
	return &Codec{}
}

// Encode turns object into JSON
func (c *Codec) Encode(obj models.Object) ([]byte, error) {
	switch obj.(type) {
	case *models.Parameter:
		return json.Marshal(obj.(*models.Parameter))
	case *models.Action:
		return json.Marshal(obj.(*models.Action))
	case *models.ActionList:
		return json.Marshal(obj.(*models.ActionList))
	case *models.Event:
		return json.Marshal(obj.(*models.Event))
	case *models.Rule:
		return json.Marshal(obj.(*models.Rule))
	default:
		return nil, fmt.Errorf("invalid type")
	}
}

// Decode turns JSON into object
func (c *Codec) Decode(data []byte, objRef models.Object) error {
	switch objRef.(type) {
	case *models.Parameter:
		return json.Unmarshal(data, objRef.(*models.Parameter))
	case *models.Action:
		return json.Unmarshal(data, objRef.(*models.Action))
	case *models.ActionList:
		return json.Unmarshal(data, objRef.(*models.ActionList))
	case *models.Event:
		return json.Unmarshal(data, objRef.(*models.Event))
	case *models.Rule:
		return json.Unmarshal(data, objRef.(*models.Rule))
	default:
		return fmt.Errorf("invalid type")
	}
}
