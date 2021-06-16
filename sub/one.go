//+build one

package sub

import "example.com/build/tags/registry"

func init() {
	registry.Register("sub/one.go")
}
