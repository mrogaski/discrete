package set_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrogaski/go-set"
)

func TestNewImmutableSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		members []rune
	}{
		{name: "empty", members: []rune{}},
		{name: "1 member", members: []rune{'A'}},
		{name: "2 members", members: []rune{'A', 'Z'}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := set.NewImmutableSet(tt.members...)

			assert.IsType(t, &set.ImmutableSet[rune]{}, s)
		})
	}
}

func TestImmutableSet_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		members []rune
		elem    rune
		want    bool
	}{
		{name: "empty", members: []rune{}, elem: 'A', want: false},
		{name: "hit", members: []rune{'U', 'A'}, elem: 'A', want: true},
		{name: "miss", members: []rune{'U', 'A'}, elem: 'Z', want: false},
		{name: "hit only member", members: []rune{'A'}, elem: 'A', want: true},
		{name: "miss only member", members: []rune{'A'}, elem: 'Z', want: false},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := set.NewImmutableSet(tt.members...)

			assert.Equal(t, tt.want, s.Contains(tt.elem))
		})
	}
}

func TestImmutableSet_Size(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		members []rune
		want    int
	}{
		{name: "empty", members: []rune{}, want: 0},
		{name: "1 member", members: []rune{'A'}, want: 1},
		{name: "2 member", members: []rune{'A', 'B'}, want: 2},
		{name: "duplicates", members: []rune{'A', 'B', 'A', 'C', 'A', 'B'}, want: 3},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := set.NewImmutableSet(tt.members...)

			assert.Equal(t, tt.want, s.Size())
		})
	}
}

func TestImmutableSet_Members(t *testing.T) {
	t.Parallel()

	sequence := make([]rune, 0)
	for i := 0x20; i < 0x80; i++ {
		sequence = append(sequence, rune(i))
	}

	tests := []struct {
		name    string
		members []rune
	}{
		{name: "empty", members: []rune{}},
		{name: "1 member", members: []rune{'A'}},
		{name: "sequence", members: sequence},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := set.NewImmutableSet(tt.members...)

			assert.Equal(t, sorted(tt.members), sorted(s.Members()))
		})
	}
}

func sorted(input []rune) []rune {
	result := make([]rune, len(input))
	copy(result, input)

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	return input
}
