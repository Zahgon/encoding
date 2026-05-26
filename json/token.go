package json

import (
	"sync"
)

// Tokenizer is an iterator-style type which can be used to progressively parse
// through a json input.
//
// Tokenizing json is useful to build highly efficient parsing operations, for
// example when doing tranformations on-the-fly where as the program reads the
// input and produces the transformed json to an output buffer.
//
// Here is a common pattern to use a tokenizer:
//
//	for t := json.NewTokenizer(b); t.Next(); {
//		switch k := t.Kind(); k.Class() {
//		case json.Null:
//			...
//		case json.Bool:
//			...
//		case json.Num:
//			...
//		case json.String:
//			...
//		case json.Array:
//			...
//		case json.Object:
//			...
//		}
//	}
type Tokenizer struct {
	// When the tokenizer is positioned on a json delimiter this field is not
	// zero. In this case the possible values are '{', '}', '[', ']', ':', and
	// ','.
	Delim Delim

	// This field contains the raw json token that the tokenizer is pointing at.
	// When Delim is not zero, this field is a single-element byte slice
	// continaing the delimiter value. Otherwise, this field holds values like
	// null, true, false, numbers, or quoted strings.
	Value RawValue

	// When the tokenizer has encountered invalid content this field is not nil.
	Err error

	// When the value is in an array or an object, this field contains the depth
	// at which it was found.
	Depth int

	// When the value is in an array or an object, this field contains the
	// position at which it was found.
	Index int

	// This field is true when the value is the key of an object.
	IsKey bool

	// Tells whether the next value read from the tokenizer is a key.
	isKey bool

	// json input for the tokenizer, pointing at data right after the last token
	// that was parsed.
	json []byte

	// Stack used to track entering and leaving arrays, objects, and keys.
	stack *stack

	// Decoder used for parsing.
	decoder
}

// NewTokenizer constructs a new Tokenizer which reads its json input from b.
func NewTokenizer(b []byte) *Tokenizer { _ = "STUB: not implemented"; return nil }

// Reset erases the state of t and re-initializes it with the json input from b.
func (t *Tokenizer) Reset(b []byte) { _ = "STUB: not implemented"; return }

// This code is similar to:
//
//	*t = Tokenizer{json: b}
//
// However, it does not compile down to an invocation of duff-copy.

// Next returns a new tokenizer pointing at the next token, or the zero-value of
// Tokenizer if the end of the json input has been reached.
//
// If the tokenizer encounters malformed json while reading the input the method
// sets t.Err to an error describing the issue, and returns false. Once an error
// has been encountered, the tokenizer will always fail until its input is
// cleared by a call to its Reset method.
func (t *Tokenizer) Next() bool { _ = "STUB: not implemented"; return false }

// Inlined code of the skipSpaces function, this give a ~15% speed boost.

func (t *Tokenizer) depth() int { _ = "STUB: not implemented"; return 0 }

func (t *Tokenizer) index() int { _ = "STUB: not implemented"; return 0 }

func (t *Tokenizer) push(typ scope) { _ = "STUB: not implemented"; return }

func (t *Tokenizer) pop(expect scope) error { _ = "STUB: not implemented"; return nil }

// Kind returns the kind of the value that the tokenizer is currently positioned
// on.
func (t *Tokenizer) Kind() Kind {
	_ = "STUB: not implemented"
	return *

	// Bool returns a bool containing the value of the json boolean that the
	// tokenizer is currently pointing at.
	//
	// This method must only be called after checking the kind of the token via a
	// call to Kind.
	//
	// If the tokenizer is not positioned on a boolean, the behavior is undefined.
	new(Kind)
}

func (t *Tokenizer) Bool() bool { _ = "STUB: not implemented"; return false }

// Int returns a byte slice containing the value of the json number that the
// tokenizer is currently pointing at.
//
// This method must only be called after checking the kind of the token via a
// call to Kind.
//
// If the tokenizer is not positioned on an integer, the behavior is undefined.
func (t *Tokenizer) Int() int64 { _ = "STUB: not implemented"; return 0 }

// Uint returns a byte slice containing the value of the json number that the
// tokenizer is currently pointing at.
//
// This method must only be called after checking the kind of the token via a
// call to Kind.
//
// If the tokenizer is not positioned on a positive integer, the behavior is
// undefined.
func (t *Tokenizer) Uint() uint64 { _ = "STUB: not implemented"; return 0 }

// Float returns a byte slice containing the value of the json number that the
// tokenizer is currently pointing at.
//
// This method must only be called after checking the kind of the token via a
// call to Kind.
//
// If the tokenizer is not positioned on a number, the behavior is undefined.
func (t *Tokenizer) Float() float64 { _ = "STUB: not implemented"; return 0 }

// String returns a byte slice containing the value of the json string that the
// tokenizer is currently pointing at.
//
// This method must only be called after checking the kind of the token via a
// call to Kind.
//
// When possible, the returned byte slice references the backing array of the
// tokenizer. A new slice is only allocated if the tokenizer needed to unescape
// the json string.
//
// If the tokenizer is not positioned on a string, the behavior is undefined.
func (t *Tokenizer) String() []byte { _ = "STUB: not implemented"; return nil }

// unquote

// Remaining returns the number of bytes left to parse.
//
// The position of the tokenizer's current Value within the original byte slice
// can be calculated like so:
//
//	end := len(b) - tok.Remaining()
//	start := end - len(tok.Value)
//
// And slicing b[start:end] will yield the tokenizer's current Value.
func (t *Tokenizer) Remaining() int {
	_ = "STUB: not implemented"

	// RawValue represents a raw json value, it is intended to carry null, true,
	// false, number, and string values only.
	return 0
}

type RawValue []byte

// String returns true if v contains a string value.
func (v RawValue) String() bool { _ = "STUB: not implemented"; return false }

// Null returns true if v contains a null value.
func (v RawValue) Null() bool { _ = "STUB: not implemented"; return false }

// True returns true if v contains a true value.
func (v RawValue) True() bool { _ = "STUB: not implemented"; return false }

// False returns true if v contains a false value.
func (v RawValue) False() bool { _ = "STUB: not implemented"; return false }

// Number returns true if v contains a number value.
func (v RawValue) Number() bool { _ = "STUB: not implemented"; return false }

// AppendUnquote writes the unquoted version of the string value in v into b.
func (v RawValue) AppendUnquote(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// Unquote returns the unquoted version of the string value in v.
func (v RawValue) Unquote() []byte { _ = "STUB: not implemented"; return nil }

type scope int

const (
	inArray scope = iota
	inObject
)

type state struct {
	typ scope
	len int
}

type stack struct {
	state []state
}

func (s *stack) push(typ scope) { _ = "STUB: not implemented"; return }

func (s *stack) pop(expect scope) bool { _ = "STUB: not implemented"; return false }

func (s *stack) is(typ scope) bool { _ = "STUB: not implemented"; return false }

func (s *stack) depth() int { _ = "STUB: not implemented"; return 0 }

func (s *stack) index() int { _ = "STUB: not implemented"; return 0 }

func acquireStack() *stack { _ = "STUB: not implemented"; return nil }

func releaseStack(s *stack) { _ = "STUB: not implemented"; return }

var stackPool sync.Pool // *stack
