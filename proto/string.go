package proto

import (
	"unsafe"
)

var stringCodec = codec{
	wire:   varlen,
	size:   sizeOfString,
	encode: encodeString,
	decode: decodeString,
}

func sizeOfString(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeString(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeString(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
