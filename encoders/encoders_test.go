package encoders

import (
	"fmt"
	"gitlab.upay.dev/golang/kvstore/model"
	"testing"
	"time"
)

func TestJSONMarshalUnmarshal(t *testing.T) {
	v1 := &model.KVStore{
		Key: "key",
		Value: []byte("value"),
		CreatedAt: uint64(time.Now().Unix()),
		UpdatedAt: uint64(time.Now().Unix()),
	}
	v2 := &model.KVStore{}

	marshal, err := Marshal(JSONCodec{}, v1)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Encoded Byte: %v\nByte to String: %v\n", marshal, string(marshal))

	err = Unmarshal(JSONCodec{}, marshal, v2)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Decoded Data: %v\n", v2)
}

func TestGOBMarshalUnmarshal(t *testing.T) {
	v1 := &model.KVStore{
		Key: "key",
		Value: []byte("value"),
		CreatedAt: uint64(time.Now().Unix()),
		UpdatedAt: uint64(time.Now().Unix()),
	}
	v2 := &model.KVStore{}

	marshal, err := Marshal(GOBCodec{}, v1)

	fmt.Printf("Encoded Byte: %v\nByte to String: %v\n", marshal, string(marshal))

	if err != nil {
		t.Error(err)
	}

	err = Unmarshal(GOBCodec{}, marshal, v2)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Decoded Data: %v\n", v2)
}

func TestMsgPackMarshalUnmarshal(t *testing.T) {
	v1 := &model.KVStore{
		Key: "key",
		Value: []byte("value"),
		CreatedAt: uint64(time.Now().Unix()),
		UpdatedAt: uint64(time.Now().Unix()),
	}
	v2 := &model.KVStore{}

	marshal, err := Marshal(MessagePackCodec{}, v1)

	fmt.Printf("Encoded Byte: %v\nByte to String: %v\n", marshal, string(marshal))

	if err != nil {
		t.Error(err)
	}

	err = Unmarshal(MessagePackCodec{}, marshal, v2)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Decoded Data: %v\n", v2)
}

func TestProtobufMarshalUnmarshal(t *testing.T) {
	v1 := &model.KVStore{
		Key: "key",
		Value: []byte("value"),
		CreatedAt: uint64(time.Now().Unix()),
		UpdatedAt: uint64(time.Now().Unix()),
	}
	v2 := &model.KVStore{}

	marshal, err := Marshal(ProtobufCodec{}, v1)

	fmt.Printf("Encoded Byte: %v\nByte to String: %v\n", marshal, string(marshal))

	if err != nil {
		t.Error(err)
	}

	err = Unmarshal(ProtobufCodec{}, marshal, v2)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Decoded Data: %v\n", v2)
}

func BenchmarkJSONMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := &model.KVStore{
			Key: "key",
			Value: []byte("value"),
			CreatedAt: uint64(time.Now().Unix()),
			UpdatedAt: uint64(time.Now().Unix()),
		}
		v2 := &model.KVStore{}

		marshal, err := Marshal(JSONCodec{}, v1)

		if err != nil {
			return
		}

		err = Unmarshal(JSONCodec{}, marshal, v2)

		if err != nil {
			return
		}
	}
}

func BenchmarkGOBMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := &model.KVStore{
			Key: "key",
			Value: []byte("value"),
			CreatedAt: uint64(time.Now().Unix()),
			UpdatedAt: uint64(time.Now().Unix()),
		}
		v2 := &model.KVStore{}

		marshal, err := Marshal(GOBCodec{}, v1)

		if err != nil {
			return
		}

		err = Unmarshal(GOBCodec{}, marshal, v2)

		if err != nil {
			return
		}
	}
}

func BenchmarkMsgPackMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := &model.KVStore{
			Key: "key",
			Value: []byte("value"),
			CreatedAt: uint64(time.Now().Unix()),
			UpdatedAt: uint64(time.Now().Unix()),
		}
		v2 := &model.KVStore{}

		marshal, err := Marshal(MessagePackCodec{}, v1)

		if err != nil {
			return
		}

		err = Unmarshal(MessagePackCodec{}, marshal, v2)

		if err != nil {
			return
		}
	}
}

func BenchmarkProtobufMarshalUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := &model.KVStore{
			Key: "key",
			Value: []byte("value"),
			CreatedAt: uint64(time.Now().Unix()),
			UpdatedAt: uint64(time.Now().Unix()),
		}
		v2 := &model.KVStore{}

		marshal, err := Marshal(ProtobufCodec{}, v1)

		if err != nil {
			return
		}

		err = Unmarshal(ProtobufCodec{}, marshal, v2)

		if err != nil {
			return
		}
	}
}
