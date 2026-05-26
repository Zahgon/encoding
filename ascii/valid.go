//go:generate go run valid_asm.go -out valid_amd64.s -stubs valid_amd64.go
package ascii

// Valid returns true if b contains only ASCII characters.
func Valid(b []byte) bool { _ = "STUB: not implemented"; return false }

// ValidBytes returns true if b is an ASCII character.
func ValidByte(b byte) bool { _ = "STUB: not implemented"; return false }

// ValidBytes returns true if b is an ASCII character.
func ValidRune(r rune) bool { _ = "STUB: not implemented"; return false }

// ValidString returns true if s contains only ASCII characters.
func ValidString(s string) bool { _ = "STUB: not implemented"; return false }
