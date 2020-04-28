// Package slimit provides size boundaries on UTF-8 data.
package slimit

// SplitAt cuts on index i, i.e., s is separated after i bytes.
// If a splitup would interrupt a UTF-8 byte sequence, then the
// incomplete tail (from head) goes to remainder.
//
// The concatenation of head and remainder equals s for any i.
// Head is empty when i ≤ 0. Remainder is empty when i ≥ len(s).
func SplitAt(s string, i int) (head, remainder string) {
	if len(s) <= i {
		return s, ""
	}
	if i <= 0 {
		return "", s
	}
	if isTail(s[i]) {
		// cut-off in multi-byte sequence
		for j := i - 1; j >= 0; j-- {
			if isTail(s[j]) {
				continue
			}

			// Check whether the tail actually reaches s[n].
			// The most significant bits are 0b110… with 2 byte
			// sequences, 0b1110… with 3 and 0b11110… with 4.
			if byte(int8(-128)>>uint(i-j)) <= s[j] {
				return s[:j], s[j:]
			}
			break
		}
		// broken sequence
	}
	return s[:i], s[i:]
}

// isTail returns whether c is a continuation of a muli-byte sequence.
func isTail(c byte) bool { return c>>6 == 0b10 }

// NBytes limits the string to a maximum number of bytes.
// If a cut-off would interrupt a UTF-8 byte sequence, then the
// incomplete tail is left out.
//
// The return is a prefix of s for any n.
// The return is empty when n ≤ 0.
// The return is s when n ≥ len(s).
func NBytes(s string, n int) string {
	head, _ := SplitAt(s, n)
	return head
}
