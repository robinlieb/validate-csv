// Package helper contains helper function
package helper

import (
	"fmt"
	"os"
)

func PrintErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
