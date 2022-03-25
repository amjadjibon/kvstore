package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/amjadjibon/kvstore/registry"
)

type EncodingGOB struct{}

func (e EncodingGOB) Marshal(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	var encoder = gob.NewEncoder(&buffer)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (e EncodingGOB) Unmarshal(b []byte, v interface{}) error {
	var reader = bytes.NewReader(b)
	var decoder = gob.NewDecoder(reader)
	return decoder.Decode(v)
}

func init() {
	registry.EncodingRegistry().AddEncoding("gob", EncodingGOB{})
}
