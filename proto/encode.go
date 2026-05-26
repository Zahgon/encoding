package proto

import (
	"unsafe"
)

// EncodeTag encodes a pair of field number and wire type into a protobuf tag.
func EncodeTag(f FieldNumber, t WireType) uint64 { _ = "STUB: not implemented"; return 0 }

// EncodeZigZag returns v as a zig-zag encoded value.
func EncodeZigZag(v int64) uint64 { _ = "STUB: not implemented"; return 0 }

func encodeZigZag64(v int64) uint64 { _ = "STUB: not implemented"; return 0 }

func encodeZigZag32(v int32) uint32 { _ = "STUB: not implemented"; return 0 }

type encodeFunc = func([]byte, unsafe.Pointer, flags) (int, error)

func encodeVarint(b []byte, v uint64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func encodeVarintZigZag(b []byte, v int64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func encodeLE32(b []byte, v uint32) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func encodeLE64(b []byte, v uint64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func encodeTag(b []byte, f fieldNumber, t wireType) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
