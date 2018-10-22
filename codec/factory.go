package codec

import (
	jsonCodec "github.com/joshuakwan/hydra/codec/json"
)

// Type refers to the data type of a codec
type Type string

// NewCodec returns a new Codec upon the given type
func NewCodec(codecType Type) Codec {
	switch codecType {
	case "json":
		return jsonCodec.NewJSONCodec()
	default:
		return jsonCodec.NewJSONCodec()
	}
}
