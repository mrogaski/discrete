// Package set implements generic set types.
//
// Member type must be comparable.
package set

// An ImmutableSet is a collection of unique members.
type ImmutableSet[T comparable] struct {
	elements map[T]nothing
}

type nothing struct{}

type set[T comparable] map[T]nothing

// NewImmutableSet returns an immutable set containing any member members passed as arguments.
// Any duplicate arguments will be ignored.
func NewImmutableSet[T comparable](elements ...T) *ImmutableSet[T] {
	return &ImmutableSet[T]{elements: newSet(elements...)}
}

func newSet[T comparable](elements ...T) set[T] {
	var nihil nothing
	result := make(set[T], len(elements))

	for _, elem := range elements {
		result[elem] = nihil
	}

	return result
}

// Contains returns true if the given element is a member of the set.
func (s *ImmutableSet[T]) Contains(elem T) bool {
	_, found := s.elements[elem]

	return found
}

// Size returns a count of the number of members contained in the set.
func (s *ImmutableSet[T]) Size() int {
	return len(s.elements)
}

// Members returns a slice containing the set members.
// There is no guaranteed ordering.
func (s *ImmutableSet[T]) Members() []T {
	return members(s.elements)
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

// Copy returns a new set with the same members as the original set.
func (s *ImmutableSet[T]) Copy() *ImmutableSet[T] {
	return NewImmutableSet(members(s.elements)...)
}

// IsEqual returns true is the set is equal to the argument set, false otherwise.
func (s *ImmutableSet[T]) IsEqual(sp *ImmutableSet[T]) bool {
	return s.IsSubset(sp) && sp.IsSubset(s)
}

// IsSubset returns true is the set is a subset of the argument set, false otherwise.
func (s *ImmutableSet[T]) IsSubset(sp *ImmutableSet[T]) bool {
	for elem := range s.elements {
		_, found := sp.elements[elem]
		if !found {
			return false
		}
	}

	return true
}

// IsProperSubset returns true is the set is a subset of the argument set but not equal to the argument set.
func (s *ImmutableSet[T]) IsProperSubset(sp *ImmutableSet[T]) bool {
	return s.IsSubset(sp) && !s.IsEqual(sp)
}

// Union returns a set whose members are a disjunction of both the receiver and argument sets, belonging to either set.
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

// Intersection returns a set whose members are a conjunction of both the receiver and argument sets, belonging
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

// Difference returns a set whose members belong to the receiver but not the set passed as an argument.
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

// SymmetricDifference returns a set whose members are an exclusive disjunction of both the receiver and argument sets,
// belonging to either set but not both.
func (s *ImmutableSet[T]) SymmetricDifference(sp *ImmutableSet[T]) *ImmutableSet[T] {
	return s.Union(sp).Difference(s.Intersection(sp))
}
