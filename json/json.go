package json

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
)

// Delim is documented at https://golang.org/pkg/encoding/json/#Delim
type Delim = json.Delim

// InvalidUTF8Error is documented at https://golang.org/pkg/encoding/json/#InvalidUTF8Error
type InvalidUTF8Error = json.InvalidUTF8Error //nolint:staticcheck // compat.

// InvalidUnmarshalError is documented at https://golang.org/pkg/encoding/json/#InvalidUnmarshalError
type InvalidUnmarshalError = json.InvalidUnmarshalError

// Marshaler is documented at https://golang.org/pkg/encoding/json/#Marshaler
type Marshaler = json.Marshaler

// MarshalerError is documented at https://golang.org/pkg/encoding/json/#MarshalerError
type MarshalerError = json.MarshalerError

// Number is documented at https://golang.org/pkg/encoding/json/#Number
type Number = json.Number

// RawMessage is documented at https://golang.org/pkg/encoding/json/#RawMessage
type RawMessage = json.RawMessage

// A SyntaxError is a description of a JSON syntax error.
type SyntaxError = json.SyntaxError

// Token is documented at https://golang.org/pkg/encoding/json/#Token
type Token = json.Token

// UnmarshalFieldError is documented at https://golang.org/pkg/encoding/json/#UnmarshalFieldError
type UnmarshalFieldError = json.UnmarshalFieldError //nolint:staticcheck // compat.

// UnmarshalTypeError is documented at https://golang.org/pkg/encoding/json/#UnmarshalTypeError
type UnmarshalTypeError = json.UnmarshalTypeError

// Unmarshaler is documented at https://golang.org/pkg/encoding/json/#Unmarshaler
type Unmarshaler = json.Unmarshaler

// UnsupportedTypeError is documented at https://golang.org/pkg/encoding/json/#UnsupportedTypeError
type UnsupportedTypeError = json.UnsupportedTypeError

// UnsupportedValueError is documented at https://golang.org/pkg/encoding/json/#UnsupportedValueError
type UnsupportedValueError = json.UnsupportedValueError

// AppendFlags is a type used to represent configuration options that can be
// applied when formatting json output.
type AppendFlags uint32

const (
	// EscapeHTML is a formatting flag used to to escape HTML in json strings.
	EscapeHTML AppendFlags = 1 << iota

	// SortMapKeys is formatting flag used to enable sorting of map keys when
	// encoding JSON (this matches the behavior of the standard encoding/json
	// package).
	SortMapKeys

	// TrustRawMessage is a performance optimization flag to skip value
	// checking of raw messages. It should only be used if the values are
	// known to be valid json (e.g., they were created by json.Unmarshal).
	TrustRawMessage

	// appendNewline is a formatting flag to enable the addition of a newline
	// in Encode (this matches the behavior of the standard encoding/json
	// package).
	appendNewline
)

// ParseFlags is a type used to represent configuration options that can be
// applied when parsing json input.
type ParseFlags uint32

func (flags ParseFlags) has(f ParseFlags) bool { _ = "STUB: not implemented"; return false }

func (f ParseFlags) kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (f ParseFlags) withKind(kind Kind) ParseFlags {
	_ = "STUB: not implemented"
	return *new(ParseFlags)
}

const (
	// DisallowUnknownFields is a parsing flag used to prevent decoding of
	// objects to Go struct values when a field of the input does not match
	// with any of the struct fields.
	DisallowUnknownFields ParseFlags = 1 << iota

	// UseNumber is a parsing flag used to load numeric values as Number
	// instead of float64.
	UseNumber

	// DontCopyString is a parsing flag used to provide zero-copy support when
	// loading string values from a json payload. It is not always possible to
	// avoid dynamic memory allocations, for example when a string is escaped in
	// the json data a new buffer has to be allocated, but when the `wire` value
	// can be used as content of a Go value the decoder will simply point into
	// the input buffer.
	DontCopyString

	// DontCopyNumber is a parsing flag used to provide zero-copy support when
	// loading Number values (see DontCopyString and DontCopyRawMessage).
	DontCopyNumber

	// DontCopyRawMessage is a parsing flag used to provide zero-copy support
	// when loading RawMessage values from a json payload. When used, the
	// RawMessage values will not be allocated into new memory buffers and
	// will instead point directly to the area of the input buffer where the
	// value was found.
	DontCopyRawMessage

	// DontMatchCaseInsensitiveStructFields is a parsing flag used to prevent
	// matching fields in a case-insensitive way. This can prevent degrading
	// performance on case conversions, and can also act as a stricter decoding
	// mode.
	DontMatchCaseInsensitiveStructFields

	// Decode integers into *big.Int.
	// Takes precedence over UseNumber for integers.
	UseBigInt

	// Decode in-range integers to int64.
	// Takes precedence over UseNumber and UseBigInt for in-range integers.
	UseInt64

	// Decode in-range positive integers to uint64.
	// Takes precedence over UseNumber, UseBigInt, and UseInt64
	// for positive, in-range integers.
	UseUint64

	// ZeroCopy is a parsing flag that combines all the copy optimizations
	// available in the package.
	//
	// The zero-copy optimizations are better used in request-handler style
	// code where none of the values are retained after the handler returns.
	ZeroCopy = DontCopyString | DontCopyNumber | DontCopyRawMessage

	// validAsciiPrint is an internal flag indicating that the input contains
	// only valid ASCII print chars (0x20 <= c <= 0x7E). If the flag is unset,
	// it's unknown whether the input is valid ASCII print.
	validAsciiPrint ParseFlags = 1 << 28

	// noBackslach is an internal flag indicating that the input does not
	// contain a backslash. If the flag is unset, it's unknown whether the
	// input contains a backslash.
	noBackslash ParseFlags = 1 << 29

	// Bit offset where the kind of the json value is stored.
	//
	// See Kind in token.go for the enum.
	kindOffset ParseFlags = 16
)

// Kind represents the different kinds of value that exist in JSON.
type Kind uint

const (
	Undefined Kind = 0

	Null Kind = 1 // Null is not zero, so we keep zero for "undefined".

	Bool  Kind = 2 // Bit two is set to 1, means it's a boolean.
	False Kind = 2 // Bool + 0
	True  Kind = 3 // Bool + 1

	Num   Kind = 4 // Bit three is set to 1, means it's a number.
	Uint  Kind = 5 // Num + 1
	Int   Kind = 6 // Num + 2
	Float Kind = 7 // Num + 3

	String    Kind = 8 // Bit four is set to 1, means it's a string.
	Unescaped Kind = 9 // String + 1

	Array  Kind = 16 // Equivalent to Delim == '['
	Object Kind = 32 // Equivalent to Delim == '{'
)

// Class returns the class of k.
func (k Kind) Class() Kind { _ = "STUB: not implemented"; return *new(Kind) }

// Append acts like Marshal but appends the json representation to b instead of
// always reallocating a new slice.
func Append(b []byte, x any, flags AppendFlags) ([]byte, error) {
	_ = "STUB: not implemented"

	// Special case for nil values because it makes the rest of the code
	// simpler to assume that it won't be seeing nil pointers.
	return nil, nil
}

// Escape is a convenience helper to construct an escaped JSON string from s.
// The function escales HTML characters, for more control over the escape
// behavior and to write to a pre-allocated buffer, use AppendEscape.
func Escape(s string) []byte {
	_ = "STUB: not implemented"
	// +10 for extra escape characters, maybe not enough and the buffer will
	// be reallocated.
	return nil
}

// AppendEscape appends s to b with the string escaped as a JSON value.
// This will include the starting and ending quote characters, and the
// appropriate characters will be escaped correctly for JSON encoding.
func AppendEscape(b []byte, s string, flags AppendFlags) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Unescape is a convenience helper to unescape a JSON value.
// For more control over the unescape behavior and
// to write to a pre-allocated buffer, use AppendUnescape.
func Unescape(s []byte) []byte { _ = "STUB: not implemented"; return nil }

// AppendUnescape appends s to b with the string unescaped as a JSON value.
// This will remove starting and ending quote characters, and the
// appropriate characters will be escaped correctly as if JSON decoded.
// New space will be reallocated if more space is needed.
func AppendUnescape(b []byte, s []byte, flags ParseFlags) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Compact is documented at https://golang.org/pkg/encoding/json/#Compact
func Compact(dst *bytes.Buffer, src []byte) error { _ = "STUB: not implemented"; return nil }

// HTMLEscape is documented at https://golang.org/pkg/encoding/json/#HTMLEscape
func HTMLEscape(dst *bytes.Buffer, src []byte) { _ = "STUB: not implemented"; return }

// Indent is documented at https://golang.org/pkg/encoding/json/#Indent
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	_ = "STUB: not implemented"
	return nil
}

// Marshal is documented at https://golang.org/pkg/encoding/json/#Marshal
func Marshal(x any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalIndent is documented at https://golang.org/pkg/encoding/json/#MarshalIndent
func MarshalIndent(x any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal is documented at https://golang.org/pkg/encoding/json/#Unmarshal
func Unmarshal(b []byte, x any) error { _ = "STUB: not implemented"; return nil }

// The encoding/json package prioritizes reporting errors caused by
// unexpected trailing bytes over other issues; here we emulate this
// behavior by overriding the error.

// Parse behaves like Unmarshal but the caller can pass a set of flags to
// configure the parsing behavior.
func Parse(b []byte, x any, flags ParseFlags) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Valid is documented at https://golang.org/pkg/encoding/json/#Valid
func Valid(data []byte) bool { _ = "STUB: not implemented"; return false }

// Decoder is documented at https://golang.org/pkg/encoding/json/#Decoder
type Decoder struct {
	reader      io.Reader
	buffer      []byte
	remain      []byte
	inputOffset int64
	err         error
	flags       ParseFlags
}

// NewDecoder is documented at https://golang.org/pkg/encoding/json/#NewDecoder
func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// Buffered is documented at https://golang.org/pkg/encoding/json/#Decoder.Buffered
func (dec *Decoder) Buffered() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// Decode is documented at https://golang.org/pkg/encoding/json/#Decoder.Decode
func (dec *Decoder) Decode(v any) error { _ = "STUB: not implemented"; return nil }

const (
	minBufferSize = 32768
	minReadSize   = 4096
)

// readValue reads one JSON value from the buffer and returns its raw bytes. It
// is optimized for the "one JSON value per line" case.
func (dec *Decoder) readValue() (v []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Parsing of the next JSON value stopped at a position other
// than the end of the input buffer, which indicaates that a
// syntax error was encountered.

// DisallowUnknownFields is documented at https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields
func (dec *Decoder) DisallowUnknownFields() { _ = "STUB: not implemented"; return }

// UseNumber is documented at https://golang.org/pkg/encoding/json/#Decoder.UseNumber
func (dec *Decoder) UseNumber() { _ = "STUB: not implemented"; return }

// DontCopyString is an extension to the standard encoding/json package
// which instructs the decoder to not copy strings loaded from the json
// payloads when possible.
func (dec *Decoder) DontCopyString() { _ = "STUB: not implemented"; return }

// DontCopyNumber is an extension to the standard encoding/json package
// which instructs the decoder to not copy numbers loaded from the json
// payloads.
func (dec *Decoder) DontCopyNumber() { _ = "STUB: not implemented"; return }

// DontCopyRawMessage is an extension to the standard encoding/json package
// which instructs the decoder to not allocate RawMessage values in separate
// memory buffers (see the documentation of the DontcopyRawMessage flag for
// more detais).
func (dec *Decoder) DontCopyRawMessage() { _ = "STUB: not implemented"; return }

// DontMatchCaseInsensitiveStructFields is an extension to the standard
// encoding/json package which instructs the decoder to not match object fields
// against struct fields in a case-insensitive way, the field names have to
// match exactly to be decoded into the struct field values.
func (dec *Decoder) DontMatchCaseInsensitiveStructFields() { _ = "STUB: not implemented"; return }

// ZeroCopy is an extension to the standard encoding/json package which enables
// all the copy optimizations of the decoder.
func (dec *Decoder) ZeroCopy() { _ = "STUB: not implemented"; return }

// InputOffset returns the input stream byte offset of the current decoder position.
// The offset gives the location of the end of the most recently returned token
// and the beginning of the next token.
func (dec *Decoder) InputOffset() int64 { _ = "STUB: not implemented"; return 0 }

// Encoder is documented at https://golang.org/pkg/encoding/json/#Encoder
type Encoder struct {
	writer io.Writer
	prefix string
	indent string
	buffer *bytes.Buffer
	err    error
	flags  AppendFlags
}

// NewEncoder is documented at https://golang.org/pkg/encoding/json/#NewEncoder
func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

// Encode is documented at https://golang.org/pkg/encoding/json/#Encoder.Encode
func (enc *Encoder) Encode(v any) error { _ = "STUB: not implemented"; return nil }

// SetEscapeHTML is documented at https://golang.org/pkg/encoding/json/#Encoder.SetEscapeHTML
func (enc *Encoder) SetEscapeHTML(on bool) { _ = "STUB: not implemented"; return }

// SetIndent is documented at https://golang.org/pkg/encoding/json/#Encoder.SetIndent
func (enc *Encoder) SetIndent(prefix, indent string) { _ = "STUB: not implemented"; return }

// SetSortMapKeys is an extension to the standard encoding/json package which
// allows the program to toggle sorting of map keys on and off.
func (enc *Encoder) SetSortMapKeys(on bool) { _ = "STUB: not implemented"; return }

// SetTrustRawMessage skips value checking when encoding a raw json message. It should only
// be used if the values are known to be valid json, e.g. because they were originally created
// by json.Unmarshal.
func (enc *Encoder) SetTrustRawMessage(on bool) { _ = "STUB: not implemented"; return }

// SetAppendNewline is an extension to the standard encoding/json package which
// allows the program to toggle the addition of a newline in Encode on or off.
func (enc *Encoder) SetAppendNewline(on bool) { _ = "STUB: not implemented"; return }

var encoderBufferPool = sync.Pool{
	New: func() any { return &encoderBuffer{data: make([]byte, 0, 4096)} },
}

type encoderBuffer struct{ data []byte }
