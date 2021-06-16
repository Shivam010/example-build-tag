//+build dir

package main

import (
	"fmt"

	dir "example.com/build/tags/directory"
	"example.com/build/tags/registry"
)

func main() {
	registry.Register("directory.go")
	Directory()
	dir.Extra()
	registry.ListAll()
	fmt.Println("After all that")
	OneRunner()
}

func Directory() {
	fmt.Println("Running directory with build tag 'dir'")
}
