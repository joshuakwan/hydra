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
	switch obj.(type) {
	case *models.Parameter:
		return yaml.Marshal(obj.(*models.Parameter))
	case *models.Action:
		return yaml.Marshal(obj.(*models.Action))
	case *models.ActionList:
		return yaml.Marshal(obj.(*models.ActionList))
	case *models.Event:
		return yaml.Marshal(obj.(*models.Event))
	case *models.Rule:
		return yaml.Marshal(obj.(*models.Rule))
	case *models.Worker:
		return yaml.Marshal(obj.(*models.Worker))
	default:
		return nil, fmt.Errorf("invalid type")
	}
}

// Decode turns YAML into object
func (c *Codec) Decode(data []byte, objRef models.Object) error {
	switch objRef.(type) {
	case *models.Parameter:
		return yaml.Unmarshal(data, objRef.(*models.Parameter))
	case *models.Action:
		return yaml.Unmarshal(data, objRef.(*models.Action))
	case *models.ActionList:
		return yaml.Unmarshal(data, objRef.(*models.ActionList))
	case *models.Event:
		return yaml.Unmarshal(data, objRef.(*models.Event))
	case *models.Rule:
		return yaml.Unmarshal(data, objRef.(*models.Rule))
	case *models.Worker:
		return yaml.Unmarshal(data, objRef.(*models.Worker))
	default:
		return fmt.Errorf("invalid type")
	}
}
