package set_test

import (
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

			s := set.NewImmutableSet(tt.members)
			assert.IsType(t, &set.ImmutableSet[rune]{}, s)
		})
	}
}
