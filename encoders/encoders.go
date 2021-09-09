package encoders

import (
	"gitlab.upay.dev/golang/kvstore/model"
)

func Marshal(encoder Encoder, v *model.InputModel) ([]byte, error) {
	return encoder.Marshal(v)
}

func Unmarshal(encoder Encoder, data []byte, v *model.InputModel) error {
	return encoder.Unmarshal(data, v)
}
