// Package slimit provides size boundaries on UTF-8 data.
package slimit

// SplitAt cuts on index i, i.e., s is separated after i bytes.
// If a splitup would interrupt a UTF-8 byte sequence, then the
// incomplete tail (from head) goes to remainder.
//
// The concatenation of head and remainder equals s for any i.
// Head is empty when i ≤ 0. Remainder is empty when i ≥ len(s).
func SplitAt(s string, i int) (head, remainder string) {
	head = NBytes(s, i)
	return head, s[len(head):]
}

// NBytes limits the string to a maximum number of bytes.
// If a cut-off would interrupt a UTF-8 byte sequence, then the
// incomplete tail is left out.
//
// The return is a prefix of s for any n.
// The return is empty when n ≤ 0.
// The return is s when n ≥ len(s).
func NBytes(s string, n int) string {
	if len(s) <= n {
		return s
	}

	for i := n; i > 0; i-- {
		if s[i]>>6 != 2 || s[i-1] < 128 {
			return s[:i]
		}
	}

	return ""
}
