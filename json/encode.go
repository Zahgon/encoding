package json

import (
	"reflect"
	"sync"
	"unsafe"
)

const hex = "0123456789abcdef"

func (e encoder) encodeNull(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeBool(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeInt(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeInt8(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeInt16(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeInt32(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeInt64(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUint(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUintptr(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUint8(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUint16(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUint32(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUint64(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeFloat32(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeFloat64(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeFloat(b []byte, f float64, bits int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert as if by ES6 number to string conversion.
// This matches most other JSON generators.
// See golang.org/issue/6384 and golang.org/issue/14135.
// Like fmt %g, but the exponent cutoffs are different
// and exponents themselves are not padded to two digits.

// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.

// clean up e-09 to e-9

func (e encoder) encodeNumber(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeString(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fast path: most of the time, printable ascii characters are used

// This encodes bytes < 0x20 except for \t, \n and \r.

// U+2028 is LINE SEPARATOR.
// U+2029 is PARAGRAPH SEPARATOR.
// They are both technically valid characters in JSON strings,
// but don't work in JSONP, which has to be evaluated as JavaScript,
// and can lead to security holes there. It is valid JSON to
// escape them, so we do so unconditionally.
// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.

func (e encoder) encodeToString(b []byte, p unsafe.Pointer, encode encodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeBytes(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeDuration(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeTime(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeArray(b []byte, p unsafe.Pointer, n int, size uintptr, t reflect.Type, encode encodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeSlice(b []byte, p unsafe.Pointer, size uintptr, t reflect.Type, encode encodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeMap(b []byte, p unsafe.Pointer, t reflect.Type, encodeKey, encodeValue encodeFunc, sortKeys sortFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type element struct {
	key string
	val any
	raw RawMessage
}

type mapslice struct {
	elements []element
}

func (m *mapslice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (m *mapslice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (m *mapslice) Swap(i, j int)      { _ = "STUB: not implemented"; return }

var mapslicePool = sync.Pool{
	New: func() any { return new(mapslice) },
}

func (e encoder) encodeMapStringInterface(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimized code path when the program does not need the map keys to be
// sorted.

func (e encoder) encodeMapStringRawMessage(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimized code path when the program does not need the map keys to be
// sorted.

// encodeString doesn't return errors so we ignore it here

func (e encoder) encodeMapStringString(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimized code path when the program does not need the map keys to be
// sorted.

// encodeString never returns an error so we ignore it here

// encodeString never returns an error so we ignore it here

func (e encoder) encodeMapStringStringSlice(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimized code path when the program does not need the map keys to be
// sorted.

func (e encoder) encodeMapStringBool(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimized code path when the program does not need the map keys to be
// sorted.

// encodeString never returns an error so we ignore it here

// encodeString never returns an error so we ignore it here

func (e encoder) encodeStruct(b []byte, p unsafe.Pointer, st *structType) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rollback struct{}

func (rollback) Error() string { _ = "STUB: not implemented"; return "" }

func (e encoder) encodeEmbeddedStructPointer(b []byte, p unsafe.Pointer, t reflect.Type, unexported bool, offset uintptr, encode encodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodePointer(b []byte, p unsafe.Pointer, t reflect.Type, encode encodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: reconstruct the reflect.Value from p + t so we can set
// the erorr's Value field?

func (e encoder) encodeInterface(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeMaybeEmptyInterface(b []byte, p unsafe.Pointer, t reflect.Type) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeUnsupportedTypeError(b []byte, p unsafe.Pointer, t reflect.Type) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeRawMessage(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// don't assume that a RawMessage starts with a token.

func (e encoder) encodeJSONMarshaler(b []byte, p unsafe.Pointer, t reflect.Type, pointer bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e encoder) encodeTextMarshaler(b []byte, p unsafe.Pointer, t reflect.Type, pointer bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendCompactEscapeHTML(dst []byte, src []byte) []byte { _ = "STUB: not implemented"; return nil }

// enter string

// skip space

// Convert U+2028 and U+2029 (E2 80 A8 and E2 80 A9).
