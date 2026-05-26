package proto

import (
	"unsafe"
)

var int32Codec = codec{
	wire:   varint,
	size:   sizeOfInt32,
	encode: encodeInt32,
	decode: decodeInt32,
}

func sizeOfInt32(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeInt32(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeInt32(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
