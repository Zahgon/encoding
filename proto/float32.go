package proto

import (
	"unsafe"
)

var float32Codec = codec{
	wire:   fixed32,
	size:   sizeOfFloat32,
	encode: encodeFloat32,
	decode: decodeFloat32,
}

func sizeOfFloat32(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeFloat32(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeFloat32(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
