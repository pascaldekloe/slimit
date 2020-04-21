package slimit_test

import (
	"fmt"

	"github.com/pascaldekloe/slimit"
)

func ExampleNBytes() {
	fmt.Println(slimit.NBytes("👋🌍", 5))
	// Output: 👋
}

func ExampleSplitAt() {
	head, remainder := slimit.SplitAt("🪨📄✂️", 5)
	fmt.Println(head, "&", remainder)
	// Output: 🪨 & 📄✂️
}
