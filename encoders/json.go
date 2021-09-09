package encoders

import (
	"encoding/json"
	"gitlab.upay.dev/golang/kvstore/model"
)

// JSONCodec encodes/decodes Go values to/from JSON.
// You can use encoding.JSON instead of creating an instance of this struct.
type JSONCodec struct{}

// Marshal encodes a Go value to JSON.
func (c JSONCodec) Marshal(v *model.InputModel) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal decodes a JSON value into a Go value.
func (c JSONCodec) Unmarshal(data []byte, v *model.InputModel) error {
	return json.Unmarshal(data, v)
}