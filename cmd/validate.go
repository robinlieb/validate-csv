package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/robinlieb/validate-csv/business/validate"
	"github.com/robinlieb/validate-csv/foundation/helper"
	"github.com/spf13/cobra"
)

func NewValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Input file to validate.",
		Long:  "Input file to validate. It should be in CSV format.",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				helper.PrintErrorAndExit(err)
			}

			err = HandleValidate(file)
			if err != nil {
				helper.PrintErrorAndExit(err)
			}

			fmt.Println("Validation successful.")
		},
	}

	return cmd
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
