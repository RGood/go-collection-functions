package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(75)
	s.Add(69)

	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(75))
	assert.True(t, s.Contains(69))
	assert.False(t, s.Contains(99))
}

func TestOrderedSet(t *testing.T) {
	s := NewOrderedSet[int]()
	s.Add(8)
	s.Add(1)
	s.Add(63)
	s.Add(42)

	s.Remove(63)

	s.Add(17)

	entries := []int{}
	s.ForEach(func(_, entry int) {
		entries = append(entries, entry)
	})

	assert.Equal(t, entries, []int{8, 1, 42, 17})
}

func TestSetCopy(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(10)

	c := s.Copy()

	s.Remove(1)

	assert.False(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(10))

	assert.True(t, c.Contains(1))
	assert.True(t, c.Contains(2))
	assert.True(t, c.Contains(10))
}
