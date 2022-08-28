package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/robinlieb/validate-csv/business/validate"
)

func CmdInit() error {
	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	filename := validateCmd.String("file", "", "Input file to validate. It should be in CSV format.")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'validate' subcommand")
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
		fmt.Println("Expected 'validate' subcommand")
		os.Exit(1)
	}

	return nil
}

func HandleValidate(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var mat [][]int

	for _, row := range data {
		var matRow []int
		for _, value := range row {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			matRow = append(matRow, i)
		}
		mat = append(mat, matRow)
	}

	err = validate.Validate(mat)
	if err != nil {
		return err
	}

	return nil
}
