package encoders

import (
	"gitlab.upay.dev/golang/kvstore/model"
	"google.golang.org/protobuf/proto"
)

type ProtobufCodec struct{}

// Marshal encodes a Go value to JSON.
func (c ProtobufCodec) Marshal(v *model.InputModel) ([]byte, error) {
	return proto.Marshal(v)
}

// Unmarshal decodes a JSON value into a Go value.
func (c ProtobufCodec) Unmarshal(data []byte, v *model.InputModel) error {
	return proto.Unmarshal(data, v)
}