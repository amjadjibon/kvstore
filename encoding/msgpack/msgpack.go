package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"

	"github.com/amjadjibon/kvstore/registry"
)

type EncodingMsgpack struct{}

func (e EncodingMsgpack) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (e EncodingMsgpack) Unmarshal(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b, v)
}

func init() {
	registry.EncodingRegistry().AddEncoding("msgpack", EncodingMsgpack{})
}
