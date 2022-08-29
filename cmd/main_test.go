package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestMainCli test main application
func TestMainCli(t *testing.T) {
	var tests = []struct {
		testname string
		Args     []string
		want     string
	}{
		{
			"Verion command",
			[]string{"./validate-csv", "version"},
			"develop",
		},
		{
			"Validate command",
			[]string{"./validate-csv", "validate", "--file", "./../example.csv"},
			"Validation successful",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			os.Args = tt.Args

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			if !strings.Contains(string(out), tt.want) {
				t.Errorf("could not find expected %s in %s", tt.want, string(out))
			}
		})
	}
}
