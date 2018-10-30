package codec

import (
	jsonCodec "github.com/joshuakwan/hydra/codec/json"
	yamlCodec "github.com/joshuakwan/hydra/codec/yaml"
)

// Type refers to the data type of a codec
type Type string

// NewCodec returns a new Codec upon the given type
func NewCodec(codecType Type) Codec {
	switch codecType {
	case "json":
		return jsonCodec.NewJSONCodec()
	case "yaml":
		return yamlCodec.NewYAMLCodec()
	default:
		return jsonCodec.NewJSONCodec()
	}
}
