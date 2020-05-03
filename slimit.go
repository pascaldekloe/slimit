// Package slimit provides size boundaries on UTF-8 data.
package slimit

import "unicode/utf8"

// SplitAt cuts on index i, i.e., s is separated after i bytes.
// If a splitup would interrupt a UTF-8 byte sequence, then the
// incomplete tail (from head) goes to remainder.
//
// The concatenation of head and remainder equals s for any i.
// Head is empty when i ≤ 0. Remainder is empty when i ≥ len(s).
func SplitAt(s string, i int) (head, remainder string) {
	if i >= len(s) {
		return s, ""
	}
	if i <= 0 {
		return "", s
	}
	if isTail(s[i]) {
		// cut-off in multi-byte sequence
		j := i - 1
		if j > 0 && isTail(s[j]) {
			j--
			if j > 0 && isTail(s[j]) {
				j--
				// encoding is limited to 4 bytes
			}
		}
		_, size := utf8.DecodeRuneInString(s[j:])
		if size > i-j {
			return s[:j], s[j:]
		}
		// broken sequence
	}
	return s[:i], s[i:]
}

// isTail returns whether c is a continuation of a muli-byte sequence.
func isTail(c byte) bool { return c>>6 == 0b10 }
