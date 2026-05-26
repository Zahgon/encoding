package proto

import (
	"reflect"
	"unsafe"
)

const (
	embedded = 1 << 0
	repeated = 1 << 1
	zigzag   = 1 << 2
)

type structField struct {
	number  uint16
	tagsize uint8
	flags   uint8
	offset  uint32
	codec   *codec
}

func (f *structField) String() string { _ = "STUB: not implemented"; return "" }

func (f *structField) fieldNumber() fieldNumber {
	_ = "STUB: not implemented"
	return *new(fieldNumber)
}

func (f *structField) wireType() wireType { _ = "STUB: not implemented"; return *new(wireType) }

func (f *structField) embedded() bool { _ = "STUB: not implemented"; return false }

func (f *structField) repeated() bool { _ = "STUB: not implemented"; return false }

func (f *structField) pointer(p unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (f *structField) makeFlags(base flags) flags { _ = "STUB: not implemented"; return *new(flags) }

func structCodecOf(t reflect.Type, seen map[reflect.Type]*codec) *codec {
	_ = "STUB: not implemented"
	return nil
}

// unexported

// []byte

func baseKindOf(t reflect.Type) reflect.Kind { _ = "STUB: not implemented"; return *new(reflect.Kind) }

func baseTypeOf(t reflect.Type) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func structSizeFuncOf(t reflect.Type, fields []structField) sizeFunc {
	_ = "STUB: not implemented"
	return *new(sizeFunc)
}

func structEncodeFuncOf(t reflect.Type, fields []structField) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func structDecodeFuncOf(t reflect.Type, fields []structField) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// `data` will only contain the section of the input buffer where
// the data for the next field is available. This is necessary to
// limit how many bytes will be consumed by embedded messages.
