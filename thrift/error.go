package thrift

type MissingField struct {
	Field Field
}

func (e *MissingField) Error() string { _ = "STUB: not implemented"; return "" }

type TypeMismatch struct {
	Expect Type
	Found  Type
	item   string
}

func (e *TypeMismatch) Error() string { _ = "STUB: not implemented"; return "" }

type decodeError struct {
	base error
	path []error
}

func (e *decodeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *decodeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func with(base, elem error) error { _ = "STUB: not implemented"; return nil }

type decodeErrorField struct {
	cause Field
}

func (d *decodeErrorField) Error() string { _ = "STUB: not implemented"; return "" }

type decodeErrorList struct {
	cause List
	index int
}

func (d *decodeErrorList) Error() string { _ = "STUB: not implemented"; return "" }

type decodeErrorSet struct {
	cause Set
	index int
}

func (d *decodeErrorSet) Error() string { _ = "STUB: not implemented"; return "" }

type decodeErrorMap struct {
	cause Map
	index int
}

func (d *decodeErrorMap) Error() string { _ = "STUB: not implemented"; return "" }

func dontExpectEOF(err error) error { _ = "STUB: not implemented"; return nil }
