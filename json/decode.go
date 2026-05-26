package json

import (
	"reflect"
	"unsafe"
)

func (d decoder) anyFlagsSet(flags ParseFlags) bool { _ = "STUB: not implemented"; return false }

func (d decoder) decodeNull(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeBool(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeInt(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeInt8(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeInt16(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeInt32(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeInt64(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeUint(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeUintptr(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeUint8(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeUint16(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeUint32(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeUint64(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeFloat32(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeFloat64(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeNumber(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeString(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeFromString(b []byte, p unsafe.Pointer, decode decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeFromStringToInt(b []byte, p unsafe.Pointer, t reflect.Type, decode decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The encoding/json package will return a *json.UnmarshalTypeError if
// the input was a floating point number representation, even tho a
// string is expected here.

// In this context the encoding/json package accepts leading zeroes because
// it is not constrained by the JSON syntax, remove them so the parsing
// functions don't return syntax errors.

// The standard library interprets sequences of '-' characters
// as numbers but still returns type errors in this case...

// When the input value was a valid number representation we retain the
// error returned by the decoder.

// When the input value valid JSON we mirror the behavior of the
// encoding/json package and return a generic error.

func (d decoder) decodeBytes(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Go 1.7- behavior: bytes slices may be decoded from array of integers.

// The input string contains escaped sequences, we need to parse it before
// decoding it to match the encoding/json package behvaior.

func (d decoder) decodeDuration(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// in order to inter-operate with the stdlib, we must be able to interpret
// durations passed as integer values.  there's some discussion about being
// flexible on how durations are formatted, but for the time being, it's
// been punted to go2 at the earliest: https://github.com/golang/go/issues/4712

// trim quotes

func (d decoder) decodeTime(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// trim quotes

func (d decoder) decodeArray(b []byte, p unsafe.Pointer, n int, size uintptr, t reflect.Type, decode decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The encoding/json package ignores extra elements found when decoding into
// array types (which have a fixed size).

// This is a placeholder used to consturct non-nil empty slices.
var empty struct{}

func (d decoder) decodeSlice(b []byte, p unsafe.Pointer, size uintptr, t reflect.Type, decode decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Go 1.7- behavior: fallback to decoding as a []byte if the element
// type is byte; allow conversions from JSON strings even tho the
// underlying type implemented unmarshaler interfaces.

func (d decoder) decodeMap(b []byte, p unsafe.Pointer, t, kt, vt reflect.Type, kz, vz reflect.Value, decodeKey, decodeValue decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeMapStringInterface(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeMapStringRawMessage(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeMapStringString(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeMapStringStringSlice(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeMapStringBool(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeStruct(b []byte, p unsafe.Pointer, st *structType) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// memory buffer used to convert short field names to lowercase

func (d decoder) decodeEmbeddedStructPointer(b []byte, p unsafe.Pointer, t reflect.Type, unexported bool, offset uintptr, decode decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodePointer(b []byte, p unsafe.Pointer, t reflect.Type, decode decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeInterface(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the destination is nil the only value that is OK to decode is
// `null`, and the encoding/json package always nils the destination
// interface value in this case.

func (d decoder) decodeDynamicNumber(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Only pre-parse for numeric kind if a conditional decode
		// has been requested.
		nil
}

// Mutually exclusive integer handling cases.

// If requested, attempt decode of positive integers as uint64.

// If uint64 decode was not requested but int64 decode was requested,
// then attempt decode of positive integers as int64.

// If int64 decode was requested,
// attempt decode of negative integers as int64.

// Fallback numeric handling cases:
// these cannot be combined into the above switch,
// since these cases also handle overflow
// from the above cases, if decode was already attempted.

// If *big.Int decode was requested, handle that case for any integer.

// If json.Number decode was requested, handle that for any number.

// Fall back to float64 decode when no special decoding has been requested.

func (d decoder) decodeMaybeEmptyInterface(b []byte, p unsafe.Pointer, t reflect.Type) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// empty interface

func (d decoder) decodeUnmarshalTypeError(b []byte, _ unsafe.Pointer, t reflect.Type) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeRawMessage(b []byte, p unsafe.Pointer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeJSONUnmarshaler(b []byte, p unsafe.Pointer, t reflect.Type, pointer bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) decodeTextUnmarshaler(b []byte, p unsafe.Pointer, t reflect.Type, pointer bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d decoder) prependField(key, field string) string { _ = "STUB: not implemented"; return "" }

func (d decoder) inputError(b []byte, t reflect.Type) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeInto[T any](dest *any, b []byte, d decoder, fn decodeFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
