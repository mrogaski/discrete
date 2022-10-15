// Package set implements generic set types.
//
// Member type must be comparable.
package set

type nothing struct{}

type set[T comparable] map[T]nothing

func newSet[T comparable](elements ...T) set[T] {
	var nihil nothing

	result := make(set[T], len(elements))

	for _, elem := range elements {
		result[elem] = nihil
	}

	return result
}

func members[T comparable](s set[T]) []T {
	result := make([]T, len(s))

	i := 0

	for elem := range s {
		result[i] = elem
		i++
	}

	return result
}
