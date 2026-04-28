package main

import (
	"fmt"
	"os"

	"github.com/engram-dev/engram/cmd"
)

// version is set at build time via ldflags.
var version = "dev"

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
