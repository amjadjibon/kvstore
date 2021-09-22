package encoders

import (
	"github.com/vmihailenco/msgpack/v5"
)

// MessagePackCodec encodes/decodes Go values to/from JSON.
// You can use encoding.JSON instead of creating an instance of this struct.
type MessagePackCodec struct{}

// Marshal encodes a Go value to JSON.
func (c MessagePackCodec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal decodes a JSON value into a Go value.
func (c MessagePackCodec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}