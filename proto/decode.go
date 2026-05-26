package proto

import (
	"errors"
	"unsafe"
)

// DecodeTag reverses the encoding applied by EncodeTag.
func DecodeTag(tag uint64) (FieldNumber, WireType) {
	_ = "STUB: not implemented"
	return *new(FieldNumber), *new(WireType)
}

// DecodeZigZag reverses the encoding applied by EncodeZigZag.
func DecodeZigZag(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

func decodeZigZag64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

func decodeZigZag32(v uint32) int32 { _ = "STUB: not implemented"; return 0 }

type decodeFunc = func([]byte, unsafe.Pointer, flags) (int, error)

var errVarintOverflow = errors.New("varint overflowed 64 bits integer")

func decodeVarint(b []byte) (uint64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

// Fast-path for decoding the common case of varints that fit on a
// single byte.
//
// This path is ~60% faster than calling binary.Uvarint.

func decodeVarintZigZag(b []byte) (int64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func decodeLE32(b []byte) (uint32, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func decodeLE64(b []byte) (uint64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func decodeTag(b []byte) (f fieldNumber, t wireType, n int, err error) {
	_ = "STUB: not implemented"
	return *new(fieldNumber), *new(wireType), 0, nil
}

func decodeVarlen(b []byte) ([]byte, int, error) { _ = "STUB: not implemented"; return nil, 0, nil }
