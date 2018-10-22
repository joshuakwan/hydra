package codec

import "github.com/joshuakwan/hydra/models"

// Codec is the interface of marshal/unmarshal objects
type Codec interface {
	Encode(obj models.Object) ([]byte, error)
	Decode(data []byte, objRef models.Object) error
}
