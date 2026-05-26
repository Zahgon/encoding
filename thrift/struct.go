package thrift

import (
	"reflect"
)

type flags int16

const (
	enum     flags = 1 << 0
	union    flags = 1 << 1
	required flags = 1 << 2
	optional flags = 1 << 3
	strict   flags = 1 << 4

	featuresBitOffset  = 8
	useDeltaEncoding   = flags(UseDeltaEncoding) << featuresBitOffset
	coalesceBoolFields = flags(CoalesceBoolFields) << featuresBitOffset

	structFlags   flags = enum | union | required | optional
	encodeFlags   flags = strict | protocolFlags
	decodeFlags   flags = strict | protocolFlags
	protocolFlags flags = useDeltaEncoding | coalesceBoolFields
)

func (f flags) have(x flags) bool { _ = "STUB: not implemented"; return false }

func (f flags) only(x flags) flags { _ = "STUB: not implemented"; return *new(flags) }

func (f flags) with(x flags) flags { _ = "STUB: not implemented"; return *new(flags) }

func (f flags) without(x flags) flags { _ = "STUB: not implemented"; return *new(flags) }

type structField struct {
	typ   reflect.Type
	index []int
	id    int16
	flags flags
}

func forEachStructField(t reflect.Type, index []int, do func(structField)) {
	_ = "STUB: not implemented"
	return
}

// unexported
