package thrift

import (
	"reflect"
	"sync/atomic"
)

// Marshal serializes v into a thrift representation according to the the
// protocol p.
//
// The function panics if v cannot be converted to a thrift representation.
func Marshal(p Protocol, v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type Encoder struct {
	w Writer
	f flags
}

func NewEncoder(w Writer) *Encoder { _ = "STUB: not implemented"; return nil }

func (e *Encoder) Encode(v any) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) Reset(w Writer) { _ = "STUB: not implemented"; return }

func encoderFlags(w Writer) flags { _ = "STUB: not implemented"; return *new(flags) }

var encoderCache atomic.Value // map[typeID]encodeFunc

type encodeFunc func(Writer, reflect.Value, flags) error

type encodeFuncCache map[reflect.Type]encodeFunc

func encodeFuncOf(t reflect.Type, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func encodeBool(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeInt8(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeInt16(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeInt32(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeInt64(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeFloat64(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeString(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeBytes(w Writer, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func encodeFuncSliceOf(t reflect.Type, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func encodeFuncMapOf(t reflect.Type, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

// map[?]struct{}

// empty map

func encodeFuncMapAsSetOf(t reflect.Type, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

// empty map

type structEncoder struct {
	fields []structEncoderField
	union  bool
}

func dereference(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func isTrue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func (enc *structEncoder) encode(w Writer, v reflect.Value, flags flags) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *structEncoder) String() string { _ = "STUB: not implemented"; return "" }

type structEncoderField struct {
	index  []int
	id     int16
	flags  flags
	typ    Type
	encode encodeFunc
}

func encodeFuncStructOf(t reflect.Type, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func encodeFuncStructFieldOf(f structField, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func encodeFuncPtrOf(t reflect.Type, seen encodeFuncCache) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}
