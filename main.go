//+build !dir

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"example.com/build/tags/registry"
	_ "example.com/build/tags/sub"
)

func main() {
	MainRunner()
	registry.Register("main.go")
	now := time.Now()
	fmt.Println(now.Format(time.Kitchen))
	registry.ListAll()
	fmt.Println()
	fmt.Println("------------")

	cmd := exec.Command("go", "run", "-tags", "dir", "example.com/build/tags")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("error in cmd", err)
	}
	fmt.Println("------------")
	fmt.Println()
	fmt.Println("Done")
}

// MainRunner ...
func MainRunner() {
	fmt.Println("Running in example.com/build/tags@main.go")
}
