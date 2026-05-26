package proto

import (
	"unsafe"
)

var intCodec = codec{
	wire:   varint,
	size:   sizeOfInt,
	encode: encodeInt,
	decode: decodeInt,
}

func sizeOfInt(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeInt(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeInt(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
