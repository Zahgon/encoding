package proto

import (
	"github.com/segmentio/encoding/json"
)

// Rewriter is an interface implemented by types that support rewriting protobuf
// messages.
type Rewriter interface {
	// The function is expected to append the new content to the byte slice
	// passed as argument. If it wasn't able to perform the rewrite, it must
	// return a non-nil error.
	Rewrite(out, in []byte) ([]byte, error)
}

type identity struct{}

func (identity) Rewrite(out, in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MultiRewriter constructs a Rewriter which applies all rewriters passed as
// arguments.
func MultiRewriter(rewriters ...Rewriter) Rewriter {
	_ = "STUB: not implemented"
	return *new(Rewriter)
}

type multiRewriter struct {
	rewriters []Rewriter
}

func (m *multiRewriter) Rewrite(out, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RewriteFunc is a function type implementing the Rewriter interface.
type RewriteFunc func([]byte, []byte) ([]byte, error)

// Rewrite satisfies the Rewriter interface.
func (r RewriteFunc) Rewrite(out, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"

	// MessageRewriter maps field numbers to rewrite rules, satisfying the Rewriter
	// interace to support composing rewrite rules.
	return nil, nil
}

type MessageRewriter []Rewriter

// Rewrite applies the rewrite rule matching f in r, satisfies the Rewriter
// interface.
func (r MessageRewriter) Rewrite(out, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fieldset []uint64

func makeFieldset(n int) fieldset { _ = "STUB: not implemented"; return *new(fieldset) }

func (f fieldset) len() int { _ = "STUB: not implemented"; return 0 }

func (f fieldset) has(i int) bool { _ = "STUB: not implemented"; return false }

func (f fieldset) set(i int) { _ = "STUB: not implemented"; return }

func (f fieldset) unset(i int) { _ = "STUB: not implemented"; return }

func (f fieldset) index(i int) (int, int) {
	_ = "STUB: not implemented"
	return 0,

		// ParseRewriteTemplate constructs a Rewriter for a protobuf type using the
		// given json template to describe the rewrite rules.
		//
		// The json template contains a representation of the message that is used as the
		// source values to overwrite in the protobuf targeted by the resulting rewriter.
		//
		// The rules are an optional set of RewriterRules that can provide alternative
		// Rewriters from the default used for the field type. These rules are given the
		// json.RawMessage bytes from the template, and they are expected to create a
		// Rewriter to be applied against the target protobuf.
		0
}

func ParseRewriteTemplate(typ Type, jsonTemplate []byte, rules ...RewriterRules) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplate(t Type, f FieldNumber, j json.RawMessage, rule any) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateBool(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateInt32(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateInt64(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateSint32(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateSint64(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateUint32(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateUint64(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateFix32(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateFix64(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateSfix32(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateSfix64(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateFloat(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateDouble(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateString(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateBytes(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateMap(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

func parseRewriteTemplateStruct(t Type, f FieldNumber, j json.RawMessage, rules ...RewriterRules) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

type embddedRewriter struct {
	number  FieldNumber
	message MessageRewriter
}

func (f *embddedRewriter) Rewrite(out, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RewriterRules defines a set of rules for overriding the Rewriter used for any
// particular field. These maps may be nested for defining rules for struct members.
//
// For example:
//
//	rules := proto.RewriterRules {
//		"flags": proto.BitOr[uint64]{},
//		"nested": proto.RewriterRules {
//			"name": myCustomRewriter,
//		},
//	}
type RewriterRules map[string]any

// Rewriterer is the interface for producing a Rewriter for a given Type, FieldNumber
// and json.RawMessage. The JSON value is the JSON-encoded payload that should be
// decoded to produce the appropriate Rewriter. Implementations of the Rewriterer
// interface are added to the RewriterRules to specify the rules for performing
// custom rewrite logic.
type Rewriterer interface {
	Rewriter(Type, FieldNumber, json.RawMessage) (Rewriter, error)
}

// BitOr implments the Rewriterer interface for providing a bitwise-or rewrite
// logic for integers rather than replacing them. Instances of this type are
// zero-size, carrying only the generic type for creating the appropriate
// Rewriter when requested.
//
// Adding these to a RewriterRules looks like:
//
//	rules := proto.RewriterRules {
//		"flags": proto.BitOr[uint64]{},
//	}
//
// When used as a rule when rewriting from a template, the BitOr expects a JSON-
// encoded integer passed into the Rewriter method. This parsed integer is then
// used to perform a bitwise-or against the protobuf message that is being rewritten.
//
// The above example can then be used like:
//
//	template := []byte(`{"flags": 8}`) // n |= 0b1000
//	rw, err := proto.ParseRewriteTemplate(typ, template, rules)
type BitOr[T integer] struct{}

// integer is the contraint used by the BitOr Rewriterer and the bitOrRW Rewriter.
// Because these perform bitwise-or operations, the types must be integer-like.
type integer interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
}

// Rewriter implements the Rewriterer interface. The JSON value provided to this
// method comes from the template used for rewriting. The returned Rewriter will use
// this JSON-encoded integer to perform a bitwise-or against the protobuf message
// that is being rewritten.
func (BitOr[T]) Rewriter(t Type, f FieldNumber, j json.RawMessage) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

// BitOrRewriter creates a bitwise-or Rewriter for a given field type and number.
// The mask is the value or'ed with values in the target protobuf.
func BitOrRewriter[T integer](t Type, f FieldNumber, mask T) (Rewriter, error) {
	_ = "STUB: not implemented"
	return *new(Rewriter), nil
}

// bitOrRW is the Rewriter returned by the BitOr Rewriter method.
type bitOrRW[T integer] struct {
	mask T
	t    Type
	f    FieldNumber
}

// Rewrite implements the Rewriter interface performing a bitwise-or between the
// template value and the input value.
func (r bitOrRW[T]) Rewrite(out, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Kind is validated when creating instances
