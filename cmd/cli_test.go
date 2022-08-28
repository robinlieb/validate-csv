package main

import (
	"testing"
)

// TestHandleValidate calls validator.HandleValidate
func TestHandleValidate(t *testing.T) {
	var tests = []struct {
		testname string
		input    string
		want     error
	}{
		{
			"Valid input",
			"./../example.csv",
			nil,
		},
		{
			"File not found",
			"not_found.csv",
			ErrFileNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := HandleValidate(tt.input)

			if err != tt.want {
				t.Errorf("got %s, want %s", err, tt.want)
			}
		})
	}
}
