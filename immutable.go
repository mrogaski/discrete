// Package set implements generic set types.
//
// Member type must be comparable.
package set

// An ImmutableSet is a collection of unique elements.
type ImmutableSet[T comparable] struct {
	members map[T]nothing
}

type nothing struct{}

// NewImmutableSet returns an immutable set containing any member elements passed as arguments.
// Any duplicate arguments will be ignored.
func NewImmutableSet[T comparable](members ...T) *ImmutableSet[T] {
	var nihil nothing

	seq := make(map[T]nothing, len(members))

	for _, elem := range members {
		seq[elem] = nihil
	}

	return &ImmutableSet[T]{members: seq}
}

// Contains returns true if the given element is a member of the set.
func (s *ImmutableSet[T]) Contains(elem T) bool {
	_, found := s.members[elem]

	return found
}

// Size returns a count of the number of members contained in the set.
func (s *ImmutableSet[T]) Size() int {
	return len(s.members)
}

// Members returns a slice containing the set members.
// There is no guaranteed ordering.
func (s *ImmutableSet[T]) Members() []T {
	result := make([]T, len(s.members))

	i := 0

	for elem := range s.members {
		result[i] = elem
		i++
	}

	return result
}

// Union returns a set whose members are disjunction of both the receiver and argument sets, belonging to either set.
func (s *ImmutableSet[T]) Union(sp *ImmutableSet[T]) *ImmutableSet[T] {
	members := make([]T, s.Size()+sp.Size())

	i := 0

	for elem := range s.members {
		members[i] = elem
		i++
	}

	for elem := range sp.members {
		members[i] = elem
		i++
	}

	return NewImmutableSet[T](members...)
}

// Difference returns a set whose members belong to the receiver but not the set passed as an argument.
func (s *ImmutableSet[T]) Difference(sp *ImmutableSet[T]) *ImmutableSet[T] {
	members := make([]T, 0)

	for k := range s.members {
		_, found := sp.members[k]

		if !found {
			members = append(members, k)
		}
	}

	return NewImmutableSet(members...)
}
