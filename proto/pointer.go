package proto

import (
	"reflect"
)

func pointerCodecOf(t reflect.Type, seen map[reflect.Type]*codec) *codec {
	_ = "STUB: not implemented"
	return nil
}

func pointerSizeFuncOf(t reflect.Type, c *codec) sizeFunc {
	_ = "STUB: not implemented"
	return *new(sizeFunc)
}

func pointerEncodeFuncOf(t reflect.Type, c *codec) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func pointerDecodeFuncOf(t reflect.Type, c *codec) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}
