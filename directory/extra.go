//+build dir

package dir

import (
	"fmt"

	"example.com/build/tags/registry"
)

func init() {
	registry.Register("directory/extra.go")
}

func Extra() {
	fmt.Println("Running extra with build tag 'dir'")
}
