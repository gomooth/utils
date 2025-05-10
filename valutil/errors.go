package valutil

import "errors"

var (
	ErrInvalidIntString  = errors.New("invalid integer string")
	ErrInvalidBoolString = errors.New("invalid boolean string")
	ErrUnsupportedType   = errors.New("unsupported type for conversion")
)
