package slimit

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestByteCuts(t *testing.T) {
	for _, s := range []string{"", "a", "ab", "abc", "𝜹🧨", "e\x8f", "0℉"} {
		for i := -2; i <= len(s)+2; i++ {
			head, remainder := SplitAt(s, i)

			switch {
			case !strings.HasPrefix(s, head):
				t.Errorf("%d bytes of %q got %q; prefix change", i, s, head)

			case head+remainder != s:
				t.Errorf("Split(%q, %d) got (%q, %q); content change", s, i, head, remainder)

			case i <= 0:
				if head != "" || remainder != s {
					t.Errorf("Split(%q, %d) got (%q, %q), want all in remainder", s, i, head, remainder)
				}

			case i >= len(s):
				if head != s || remainder != "" {
					t.Errorf("Split(%q, %d) got (%q, %q), want all in head", s, i, head, remainder)
				}

			default:
				want := s[:i]
				for !utf8.ValidString(want) {
					want = want[:len(want)-1]
				}
				if head != want {
					t.Errorf("%d bytes of %q got %q, want %q", i, s, head, want)
				}
			}
		}
	}
}
