package encoders

import (
	"gitlab.upay.dev/golang/kvstore/model"
)

type Encoder interface {
	Marshal(v *model.InputModel) ([]byte, error)
	Unmarshal(data []byte, v *model.InputModel) error
}
