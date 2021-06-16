//+build one dir

package main

import (
	"fmt"

	"example.com/build/tags/registry"
)

func init() {
	registry.Register("one.go")
}

// OneRunner ...
func OneRunner() {
	fmt.Println("Running in example.com/build/tags@one.go")
}
