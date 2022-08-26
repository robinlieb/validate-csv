package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	filename := validateCmd.String("file", "", "Input file to validate")

	if len(os.Args) < 2 {
		fmt.Println("expected 'validate' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "validate":
		validateCmd.Parse(os.Args[2:])
		err := HandleValidate(*filename)
		if err != nil {
			fmt.Printf("Recieved error: %s\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("expected 'validate' subcommand")
		os.Exit(1)
	}
}

func HandleValidate(filename string) error {
	return nil
}
