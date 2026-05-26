package proto

import "unsafe"

var uint64Codec = codec{
	wire:   varint,
	size:   sizeOfUint64,
	encode: encodeUint64,
	decode: decodeUint64,
}

func sizeOfUint64(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeUint64(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeUint64(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var fixed64Codec = codec{
	wire:   fixed64,
	size:   sizeOfFixed64,
	encode: encodeFixed64,
	decode: decodeFixed64,
}

func sizeOfFixed64(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeFixed64(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeFixed64(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
