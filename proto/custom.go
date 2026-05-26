package proto

import (
	"reflect"
)

func customCodecOf(t reflect.Type) *codec { _ = "STUB: not implemented"; return nil }

func customSizeFuncOf(t reflect.Type) sizeFunc { _ = "STUB: not implemented"; return *new(sizeFunc) }

func customEncodeFuncOf(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func customDecodeFuncOf(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}
