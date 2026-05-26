package proto

import (
	"errors"
)

var ErrWireTypeUnknown = errors.New("unknown wire type")

type UnmarshalFieldError struct {
	FieldNumer int
	WireType   int
	Err        error
}

func (e *UnmarshalFieldError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *UnmarshalFieldError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func fieldError(f fieldNumber, t wireType, err error) error { _ = "STUB: not implemented"; return nil }
