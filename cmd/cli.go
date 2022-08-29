package main

import (
	"github.com/spf13/cobra"
)

func CmdInit(version string) error {

	rootCmd := &cobra.Command{
		Use:   "validate-cli",
		Short: "CLI tool to validate CSV by duplicates and sums of rows.",
		Long: ` CLI tool to validate CSV by duplicates and sums of rows.
Complete documentation is available at https://github.com/robinlieb/validate-csv
		`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	validateCmd := NewValidateCommand()
	versionCmd := NewVersionCommand(version)

	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(versionCmd)

	validateCmd.PersistentFlags().String("file", "", "Input file to validate. It should be in CSV format.")

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
