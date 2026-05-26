package thrift

import (
	"reflect"
)

type Message struct {
	Type  MessageType
	Name  string
	SeqID int32
}

type MessageType int8

const (
	Call MessageType = iota
	Reply
	Exception
	Oneway
)

func (m MessageType) String() string { _ = "STUB: not implemented"; return "" }

type Field struct {
	ID    int16
	Type  Type
	Delta bool // whether the field id is a delta
}

func (f Field) String() string { _ = "STUB: not implemented"; return "" }

type Type int8

const (
	STOP Type = iota
	TRUE
	FALSE
	I8
	I16
	I32
	I64
	DOUBLE
	BINARY
	LIST
	SET
	MAP
	STRUCT
	BOOL = FALSE
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func (t Type) GoString() string { _ = "STUB: not implemented"; return "" }

type List struct {
	Size int32
	Type Type
}

func (l List) String() string { _ = "STUB: not implemented"; return "" }

type Set List

func (s Set) String() string { _ = "STUB: not implemented"; return "" }

type Map struct {
	Size  int32
	Key   Type
	Value Type
}

func (m Map) String() string { _ = "STUB: not implemented"; return "" }

func TypeOf(t reflect.Type) Type { _ = "STUB: not implemented"; return *new(Type) }

// []byte
