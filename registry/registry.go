package registry

import "fmt"

var registry = map[string]struct{}{
	"registry/registry.go": {},
}

// Register object in registry
func Register(val string) {
	registry[val] = struct{}{}
}

// ListAll objects of registry
func ListAll() {
	fmt.Println("Listing all registeries:")
	for r := range registry {
		fmt.Println(" -", r)
	}
	fmt.Println()
}
