package main

import (
	"fmt"
	"os"
)

func main() {
	if err := CmdInit(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
