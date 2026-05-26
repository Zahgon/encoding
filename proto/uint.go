package proto

import "unsafe"

var uintCodec = codec{
	wire:   varint,
	size:   sizeOfUint,
	encode: encodeUint,
	decode: decodeUint,
}

func sizeOfUint(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeUint(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeUint(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
