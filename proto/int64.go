package proto

import "unsafe"

var int64Codec = codec{
	wire:   varint,
	size:   sizeOfInt64,
	encode: encodeInt64,
	decode: decodeInt64,
}

func sizeOfInt64(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeInt64(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeInt64(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
