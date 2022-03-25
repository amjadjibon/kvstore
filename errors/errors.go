package errors

import (
	"github.com/mkawserm/abesh/errors"
)

var (
	ErrEncodingNotFound = errors.New(1, "ABESH_KVSTORE_ERROR", "Encoding not found", nil)
)
