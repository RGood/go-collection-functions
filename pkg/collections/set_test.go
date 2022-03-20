package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)

	assert.True(t, s.Contains(1))
}
