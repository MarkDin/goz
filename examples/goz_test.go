package goz

import (
	"fmt"

	"github.com/MarkDin/goz"
)

func ExampleNewClient() {
	cli := goz.NewClient()

	fmt.Printf("%T", cli)
	// Output: *goz.Request
}
