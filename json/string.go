package json

const (
	lsb = 0x0101010101010101
	msb = 0x8080808080808080
)

// escapeIndex finds the index of the first char in `s` that requires escaping.
// A char requires escaping if it's outside of the range of [0x20, 0x7F] or if
// it includes a double quote or backslash. If the escapeHTML mode is enabled,
// the chars <, > and & also require escaping. If no chars in `s` require
// escaping, the return value is -1.
func escapeIndex(s string, escapeHTML bool) int { _ = "STUB: not implemented"; return 0 }

// combine masks before checking for the MSB of each byte. We include
// `n` in the mask to check whether any of the *input* byte MSBs were
// set (i.e. the byte was outside the ASCII range).

func escapeByteRepr(b byte) byte { _ = "STUB: not implemented"; return 0 }

// below return a mask that can be used to determine if any of the bytes
// in `n` are below `b`. If a byte's MSB is set in the mask then that byte was
// below `b`. The result is only valid if `b`, and each byte in `n`, is below
// 0x80.
func below(n uint64, b byte) uint64 { _ = "STUB: not implemented"; return 0 }

// contains returns a mask that can be used to determine if any of the
// bytes in `n` are equal to `b`. If a byte's MSB is set in the mask then
// that byte is equal to `b`. The result is only valid if `b`, and each
// byte in `n`, is below 0x80.
func contains(n uint64, b byte) uint64 { _ = "STUB: not implemented"; return 0 }

// expand puts the specified byte into each of the 8 bytes of a uint64.
func expand(b byte) uint64 { _ = "STUB: not implemented"; return 0 }

func stringToUint64(s string) []uint64 { _ = "STUB: not implemented"; return nil }
