//go:generate go run equal_fold_asm.go -out equal_fold_amd64.s -stubs equal_fold_amd64.go
package ascii

// EqualFold is a version of bytes.EqualFold designed to work on ASCII input
// instead of UTF-8.
//
// When the program has guarantees that the input is composed of ASCII
// characters only, it allows for greater optimizations.
func EqualFold(a, b []byte) bool { _ = "STUB: not implemented"; return false }

func HasPrefixFold(s, prefix []byte) bool { _ = "STUB: not implemented"; return false }

func HasSuffixFold(s, suffix []byte) bool { _ = "STUB: not implemented"; return false }

// EqualFoldString is a version of strings.EqualFold designed to work on ASCII
// input instead of UTF-8.
//
// When the program has guarantees that the input is composed of ASCII
// characters only, it allows for greater optimizations.
func EqualFoldString(a, b string) bool { _ = "STUB: not implemented"; return false }

func HasPrefixFoldString(s, prefix string) bool { _ = "STUB: not implemented"; return false }

func HasSuffixFoldString(s, suffix string) bool { _ = "STUB: not implemented"; return false }
