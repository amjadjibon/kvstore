package proto

import (
	"google.golang.org/protobuf/proto"

	"github.com/amjadjibon/kvstore/registry"
)

type EncodingProto struct{}

func (e EncodingProto) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (e EncodingProto) Unmarshal(b []byte, v interface{}) error {
	return proto.Unmarshal(b, v.(proto.Message))
}

func init() {
	registry.EncodingRegistry().AddEncoding("proto", EncodingProto{})
}
