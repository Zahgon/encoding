package proto

import (
	"reflect"
)

// Message is an interface implemented by types that supported being encoded to
// and decoded from protobuf.
type Message interface {
	// Size is the size of the protobuf representation (in bytes).
	Size() int

	// Marshal writes the message to the byte slice passed as argument.
	Marshal([]byte) error

	// Unmarshal reads the message from the byte slice passed as argument.
	Unmarshal([]byte) error
}

// RawMessage represents a raw protobuf-encoded message.
type RawMessage []byte

// Size satisfies the Message interface.
func (m RawMessage) Size() int {
	_ = "STUB: not implemented"

	// Marshal satisfies the Message interface.
	return 0
}

func (m RawMessage) Marshal(b []byte) error { _ = "STUB: not implemented"; return nil }

// Unmarshal satisfies the Message interface.
func (m *RawMessage) Unmarshal(b []byte) error { _ = "STUB: not implemented"; return nil }

// Rewrite satisfies the Rewriter interface.
func (m RawMessage) Rewrite(out, _ []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// FieldNumber represents a protobuf field number.
}

type FieldNumber uint

func (f FieldNumber) Bool(v bool) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

func (f FieldNumber) Int(v int) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

func (f FieldNumber) Int32(v int32) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

func (f FieldNumber) Int64(v int64) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

func (f FieldNumber) Uint(v uint) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

func (f FieldNumber) Uint32(v uint32) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) Uint64(v uint64) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) Fixed32(v uint32) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) Fixed64(v uint64) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) Float32(v float32) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) Float64(v float64) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) String(v string) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func (f FieldNumber) Bytes(v []byte) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

// Value constructs a RawMessage for field number f from v.
func (f FieldNumber) Value(v any) RawMessage { _ = "STUB: not implemented"; return *new(RawMessage) }

// The WireType enumeration represents the different protobuf wire types.
type WireType uint

const (
	Varint  WireType = 0
	Fixed64 WireType = 1
	Varlen  WireType = 2
	Fixed32 WireType = 5
	// Wire types 3 and 4 were used for StartGroup and EndGroup, but are
	// deprecated so we don't expose them here.
	//
	// https://developers.google.com/protocol-buffers/docs/encoding#structure
)

func (wt WireType) String() string { _ = "STUB: not implemented"; return "" }

func Append(m RawMessage, f FieldNumber, t WireType, v []byte) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func AppendVarint(m RawMessage, f FieldNumber, v uint64) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func AppendVarlen(m RawMessage, f FieldNumber, v []byte) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func AppendFixed32(m RawMessage, f FieldNumber, v uint32) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func AppendFixed64(m RawMessage, f FieldNumber, v uint64) RawMessage {
	_ = "STUB: not implemented"
	return *new(RawMessage)
}

func Parse(m []byte) (FieldNumber, WireType, RawValue, RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(FieldNumber), *new(WireType), *new(RawValue), *new(RawMessage), nil
}

// length

// Scan calls fn for each protobuf field in the message b.
//
// The iteration stops when all fields have been scanned, fn returns false, or
// an error is seen.
func Scan(b []byte, fn func(FieldNumber, WireType, RawValue) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// RawValue represents a single protobuf value.
//
// RawValue instances are returned by Parse and share the backing array of the
// RawMessage that they were decoded from.
type RawValue []byte

// Varint decodes v as a varint.
//
// The content of v will always be a valid varint if v was returned by a call to
// Parse and the associated wire type was Varint. In other cases, the behavior
// of Varint is undefined.
func (v RawValue) Varint() uint64 { _ = "STUB: not implemented"; return 0 }

// Fixed32 decodes v as a fixed32.
//
// The content of v will always be a valid fixed32 if v was returned by a call
// to Parse and the associated wire type was Fixed32. In other cases, the
// behavior of Fixed32 is undefined.
func (v RawValue) Fixed32() uint32 { _ = "STUB: not implemented"; return 0 }

// Fixed64 decodes v as a fixed64.
//
// The content of v will always be a valid fixed64 if v was returned by a call
// to Parse and the associated wire type was Fixed64. In other cases, the
// behavior of Fixed64 is undefined.
func (v RawValue) Fixed64() uint64 { _ = "STUB: not implemented"; return 0 }

var (
	_ Message  = &RawMessage{}
	_ Rewriter = RawMessage{}
)

func messageCodecOf(t reflect.Type) *codec { _ = "STUB: not implemented"; return nil }

func messageSizeFuncOf(t reflect.Type) sizeFunc { _ = "STUB: not implemented"; return *new(sizeFunc) }

func messageEncodeFuncOf(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func messageDecodeFuncOf(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}
