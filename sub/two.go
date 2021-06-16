//+build two

package sub

import "example.com/build/tags/registry"

func init() {
	registry.Register("sub/two.go")
}
