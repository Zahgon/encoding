package proto

import (
	"reflect"

	. "github.com/segmentio/encoding/internal/runtime_reflect"
)

const (
	zeroSize = 1 // sizeOfVarint(0)
)

type mapField struct {
	number   uint16
	keyFlags uint8
	valFlags uint8
	keyCodec *codec
	valCodec *codec
}

func mapCodecOf(t reflect.Type, f *mapField, seen map[reflect.Type]*codec) *codec {
	_ = "STUB: not implemented"
	return nil
}

func mapSizeFuncOf(t reflect.Type, f *mapField) sizeFunc {
	_ = "STUB: not implemented"
	return *new(sizeFunc)
}

func mapEncodeFuncOf(t reflect.Type, f *mapField) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func mapDecodeFuncOf(t reflect.Type, f *mapField, seen map[reflect.Type]*codec) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}
