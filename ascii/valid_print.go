//go:generate go run valid_print_asm.go -out valid_print_amd64.s -stubs valid_print_amd64.go
package ascii

// Valid returns true if b contains only printable ASCII characters.
func ValidPrint(b []byte) bool { _ = "STUB: not implemented"; return false }

// ValidBytes returns true if b is an ASCII character.
func ValidPrintByte(b byte) bool { _ = "STUB: not implemented"; return false }

// ValidBytes returns true if b is an ASCII character.
func ValidPrintRune(r rune) bool { _ = "STUB: not implemented"; return false }

// ValidString returns true if s contains only printable ASCII characters.
func ValidPrintString(s string) bool { _ = "STUB: not implemented"; return false }
