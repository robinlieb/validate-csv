package validate

import (
	"testing"
)

// TestValidate calls validator.Validate
func TestValidate(t *testing.T) {
	var tests = []struct {
		testname string
		input    [][]int
		want     error
	}{
		{
			"Valid input",
			[][]int{
				{1, 2, 3},
				{3, 2, 1},
				{1, 2, 3},
			},
			nil,
		},
		{
			"Duplicate numbers",
			[][]int{
				{1, 2, 2},
				{1, 2, 2},
				{1, 2, 2},
			},
			ErrDuplicateNumbers,
		},
		{
			"Sums not equal",
			[][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 2, 8},
			},
			ErrSumsNotEqual,
		},
		{
			"Duplicate numbers and sums not equal",
			[][]int{
				{1, 2, 2},
				{1, 2, 4},
				{1, 2, 8},
			},
			ErrDuplicateNumbers,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := Validate(tt.input)

			if err != tt.want {
				t.Errorf("got %s, want %s", err, tt.want)
			}
		})
	}
}
