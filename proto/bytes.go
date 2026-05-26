package proto

import (
	"reflect"
	"unsafe"
)

var bytesCodec = codec{
	wire:   varlen,
	size:   sizeOfBytes,
	encode: encodeBytes,
	decode: decodeBytes,
}

func sizeOfBytes(p unsafe.Pointer, flags flags) int { _ = "STUB: not implemented"; return 0 }

func encodeBytes(b []byte, p unsafe.Pointer, flags flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func decodeBytes(b []byte, p unsafe.Pointer, _ flags) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func makeBytes(p unsafe.Pointer, n int) []byte { _ = "STUB: not implemented"; return nil }

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// isZeroBytes is an optimized version of this loop:
//
//	for i := range b {
//		if b[i] != 0 {
//			return false
//		}
//	}
//	return true
//
// This implementation significantly reduces the CPU footprint of checking for
// slices to be zero, especially when the length increases (these cases should
// be rare tho).
//
// name            old time/op  new time/op  delta
// IsZeroBytes0    1.78ns ± 1%  2.29ns ± 4%  +28.65%  (p=0.000 n=8+10)
// IsZeroBytes4    3.17ns ± 3%  2.37ns ± 3%  -25.21%  (p=0.000 n=10+10)
// IsZeroBytes7    3.97ns ± 4%  3.26ns ± 3%  -18.02%  (p=0.000 n=10+10)
// IsZeroBytes64K  14.8µs ± 3%   1.9µs ± 3%  -87.34%  (p=0.000 n=10+10)
func isZeroBytes(b []byte) bool { _ = "STUB: not implemented"; return false }

func bto32(b []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func bto16(b []byte) uint16 { _ = "STUB: not implemented"; return 0 }

func isZeroUint64(b []uint64) bool { _ = "STUB: not implemented"; return false }

func byteArrayCodecOf(t reflect.Type, seen map[reflect.Type]*codec) *codec {
	_ = "STUB: not implemented"
	return nil
}

func byteArraySizeFuncOf(n int) sizeFunc { _ = "STUB: not implemented"; return *new(sizeFunc) }

func byteArrayEncodeFuncOf(n int) encodeFunc { _ = "STUB: not implemented"; return *new(encodeFunc) }

func byteArrayDecodeFuncOf(n int) decodeFunc { _ = "STUB: not implemented"; return *new(decodeFunc) }
