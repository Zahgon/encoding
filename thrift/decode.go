package thrift

import (
	"io"
	"reflect"
	"sync/atomic"
)

// Unmarshal deserializes the thrift data from b to v using to the protocol p.
//
// The function errors if the data in b does not match the type of v.
//
// The function panics if v cannot be converted to a thrift representation.
//
// As an optimization, the value passed in v may be reused across multiple calls
// to Unmarshal, allowing the function to reuse objects referenced by pointer
// fields of struct values. When reusing objects, the application is responsible
// for resetting the state of v before calling Unmarshal again.
func Unmarshal(p Protocol, b []byte, v any) error { _ = "STUB: not implemented"; return nil }

type Decoder struct {
	r Reader
	f flags
}

func NewDecoder(r Reader) *Decoder { _ = "STUB: not implemented"; return nil }

func (d *Decoder) Decode(v any) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) Reset(r Reader) { _ = "STUB: not implemented"; return }

func (d *Decoder) SetStrict(enabled bool) { _ = "STUB: not implemented"; return }

func decoderFlags(r Reader) flags { _ = "STUB: not implemented"; return *new(flags) }

var decoderCache atomic.Value // map[typeID]decodeFunc

type decodeFunc func(Reader, reflect.Value, flags) error

type decodeFuncCache map[reflect.Type]decodeFunc

func decodeFuncOf(t reflect.Type, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// []byte

func decodeBool(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeInt8(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeInt16(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeInt32(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeInt64(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeFloat64(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeString(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeBytes(r Reader, v reflect.Value, _ flags) error { _ = "STUB: not implemented"; return nil }

func decodeFuncSliceOf(t reflect.Type, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// Sometimes the list type is set to TRUE when the list contains only
// TRUE values. Thrift does not seem to optimize the encoding by
// omitting the boolean values that are known to all be TRUE, we still
// need to decode them.

// TODO: implement type conversions?

func decodeFuncMapOf(t reflect.Type, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// map[?]struct{}

// empty map

// TODO: implement type conversions?

func decodeFuncMapAsSetOf(t reflect.Type, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// See decodeFuncSliceOf for details about why this type conversion
// needs to be done.

// TODO: implement type conversions?

type structDecoder struct {
	fields   []structDecoderField
	union    []int
	minID    int16
	zero     reflect.Value
	required []uint64
}

func (dec *structDecoder) decode(r Reader, v reflect.Value, flags flags) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: implement type conversions?

type structDecoderField struct {
	index  []int
	id     int16
	flags  flags
	typ    Type
	decode decodeFunc
}

func decodeFuncStructOf(t reflect.Type, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func decodeFuncStructFieldOf(f structField, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func decodeFuncPtrOf(t reflect.Type, seen decodeFuncCache) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func readBinary(r Reader, f func(io.Reader) error) error { _ = "STUB: not implemented"; return nil }

func readList(r Reader, f func(Reader, Type) error) error { _ = "STUB: not implemented"; return nil }

func readSet(r Reader, f func(Reader, Type) error) error { _ = "STUB: not implemented"; return nil }

func readMap(r Reader, f func(Reader, Type, Type) error) error {
	_ = "STUB: not implemented"
	return nil
}

func readStruct(r Reader, f func(Reader, Field) error) error { _ = "STUB: not implemented"; return nil }

func skip(r Reader, t Type) error { _ = "STUB: not implemented"; return nil }

func skipBinary(r Reader) error { _ = "STUB: not implemented"; return nil }

func skipList(r Reader) error { _ = "STUB: not implemented"; return nil }

func skipSet(r Reader) error { _ = "STUB: not implemented"; return nil }

func skipMap(r Reader) error { _ = "STUB: not implemented"; return nil }

func skipStruct(r Reader) error { _ = "STUB: not implemented"; return nil }

func skipField(r Reader, f Field) error { _ = "STUB: not implemented"; return nil }
