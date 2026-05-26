package proto

import (
	"reflect"

	. "github.com/segmentio/encoding/internal/runtime_reflect"
)

type repeatedField struct {
	codec       *codec
	fieldNumber fieldNumber
	wireType    wireType
	embedded    bool
}

func sliceCodecOf(t reflect.Type, f structField, seen map[reflect.Type]*codec) *codec {
	_ = "STUB: not implemented"
	return nil
}

func sliceSizeFuncOf(t reflect.Type, r *repeatedField) sizeFunc {
	_ = "STUB: not implemented"
	return *new(sizeFunc)
}

func sliceEncodeFuncOf(t reflect.Type, r *repeatedField) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func sliceDecodeFuncOf(t reflect.Type, r *repeatedField) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func alignedSize(t reflect.Type) uintptr { _ = "STUB: not implemented"; return 0 }

func align(align, size uintptr) uintptr { _ = "STUB: not implemented"; return 0 }

func growSlice(t reflect.Type, s *Slice) Slice { _ = "STUB: not implemented"; return *new(Slice) }
