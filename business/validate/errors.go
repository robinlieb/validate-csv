package validate

import (
	"errors"
)

// ErrDuplicateNumbers is the error used when duplicate numbers found in row.
var ErrDuplicateNumbers = errors.New("duplicate numbers provided in input")

// ErrSumsNotEqual is the error used when sums of rows are not equal.
var ErrSumsNotEqual = errors.New("sums of all rows not equal")
