// Package set implements generic set types.
//
// Member type must be comparable.
package set

// An ImmutableSet is a collection of unique elements.
type ImmutableSet[T comparable] struct {
	elements map[T]nothing
}

type nothing struct{}

// NewImmutableSet returns an immutable set containing any member elements passed as arguments.
// Any duplicate arguments will be ignored.
func NewImmutableSet[T comparable](elements ...T) *ImmutableSet[T] {
	var nihil nothing

	seq := make(map[T]nothing, len(elements))

	for _, elem := range elements {
		seq[elem] = nihil
	}

	return &ImmutableSet[T]{elements: seq}
}

// Contains returns true if the given element is a member of the set.
func (s *ImmutableSet[T]) Contains(elem T) bool {
	_, found := s.elements[elem]

	return found
}

// Size returns a count of the number of elements contained in the set.
func (s *ImmutableSet[T]) Size() int {
	return len(s.elements)
}

// Members returns a slice containing the set elements.
// There is no guaranteed ordering.
func (s *ImmutableSet[T]) Members() []T {
	result := make([]T, len(s.elements))

	i := 0

	for elem := range s.elements {
		result[i] = elem
		i++
	}

	return result
}

// Copy returns a new set with the same elements as the original set.
func (s *ImmutableSet[T]) Copy() *ImmutableSet[T] {
	return NewImmutableSet(s.Members()...)
}

// IsEqual returns true is the set is equal to the argument set, false otherwise.
func (s *ImmutableSet[T]) IsEqual(sp *ImmutableSet[T]) bool {
	return s.IsSubset(sp) && sp.IsSubset(s)
}

// IsSubset returns true is the set is a subset of the argument set, false otherwise.
func (s *ImmutableSet[T]) IsSubset(sp *ImmutableSet[T]) bool {
	for _, elem := range s.Members() {
		if !sp.Contains(elem) {
			return false
		}
	}

	return true
}

// IsProperSubset returns true is the set is a subset of the argument set but not equal to the argument set.
func (s *ImmutableSet[T]) IsProperSubset(sp *ImmutableSet[T]) bool {
	return s.IsSubset(sp) && !s.IsEqual(sp)
}

// Union returns a set whose elements are a disjunction of both the receiver and argument sets, belonging to either set.
func (s *ImmutableSet[T]) Union(sp *ImmutableSet[T]) *ImmutableSet[T] {
	members := make([]T, s.Size()+sp.Size())

	i := 0

	for elem := range s.elements {
		members[i] = elem
		i++
	}

	for elem := range sp.elements {
		members[i] = elem
		i++
	}

	return NewImmutableSet[T](members...)
}

// Intersection returns a set whose elements are a conjunction of both the receiver and argument sets, belonging
// to one set or the other but not both.
func (s *ImmutableSet[T]) Intersection(sp *ImmutableSet[T]) *ImmutableSet[T] {
	members := make([]T, 0)

	for k := range s.elements {
		_, found := sp.elements[k]

		if found {
			members = append(members, k)
		}
	}

	return NewImmutableSet(members...)
}

// Difference returns a set whose elements belong to the receiver but not the set passed as an argument.
func (s *ImmutableSet[T]) Difference(sp *ImmutableSet[T]) *ImmutableSet[T] {
	members := make([]T, 0)

	for k := range s.elements {
		_, found := sp.elements[k]

		if !found {
			members = append(members, k)
		}
	}

	return NewImmutableSet(members...)
}

// SymmetricDifference returns a set whose elements are an exclusive disjunction of both the receiver and argument sets,
// belonging to either set but not both.
func (s *ImmutableSet[T]) SymmetricDifference(sp *ImmutableSet[T]) *ImmutableSet[T] {
	return s.Union(sp).Difference(s.Intersection(sp))
}
