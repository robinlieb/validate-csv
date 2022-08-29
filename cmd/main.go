package main

import (
	"fmt"
	"os"
)

var version = "develop"

func main() {
	if err := CmdInit(version); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
