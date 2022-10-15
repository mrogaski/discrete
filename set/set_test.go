package set_test

import "sort"

func sorted(input []rune) []rune {
	result := make([]rune, len(input))
	copy(result, input)

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	return input
}
