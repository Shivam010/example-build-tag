//+build two

package main

import "example.com/build/tags/registry"

func init() {
	registry.Register("two.go")
}
