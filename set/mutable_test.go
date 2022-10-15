package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrogaski/discrete/set"
)

func TestNewMutableSet(t *testing.T) {
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

			s := set.NewMutableSet(tt.members...)

			assert.IsType(t, &set.MutableSet[rune]{}, s)
		})
	}
}

func TestMutableSet_Contains(t *testing.T) {
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

			s := set.NewMutableSet(tt.members...)

			assert.Equal(t, tt.want, s.Contains(tt.elem))
		})
	}
}

func TestMutableSet_Size(t *testing.T) {
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

			s := set.NewMutableSet(tt.members...)

			assert.Equal(t, tt.want, s.Size())
		})
	}
}

func TestMutableSet_Members(t *testing.T) {
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

			s := set.NewMutableSet(tt.members...)

			assert.Equal(t, sorted(tt.members), sorted(s.Members()))
		})
	}
}

func TestMutableSet_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		members []rune
		element rune
		want    []rune
	}{
		{name: "empty", members: []rune{}, element: 'A', want: []rune{'A'}},
		{name: "new member", members: []rune{'A', 'B'}, element: 'C', want: []rune{'A', 'B', 'C'}},
		{name: "existing member", members: []rune{'A', 'B'}, element: 'B', want: []rune{'A', 'B'}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := set.NewMutableSet(tt.members...)

			assert.Equal(t, tt.element, s.Insert(tt.element))
			assert.Equal(t, sorted(tt.want), sorted(s.Members()))
		})
	}
}

func TestMutableSet_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		members []rune
		element rune
		want    []rune
	}{
		{name: "empty", members: []rune{}, element: 'A', want: []rune{}},
		{name: "existing member", members: []rune{'A', 'B', 'C'}, element: 'C', want: []rune{'A', 'B'}},
		{name: "missing member", members: []rune{'A', 'B'}, element: 'C', want: []rune{'A', 'B'}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := set.NewMutableSet(tt.members...)

			assert.Equal(t, tt.element, s.Delete(tt.element))
			assert.Equal(t, sorted(tt.want), sorted(s.Members()))
		})
	}
}

func TestMutableSet_Copy(t *testing.T) {
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

			s := set.NewMutableSet(tt.members...)
			sp := s.Copy()

			assert.Equal(t, s, sp)
			assert.NotSame(t, s, sp)
		})
	}
}

func TestMutableSet_IsSubset(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want bool
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: true},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: false},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: true},
		{name: "identical", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Y', 'Z'}, want: true},
		{name: "overlap", a: []rune{'X', 'Y', 'Z'}, b: []rune{'W', 'X', 'Y'}, want: false},
		{name: "disjoint", a: []rune{'X', 'Y', 'Z'}, b: []rune{'x', 'y', 'z'}, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.IsSubset(b)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMutableSet_IsEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want bool
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: true},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: false},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: false},
		{name: "identical", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Y', 'Z'}, want: true},
		{name: "overlap", a: []rune{'X', 'Y', 'Z'}, b: []rune{'W', 'X', 'Y'}, want: false},
		{name: "disjoint", a: []rune{'X', 'Y', 'Z'}, b: []rune{'x', 'y', 'z'}, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.IsEqual(b)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMutableSet_IsProperSubset(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want bool
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: false},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: false},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: true},
		{name: "identical", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Y', 'Z'}, want: false},
		{name: "overlap", a: []rune{'X', 'Y', 'Z'}, b: []rune{'W', 'X', 'Y'}, want: false},
		{name: "disjoint", a: []rune{'X', 'Y', 'Z'}, b: []rune{'x', 'y', 'z'}, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.IsProperSubset(b)

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMutableSet_Union(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want []rune
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: []rune{}},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: []rune{'X', 'Y', 'Z'}},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: []rune{'x', 'y', 'z'}},
		{name: "identical", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Y', 'Z'}, want: []rune{'X', 'Y', 'Z'}},
		{name: "overlap", a: []rune{'X', 'Y', 'Z'}, b: []rune{'W', 'X', 'Y'}, want: []rune{'W', 'X', 'Y', 'Z'}},
		{name: "disjoint", a: []rune{'X', 'Y', 'Z'}, b: []rune{'x', 'y', 'z'}, want: []rune{'X', 'Y', 'Z', 'x', 'y', 'z'}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.Union(b)

			assert.Equal(t, sorted(tt.want), sorted(result.Members()))
		})
	}
}

func TestMutableSet_Intersection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want []rune
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: []rune{}},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: []rune{}},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: []rune{}},
		{name: "identical", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Y', 'Z'}, want: []rune{'X', 'Y', 'Z'}},
		{name: "overlap", a: []rune{'X', 'Y', 'Z'}, b: []rune{'W', 'X', 'Y'}, want: []rune{'X', 'Y'}},
		{name: "disjoint", a: []rune{'X', 'Y', 'Z'}, b: []rune{'x', 'y', 'z'}, want: []rune{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.Intersection(b)

			assert.Equal(t, sorted(tt.want), sorted(result.Members()))
		})
	}
}

func TestMutableSet_Difference(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want []rune
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: []rune{}},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: []rune{'X', 'Y', 'Z'}},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: []rune{}},
		{name: "subtract 1", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X'}, want: []rune{'Y', 'Z'}},
		{name: "subtract 2", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Z'}, want: []rune{'Y'}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.Difference(b)

			assert.Equal(t, sorted(tt.want), sorted(result.Members()))
		})
	}
}

func TestMutableSet_SymmetricDifference(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []rune
		b    []rune
		want []rune
	}{
		{name: "both empty", a: []rune{}, b: []rune{}, want: []rune{}},
		{name: "A + empty", a: []rune{'X', 'Y', 'Z'}, b: []rune{}, want: []rune{'X', 'Y', 'Z'}},
		{name: "empty + B", a: []rune{}, b: []rune{'x', 'y', 'z'}, want: []rune{'x', 'y', 'z'}},
		{name: "identical", a: []rune{'X', 'Y', 'Z'}, b: []rune{'X', 'Y', 'Z'}, want: []rune{}},
		{name: "overlap", a: []rune{'X', 'Y', 'Z'}, b: []rune{'W', 'X', 'Y'}, want: []rune{'W', 'Z'}},
		{name: "disjoint", a: []rune{'X', 'Y', 'Z'}, b: []rune{'x', 'y', 'z'}, want: []rune{'X', 'Y', 'Z', 'x', 'y', 'z'}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := set.NewMutableSet(tt.a...)
			b := set.NewMutableSet(tt.b...)
			result := a.SymmetricDifference(b)

			assert.Equal(t, sorted(tt.want), sorted(result.Members()))
		})
	}
}
