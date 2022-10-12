// Package set implements generic set types.
package set

type ImmutableSet[T comparable] struct{}

func NewImmutableSet[T comparable](members ...[]T) *ImmutableSet[T] {
	return nil
}
