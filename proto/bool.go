package proto

import (
	"unsafe"
)

var boolCodec = codec{
	wire:   varint,
	size:   sizeOfBool,
	encode: encodeBool,
	decode: decodeBool,
}

func sizeOfBool(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeBool(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeBool(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
