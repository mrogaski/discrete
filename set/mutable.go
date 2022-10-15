// Package set implements generic set types.
//
// Member type must be comparable.
package set

// An MutableSet is a collection of unique members with mutable state.
type MutableSet[T comparable] struct {
	elements map[T]nothing
}

// NewMutableSet returns an immutable set containing any member members passed as arguments.
// Any duplicate arguments will be ignored.
func NewMutableSet[T comparable](elements ...T) *MutableSet[T] {
	return &MutableSet[T]{elements: newSet(elements...)}
}

// Contains returns true if the given element is a member of the set.
func (s *MutableSet[T]) Contains(elem T) bool {
	_, found := s.elements[elem]

	return found
}

// Size returns a count of the number of members contained in the set.
func (s *MutableSet[T]) Size() int {
	return len(s.elements)
}

// Members returns a slice containing the set members.
// There is no guaranteed ordering.
func (s *MutableSet[T]) Members() []T {
	return members(s.elements)
}

// Insert adds a new element to the set.
// This method is idempotent.
func (s *MutableSet[T]) Insert(elem T) T {
	var nihil nothing

	s.elements[elem] = nihil

	return elem
}

// Delete removes the specified element from the set.
// This method is idempotent.
func (s *MutableSet[T]) Delete(elem T) T {
	delete(s.elements, elem)

	return elem
}

// Copy returns a new set with the same members as the original set.
func (s *MutableSet[T]) Copy() *MutableSet[T] {
	return NewMutableSet(members(s.elements)...)
}

// IsSubset returns true is the set is a subset of the argument set, false otherwise.
func (s *MutableSet[T]) IsSubset(sp *MutableSet[T]) bool {
	for elem := range s.elements {
		_, found := sp.elements[elem]
		if !found {
			return false
		}
	}

	return true
}

// IsEqual returns true is the set is equal to the argument set, false otherwise.
func (s *MutableSet[T]) IsEqual(sp *MutableSet[T]) bool {
	return s.IsSubset(sp) && sp.IsSubset(s)
}

// IsProperSubset returns true is the set is a subset of the argument set but not equal to the argument set.
func (s *MutableSet[T]) IsProperSubset(sp *MutableSet[T]) bool {
	return s.IsSubset(sp) && !s.IsEqual(sp)
}

// Union returns a set whose members are a disjunction of both the receiver and argument sets, belonging to either set.
func (s *MutableSet[T]) Union(sp *MutableSet[T]) *MutableSet[T] {
	combined := make([]T, 0, len(s.elements)+len(sp.elements))
	combined = append(combined, members(s.elements)...)
	combined = append(combined, members(sp.elements)...)

	return NewMutableSet(combined...)
}

// Intersection returns a set whose members are a conjunction of both the receiver and argument sets, belonging
// to one set or the other but not both.
func (s *MutableSet[T]) Intersection(sp *MutableSet[T]) *MutableSet[T] {
	members := make([]T, 0)

	for k := range s.elements {
		_, found := sp.elements[k]

		if found {
			members = append(members, k)
		}
	}

	return NewMutableSet(members...)
}

// Difference returns a set whose members belong to the receiver but not the set passed as an argument.
func (s *MutableSet[T]) Difference(sp *MutableSet[T]) *MutableSet[T] {
	members := make([]T, 0)

	for k := range s.elements {
		_, found := sp.elements[k]

		if !found {
			members = append(members, k)
		}
	}

	return NewMutableSet(members...)
}

// SymmetricDifference returns a set whose members are an exclusive disjunction of both the receiver and argument sets,
// belonging to either set but not both.
func (s *MutableSet[T]) SymmetricDifference(sp *MutableSet[T]) *MutableSet[T] {
	return s.Union(sp).Difference(s.Intersection(sp))
}
