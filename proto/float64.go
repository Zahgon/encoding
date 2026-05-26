package proto

import (
	"unsafe"
)

var float64Codec = codec{
	wire:   fixed64,
	size:   sizeOfFloat64,
	encode: encodeFloat64,
	decode: decodeFloat64,
}

func sizeOfFloat64(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeFloat64(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeFloat64(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
