package slimit_test

import (
	"fmt"

	"github.com/pascaldekloe/slimit"
)

func ExampleSplitAt() {
	head, remainder := slimit.SplitAt("🪨📄✂️", 5)
	fmt.Println(head, "&", remainder)
	// Output: 🪨 & 📄✂️
}
