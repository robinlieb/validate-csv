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
		return ErrCommand
	}

	switch os.Args[1] {
	case "validate":
		validateCmd.Parse(os.Args[2:])
		err := HandleValidate(*filename)
		if err != nil {
			fmt.Printf("Recieved error: %s\n", err)
			return err
		}
	default:
		return ErrCommand
	}

	return nil
}

func HandleValidate(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return ErrFileNotFound
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
