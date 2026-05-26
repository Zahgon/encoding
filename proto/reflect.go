package proto

import (
	"reflect"
	"sync"
	"sync/atomic"
)

// Kind is an enumeration representing the various data types supported by the
// protobuf language.
type Kind int

const (
	Bool Kind = iota
	Int32
	Int64
	Sint32
	Sint64
	Uint32
	Uint64
	Fix32
	Fix64
	Sfix32
	Sfix64
	Float
	Double
	String
	Bytes
	Map
	Struct
)

// Type is an interface similar to reflect.Type. Values implementing this
// interface represent high level protobuf types.
//
// Type values are safe to use concurrently from multiple goroutines.
//
// Types are comparable value.
type Type interface {
	// Returns a human-readable representation of the type.
	String() string

	// Returns the name of the type.
	Name() string

	// Kind returns the kind of protobuf values that are represented.
	Kind() Kind

	// When the Type represents a protobuf map, calling this method returns the
	// type of the map keys.
	//
	// If the Type is not a map type, the method panics.
	Key() Type

	// When the Type represents a protobuf map, calling this method returns the
	// type of the map values.
	//
	// If the Type is not a map type, the method panics.
	Elem() Type

	// Returns the protobuf wire type for the Type it is called on.
	WireType() WireType

	// Returns the number of fields in the protobuf message.
	//
	// If the Type does not represent a struct type, the method returns zero.
	NumField() int

	// Returns the Field at the given in Type.
	//
	// If the Type does not represent a struct type, the method panics.
	Field(int) Field

	// Returns the Field with the given name in Type.
	//
	// If the Type does not represent a struct type, or if the field does not
	// exist, the method panics.
	FieldByName(string) Field

	// Returns the Field with the given number in Type.
	//
	// If the Type does not represent a struct type, or if the field does not
	// exist, the method panics.
	FieldByNumber(FieldNumber) Field

	// For unsigned types, convert to their zig-zag form.
	//
	// The method uses the following table to perform the conversion:
	//
	//  base    | zig-zag
	//	--------+---------
	//	int32   | sint32
	//	int64   | sint64
	//	uint32  | sint32
	//	uint64  | sint64
	//	fixed32 | sfixed32
	//	fixed64 | sfixed64
	//
	// If the type cannot be converted to a zig-zag type, the method panics.
	ZigZag() Type
}

// TypeOf returns the protobuf type used to represent a go type.
//
// The function uses the following table to map Go types to Protobuf:
//
//	Go      | Protobuf
//	--------+---------
//	bool    | bool
//	int     | int64
//	int32   | int32
//	int64   | int64
//	uint    | uint64
//	uint32  | uint32
//	uint64  | uint64
//	float32 | float
//	float64 | double
//	string  | string
//	[]byte  | bytes
//	map     | map
//	struct  | message
//
// Pointer types are also supported and automatically dereferenced.
func TypeOf(t reflect.Type) Type { _ = "STUB: not implemented"; return *new(Type) }

func typeOf(t reflect.Type, seen map[reflect.Type]Type) Type {
	_ = "STUB: not implemented"
	return *new(Type)
}

var (
	typesMutex sync.Mutex
	typesCache atomic.Value // map[reflect.Type]Type{}
)

type Field struct {
	Index    int
	Number   FieldNumber
	Name     string
	Type     Type
	Repeated bool
}

type primitiveType struct {
	name   string
	kind   Kind
	wire   WireType
	zigzag Kind
}

func (t *primitiveType) String() string { _ = "STUB: not implemented"; return "" }

func (t *primitiveType) Name() string { _ = "STUB: not implemented"; return "" }

func (t *primitiveType) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (t *primitiveType) Key() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *primitiveType) Elem() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *primitiveType) WireType() WireType { _ = "STUB: not implemented"; return *new(WireType) }

func (t *primitiveType) NumField() int { _ = "STUB: not implemented"; return 0 }

func (t *primitiveType) Field(int) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *primitiveType) FieldByName(string) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *primitiveType) FieldByNumber(FieldNumber) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func (t *primitiveType) ZigZag() Type { _ = "STUB: not implemented"; return *new(Type) }

var primitiveTypes = [...]primitiveType{
	{name: "bool", kind: Bool, wire: Varint},
	{name: "int32", kind: Int32, wire: Varint, zigzag: Sint32},
	{name: "int64", kind: Int64, wire: Varint, zigzag: Sint64},
	{name: "sint32", kind: Sint32, wire: Varint},
	{name: "sint64", kind: Sint64, wire: Varint},
	{name: "uint32", kind: Uint32, wire: Varint, zigzag: Sint32},
	{name: "uint64", kind: Uint64, wire: Varint, zigzag: Sint64},
	{name: "fixed32", kind: Fix32, wire: Fixed32, zigzag: Sfix32},
	{name: "fixed64", kind: Fix64, wire: Fixed64, zigzag: Sfix64},
	{name: "sfixed32", kind: Sfix32, wire: Fixed32},
	{name: "sfixed64", kind: Sfix64, wire: Fixed64},
	{name: "float", kind: Float, wire: Fixed32},
	{name: "double", kind: Double, wire: Fixed64},
	{name: "string", kind: String, wire: Varlen},
	{name: "bytes", kind: Bytes, wire: Varlen},
}

func mapTypeOf(t reflect.Type, seen map[reflect.Type]Type) *mapType {
	_ = "STUB: not implemented"
	return nil
}

type mapType struct {
	key  Type
	elem Type
}

func (t *mapType) String() string { _ = "STUB: not implemented"; return "" }

func (t *mapType) Name() string { _ = "STUB: not implemented"; return "" }

func (t *mapType) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (t *mapType) Key() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *mapType) Elem() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *mapType) WireType() WireType { _ = "STUB: not implemented"; return *new(WireType) }

func (t *mapType) NumField() int { _ = "STUB: not implemented"; return 0 }

func (t *mapType) Field(int) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *mapType) FieldByName(string) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *mapType) FieldByNumber(FieldNumber) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *mapType) ZigZag() Type { _ = "STUB: not implemented"; return *new(Type) }

func structTypeOf(t reflect.Type, seen map[reflect.Type]Type) *structType {
	_ = "STUB: not implemented"
	return nil
}

// unexported

// for typeOf

// Because maps are represented as repeated varlen fields on the
// wire, the generated protobuf code sets the `rep` attribute on
// the struct fields.

type structType struct {
	name           string
	fields         []Field
	fieldsByName   map[string]int
	fieldsByNumber map[FieldNumber]int
}

func (t *structType) String() string { _ = "STUB: not implemented"; return "" }

func (t *structType) Name() string { _ = "STUB: not implemented"; return "" }

func (t *structType) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (t *structType) Key() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *structType) Elem() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *structType) WireType() WireType { _ = "STUB: not implemented"; return *new(WireType) }

func (t *structType) NumField() int { _ = "STUB: not implemented"; return 0 }

func (t *structType) Field(index int) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *structType) FieldByName(name string) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *structType) FieldByNumber(number FieldNumber) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func (t *structType) ZigZag() Type { _ = "STUB: not implemented"; return *new(Type) }

type structTag struct {
	name        string
	enum        string
	json        string
	version     int
	wireType    WireType
	fieldNumber FieldNumber
	extensions  map[string]string
	repeated    bool
	zigzag      bool
}

func parseStructTag(tag string) (structTag, error) {
	_ = "STUB: not implemented"
	return *new(structTag), nil
}

// not sure what this is for

func splitFields(s string) []string { _ = "STUB: not implemented"; return nil }

func splitNameValue(s string) (name, value string) { _ = "STUB: not implemented"; return "", "" }

type opaqueMessageType struct{}

func (t *opaqueMessageType) String() string { _ = "STUB: not implemented"; return "" }

func (t *opaqueMessageType) Name() string { _ = "STUB: not implemented"; return "" }

func (t *opaqueMessageType) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (t *opaqueMessageType) Key() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *opaqueMessageType) Elem() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *opaqueMessageType) WireType() WireType { _ = "STUB: not implemented"; return *new(WireType) }

func (t *opaqueMessageType) NumField() int { _ = "STUB: not implemented"; return 0 }

func (t *opaqueMessageType) Field(int) Field { _ = "STUB: not implemented"; return *new(Field) }

func (t *opaqueMessageType) FieldByName(string) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func (t *opaqueMessageType) FieldByNumber(FieldNumber) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func (t *opaqueMessageType) ZigZag() Type { _ = "STUB: not implemented"; return *new(Type) }
