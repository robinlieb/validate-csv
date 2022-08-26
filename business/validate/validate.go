// Package validate contains functions to validate two dimensional arrays
package validate

// Validate validates two dimensinal array input by checking if rows contains
// duplicates and sums of rows are equal. Returns either ErrDuplicateNumbers
// or ErrSumsNotEqual Error in case of unsuccessful validation. In case input
// is corrupted in both ways it will return ErrDuplicateNumbers.
func Validate(in [][]int) error {
	var sums []int
	for _, row := range in {

		for j, elem := range row {
			for k := j + 1; k < len(row); k++ {
				if elem == row[k] {
					return ErrDuplicateNumbers
				}
			}
		}

		var sum int
		for _, elem := range row {
			sum += elem
		}
		sums = append(sums, sum)
	}

	for _, sum := range sums {
		if sum != sums[0] {
			return ErrSumsNotEqual
		}
	}

	return nil
}
