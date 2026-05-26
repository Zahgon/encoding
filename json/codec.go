package json

import (
	"encoding"
	"encoding/json"
	"math/big"
	"reflect"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	// 1000 is the value used by the standard encoding/json package.
	//
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/encoding/json/encode.go;drc=refs%2Ftags%2Fgo1.17.3;l=300
	startDetectingCyclesAfter = 1000
)

type codec struct {
	encode encodeFunc
	decode decodeFunc
}

type encoder struct {
	flags AppendFlags
	// ptrDepth tracks the depth of pointer cycles, when it reaches the value
	// of startDetectingCyclesAfter, the ptrSeen map is allocated and the
	// encoder starts tracking pointers it has seen as an attempt to detect
	// whether it has entered a pointer cycle and needs to error before the
	// goroutine runs out of stack space.
	ptrDepth uint32
	ptrSeen  map[unsafe.Pointer]struct{}
}

type decoder struct {
	flags ParseFlags
}

type (
	encodeFunc func(encoder, []byte, unsafe.Pointer) ([]byte, error)
	decodeFunc func(decoder, []byte, unsafe.Pointer) ([]byte, error)
)

type (
	emptyFunc func(unsafe.Pointer) bool
	sortFunc  func([]reflect.Value)
)

// Eventually consistent cache mapping go types to dynamically generated
// codecs.
//
// Note: using a uintptr as key instead of reflect.Type shaved ~15ns off of
// the ~30ns Marhsal/Unmarshal functions which were dominated by the map
// lookup time for simple types like bool, int, etc..
var cache atomic.Pointer[map[unsafe.Pointer]codec]

func cacheLoad() map[unsafe.Pointer]codec { _ = "STUB: not implemented"; return nil }

func cacheStore(typ reflect.Type, cod codec, oldCodecs map[unsafe.Pointer]codec) {
	_ = "STUB: not implemented"
	return
}

func typeid(t reflect.Type) unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func constructCachedCodec(t reflect.Type, cache map[unsafe.Pointer]codec) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructCodec(t reflect.Type, seen map[reflect.Type]*structType, canAddr bool) (c codec) {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructStringCodec(t reflect.Type, seen map[reflect.Type]*structType, canAddr bool) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructStringEncodeFunc(encode encodeFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructStringDecodeFunc(decode decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructStringToIntDecodeFunc(t reflect.Type, decode decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructArrayCodec(t reflect.Type, seen map[reflect.Type]*structType, canAddr bool) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructArrayEncodeFunc(size uintptr, t reflect.Type, encode encodeFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructArrayDecodeFunc(size uintptr, t reflect.Type, decode decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructSliceCodec(t reflect.Type, seen map[reflect.Type]*structType) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

// Go 1.7+ behavior: slices of byte types (and aliases) may override the
// default encoding and decoding behaviors by implementing marshaler and
// unmarshaler interfaces.

func constructSliceEncodeFunc(size uintptr, t reflect.Type, encode encodeFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructSliceDecodeFunc(size uintptr, t reflect.Type, decode decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructMapCodec(t reflect.Type, seen map[reflect.Type]*structType) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

// Faster implementations for some common cases.

// This is a performance abomination but the use case is rare
// enough that it shouldn't be a problem in practice.

func constructMapEncodeFunc(t reflect.Type, encodeKey, encodeValue encodeFunc, sortKeys sortFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructMapDecodeFunc(t reflect.Type, decodeKey, decodeValue decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructStructCodec(t reflect.Type, seen map[reflect.Type]*structType, canAddr bool) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructStructType(t reflect.Type, seen map[reflect.Type]*structType, canAddr bool) *structType {
	_ = "STUB: not implemented"
	// Used for preventing infinite recursion on types that have pointers to
	// themselves.
	return nil
}

// When there is ambiguity because multiple fields have the same
// case-insensitive representation, the first field must win.

// At a certain point the linear scan provided by keyset is less
// efficient than a map. The 32 was chosen based on benchmarks in the
// segmentio/asm repo run with an Intel Kaby Lake processor and go1.17.

func constructStructEncodeFunc(st *structType) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructStructDecodeFunc(st *structType) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructEmbeddedStructPointerCodec(t reflect.Type, unexported bool, offset uintptr, field codec) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructEmbeddedStructPointerEncodeFunc(t reflect.Type, unexported bool, offset uintptr, encode encodeFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructEmbeddedStructPointerDecodeFunc(t reflect.Type, unexported bool, offset uintptr, decode decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func appendStructFields(fields []structField, t reflect.Type, offset uintptr, seen map[reflect.Type]*structType, canAddr bool) []structField {
	_ = "STUB: not implemented"
	return nil
}

// unexported

// ignored

// embedded

// When the embedded fields is inlined the fields can be looked
// up by offset from the address of the wrapping object, so we
// simply add the embedded struct fields to the list of fields
// of the current struct type.

// ignore unexported non-struct types

// https://golang.org/pkg/encoding/json/#Marshal
//
// The "string" option signals that a field is stored as JSON inside
// a JSON-encoded string. It applies only to fields of string,
// floating point, integer, or boolean types. This extra level of
// encoding is sometimes used when communicating with JavaScript
// programs:

// Only unambiguous embedded fields must be serialized.

// Embedded types can never override a field that was already present at
// the top-level.

// ambiguous embedded field

// To prevent dominant flags more than one level below the embedded one.

// To ensure the order of the fields in the output is the same is in the
// struct type.

func encodeKeyFragment(s string, flags AppendFlags) string { _ = "STUB: not implemented"; return "" }

func constructPointerCodec(t reflect.Type, seen map[reflect.Type]*structType) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructPointerEncodeFunc(t reflect.Type, encode encodeFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructPointerDecodeFunc(t reflect.Type, decode decodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructInterfaceCodec(t reflect.Type) codec { _ = "STUB: not implemented"; return *new(codec) }

func constructMaybeEmptyInterfaceEncoderFunc(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructMaybeEmptyInterfaceDecoderFunc(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructUnsupportedTypeCodec(t reflect.Type) codec {
	_ = "STUB: not implemented"
	return *new(codec)
}

func constructUnsupportedTypeEncodeFunc(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructUnsupportedTypeDecodeFunc(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructJSONMarshalerEncodeFunc(t reflect.Type, pointer bool) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructJSONUnmarshalerDecodeFunc(t reflect.Type, pointer bool) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructTextMarshalerEncodeFunc(t reflect.Type, pointer bool) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func constructTextUnmarshalerDecodeFunc(t reflect.Type, pointer bool) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func constructInlineValueEncodeFunc(encode encodeFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input. noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func alignedSize(t reflect.Type) uintptr { _ = "STUB: not implemented"; return 0 }

func align(align, size uintptr) uintptr { _ = "STUB: not implemented"; return 0 }

func inlined(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isValidTag(s string) bool { _ = "STUB: not implemented"; return false }

// Backslash and quote chars are reserved, but
// otherwise any punctuation chars are allowed
// in a tag name.

func emptyFuncOf(t reflect.Type) emptyFunc { _ = "STUB: not implemented"; return *new(emptyFunc) }

type iface struct {
	typ unsafe.Pointer
	ptr unsafe.Pointer
}

type slice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

type structType struct {
	fields      []structField
	fieldsIndex map[string]*structField
	ficaseIndex map[string]*structField
	keyset      []byte
	typ         reflect.Type
}

type structField struct {
	codec     codec
	offset    uintptr
	empty     emptyFunc
	tag       bool
	omitempty bool
	json      string
	html      string
	name      string
	typ       reflect.Type
	zero      reflect.Value
	index     int
}

func unmarshalTypeError(b []byte, t reflect.Type) error { _ = "STUB: not implemented"; return nil }

func unmarshalOverflow(b []byte, t reflect.Type) error { _ = "STUB: not implemented"; return nil }

func unexpectedEOF(b []byte) error { _ = "STUB: not implemented"; return nil }

var syntaxErrorMsgOffset = ^uintptr(0)

func init() {
	t := reflect.TypeOf(SyntaxError{})
	for i := range t.NumField() {
		if f := t.Field(i); f.Type.Kind() == reflect.String {
			syntaxErrorMsgOffset = f.Offset
		}
	}
}

func syntaxError(b []byte, msg string, args ...any) error { _ = "STUB: not implemented"; return nil }

// Hack to set the unexported `msg` field.

func objectKeyError(b []byte, err error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prefix(b []byte) string { _ = "STUB: not implemented"; return "" }

func intStringsAreSorted(i0, i1 int64) bool { _ = "STUB: not implemented"; return false }

func uintStringsAreSorted(u0, u1 uint64) bool { _ = "STUB: not implemented"; return false }

func stringToBytes(s string) []byte { _ = "STUB: not implemented"; return nil }

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

var (
	nullType = reflect.TypeOf(nil)
	boolType = reflect.TypeOf(false)

	intType   = reflect.TypeOf(int(0))
	int8Type  = reflect.TypeOf(int8(0))
	int16Type = reflect.TypeOf(int16(0))
	int32Type = reflect.TypeOf(int32(0))
	int64Type = reflect.TypeOf(int64(0))

	uintType    = reflect.TypeOf(uint(0))
	uint8Type   = reflect.TypeOf(uint8(0))
	uint16Type  = reflect.TypeOf(uint16(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	uint64Type  = reflect.TypeOf(uint64(0))
	uintptrType = reflect.TypeOf(uintptr(0))

	float32Type = reflect.TypeOf(float32(0))
	float64Type = reflect.TypeOf(float64(0))

	bigIntType     = reflect.TypeOf(new(big.Int))
	numberType     = reflect.TypeOf(json.Number(""))
	stringType     = reflect.TypeOf("")
	stringsType    = reflect.TypeOf([]string(nil))
	bytesType      = reflect.TypeOf(([]byte)(nil))
	durationType   = reflect.TypeOf(time.Duration(0))
	timeType       = reflect.TypeOf(time.Time{})
	rawMessageType = reflect.TypeOf(RawMessage(nil))

	numberPtrType     = reflect.PointerTo(numberType)
	durationPtrType   = reflect.PointerTo(durationType)
	timePtrType       = reflect.PointerTo(timeType)
	rawMessagePtrType = reflect.PointerTo(rawMessageType)

	sliceInterfaceType       = reflect.TypeOf(([]any)(nil))
	sliceStringType          = reflect.TypeOf(([]any)(nil))
	mapStringInterfaceType   = reflect.TypeOf((map[string]any)(nil))
	mapStringRawMessageType  = reflect.TypeOf((map[string]RawMessage)(nil))
	mapStringStringType      = reflect.TypeOf((map[string]string)(nil))
	mapStringStringSliceType = reflect.TypeOf((map[string][]string)(nil))
	mapStringBoolType        = reflect.TypeOf((map[string]bool)(nil))

	interfaceType       = reflect.TypeOf((*any)(nil)).Elem()
	jsonMarshalerType   = reflect.TypeOf((*Marshaler)(nil)).Elem()
	jsonUnmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
	textMarshalerType   = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()

	bigIntDecoder = constructJSONUnmarshalerDecodeFunc(bigIntType, false)
)

// =============================================================================
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// appendDuration appends a human-readable representation of d to b.
//
// The function copies the implementation of time.Duration.String but prevents
// Go from making a dynamic memory allocation on the returned value.
func appendDuration(b []byte, d time.Duration) []byte {
	_ = "STUB: not implemented"
	// Largest time is 2540400h10m10.000000000s
	return nil
}

// Special case: if duration is smaller than a second,
// use smaller units, like 1.2ms

// print nanoseconds

// print microseconds

// U+00B5 'µ' micro sign == 0xC2 0xB5
// Need room for two bytes.

// print milliseconds

// u is now integer seconds

// u is now integer minutes

// u is now integer hours
// Stop at hours because days can be different lengths.

// fmtFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros.  it omits the decimal
// point too when the fraction is 0.  It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	_ = "STUB: not implemented"
	// Omit trailing zeros up to and including decimal point.
	return 0, 0
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int { _ = "STUB: not implemented"; return 0 }

// =============================================================================
