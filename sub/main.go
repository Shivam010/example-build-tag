package sub

import (
	"example.com/build/tags/indirect"
	"example.com/build/tags/registry"
)

func init() {
	registry.Register("sub/main.go")
}

func IndirectUsage() {
	_ = indirect.Usage{}
}
