package yaml

import (
	"fmt"

	"github.com/joshuakwan/hydra/models"
	yaml "gopkg.in/yaml.v2"
)

// Codec encode/decode objects against YAML
type Codec struct {
}

// NewYAMLCodec creates a new YAML codec
func NewYAMLCodec() *Codec {
	return &Codec{}
}

// Encode turns object into YAML
func (c *Codec) Encode(obj models.Object) ([]byte, error) {
	switch obj.ObjectType() {
	case "Parameter":
		return yaml.Marshal(obj.(*models.Parameter))
	case "Action":
		return yaml.Marshal(obj.(*models.Action))
	case "ActionList":
		return yaml.Marshal(obj.(*models.ActionList))
	case "Event":
		return yaml.Marshal(obj.(*models.Event))
	case "Rule":
		return yaml.Marshal(obj.(*models.Rule))
	default:
		return nil, fmt.Errorf("invalid type")
	}
}

// Decode turns YAML into object
func (c *Codec) Decode(data []byte, objRef models.Object) error {
	switch objRef.ObjectType() {
	case "Parameter":
		return yaml.Unmarshal(data, objRef.(*models.Parameter))
	case "Action":
		return yaml.Unmarshal(data, objRef.(*models.Action))
	case "ActionList":
		return yaml.Unmarshal(data, objRef.(*models.ActionList))
	case "Event":
		return yaml.Unmarshal(data, objRef.(*models.Event))
	case "Rule":
		return yaml.Unmarshal(data, objRef.(*models.Rule))
	default:
		return fmt.Errorf("invalid type")
	}
}
