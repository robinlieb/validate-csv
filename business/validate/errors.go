package validate

import (
	"errors"
)

// ErrDuplicateNumbers is the error used when duplicate numbers found in row.
var ErrDuplicateNumbers = errors.New("Duplicate numbers provided in input")

// ErrSumsNotEqual is the error used when sums of rows are not equal.
var ErrSumsNotEqual = errors.New("Sums of all rows not equal")
