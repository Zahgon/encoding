package proto

import (
	"reflect"
	"sync/atomic"
	"unsafe"
)

func Size(v any) int { _ = "STUB: not implemented"; return 0 }

func Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MarshalTo(b []byte, v any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func Unmarshal(b []byte, v any) error {
	_ = "STUB: not implemented"

	// An empty input is a valid protobuf message with all fields set to the
	// zero-value.
	return nil
}

// Unmarshal must be passed a pointer

type flags uintptr

const (
	noflags  flags = 0
	inline   flags = 1 << 0
	wantzero flags = 1 << 1
	// Shared with structField.flags in struct.go:
	// zigzag flags = 1 << 2
	toplevel flags = 1 << 3
)

func (f flags) has(x flags) bool { _ = "STUB: not implemented"; return false }

func (f flags) with(x flags) flags { _ = "STUB: not implemented"; return *new(flags) }

func (f flags) without(x flags) flags { _ = "STUB: not implemented"; return *new(flags) }

func (f flags) uint64(i int64) uint64 { _ = "STUB: not implemented"; return 0 }

func (f flags) int64(u uint64) int64 { _ = "STUB: not implemented"; return 0 }

type iface struct {
	typ unsafe.Pointer
	ptr unsafe.Pointer
}

func inspect(v any) (reflect.Type, unsafe.Pointer) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), *new(unsafe.Pointer)
}

func pointer(v any) unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func inlined(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

type fieldNumber uint

type wireType uint

const (
	varint  wireType = 0
	fixed64 wireType = 1
	varlen  wireType = 2
	fixed32 wireType = 5
)

func (wt wireType) String() string { _ = "STUB: not implemented"; return "" }

type codec struct {
	wire   wireType
	size   sizeFunc
	encode encodeFunc
	decode decodeFunc
}

var codecCache atomic.Value // map[unsafe.Pointer]*codec

func loadCachedCodec(t reflect.Type) (*codec, map[unsafe.Pointer]*codec) {
	_ = "STUB: not implemented"
	return nil, nil
}

func storeCachedCodec(newCache map[unsafe.Pointer]*codec) { _ = "STUB: not implemented"; return }

func cachedCodecOf(t reflect.Type) *codec { _ = "STUB: not implemented"; return nil }

func codecOf(t reflect.Type, seen map[reflect.Type]*codec) *codec {
	_ = "STUB: not implemented"
	return nil
}

// backward compatibility with gogoproto custom types.
type customMessage interface {
	Size() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
}

type protoMessage interface {
	ProtoMessage()
}

var (
	messageType       = reflect.TypeOf((*Message)(nil)).Elem()
	customMessageType = reflect.TypeOf((*customMessage)(nil)).Elem()
	protoMessageType  = reflect.TypeOf((*protoMessage)(nil)).Elem()
)

func implements(t, iface reflect.Type) bool { _ = "STUB: not implemented"; return false }
