package encoders

import (
	"bytes"
	"encoding/gob"
)

// GOBCodec encodes/decodes Go values to/from gob.
// You can use encoding.Gob instead of creating an instance of this struct.
type GOBCodec struct{}

// Marshal encodes a Go value to gob.
func (c GOBCodec) Marshal(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Unmarshal decodes a gob value into a Go value.
func (c GOBCodec) Unmarshal(data []byte, v interface{}) error {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(v)
}
