package json

import (
	"reflect"
)

// All spaces characters defined in the json specification.
const (
	sp = ' '
	ht = '\t'
	nl = '\n'
	cr = '\r'
)

func internalParseFlags(b []byte) (flags ParseFlags) {
	_ = "STUB: not implemented"
	// Don't consider surrounding whitespace
	return *new(ParseFlags)
}

func skipSpaces(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func skipSpacesN(b []byte) ([]byte, int) { _ = "STUB: not implemented"; return nil, 0 }

func trimTrailingSpaces(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func trimTrailingSpacesN(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// parseInt parses a decimal representation of an int64 from b.
//
// The function is equivalent to calling strconv.ParseInt(string(b), 10, 64) but
// it prevents Go from making a memory allocation for converting a byte slice to
// a string (escape analysis fails due to the error returned by strconv.ParseInt).
//
// Because it only works with base 10 the function is also significantly faster
// than strconv.ParseInt.
func (d decoder) parseInt(b []byte, t reflect.Type) (int64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// was this actually a float?

// parseUint is like parseInt but for unsigned integers.
func (d decoder) parseUint(b []byte, t reflect.Type) (uint64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// was this actually a float?

// parseUintHex parses a hexadecimanl representation of a uint64 from b.
//
// The function is equivalent to calling strconv.ParseUint(string(b), 16, 64) but
// it prevents Go from making a memory allocation for converting a byte slice to
// a string (escape analysis fails due to the error returned by strconv.ParseUint).
//
// Because it only works with base 16 the function is also significantly faster
// than strconv.ParseUint.
func (d decoder) parseUintHex(b []byte) (uint64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (d decoder) parseNull(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

func (d decoder) parseTrue(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

func (d decoder) parseFalse(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

func (d decoder) parseNumber(b []byte) (v, r []byte, kind Kind, err error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

// Assume it's an unsigned integer at first.

// sign

// integer part

// decimal part

// exponent part

func (d decoder) parseUnicode(b []byte) (rune, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (d decoder) parseString(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

// This is an optimization for short strings. We read 8/16 bytes,
// and XOR each with 0x22 (") so that these bytes (and only
// these bytes) are now zero. We use the hasless(u,1) trick
// from https://graphics.stanford.edu/~seander/bithacks.html#ZeroInWord
// to determine whether any bytes are zero. Finally, we CTZ
// to find the index of that byte.

func (d decoder) parseStringUnquote(b []byte, r []byte) ([]byte, []byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, false, nil
}

// trim the quotes

// simple escaped character

// not sure what this escape sequence is

func appendRune(b []byte, r rune) []byte { _ = "STUB: not implemented"; return nil }

func appendCoerceInvalidUTF8(b []byte, s []byte) []byte { _ = "STUB: not implemented"; return nil }

func (d decoder) parseObject(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

func (d decoder) parseArray(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

func (d decoder) parseValue(b []byte) ([]byte, []byte, Kind, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(Kind), nil
}

func hasNullPrefix(b []byte) bool { _ = "STUB: not implemented"; return false }

func hasTruePrefix(b []byte) bool { _ = "STUB: not implemented"; return false }

func hasFalsePrefix(b []byte) bool { _ = "STUB: not implemented"; return false }

func hasPrefix(b []byte, s string) bool { _ = "STUB: not implemented"; return false }

func hasLeadingSign(b []byte) bool { _ = "STUB: not implemented"; return false }

func hasLeadingZeroes(b []byte) bool { _ = "STUB: not implemented"; return false }

func appendToLower(b, s []byte) []byte {
	_ = "STUB: not implemented"
	// fast path for ascii strings
	return nil
}

func foldRune(r rune) rune { _ = "STUB: not implemented"; return 0 }
