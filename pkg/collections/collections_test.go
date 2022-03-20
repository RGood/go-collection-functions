package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	elements := []bool{false, false, false, true, false}
	assert.True(t, Any(elements))

	elements = []bool{false, false, false, false, false}
	assert.False(t, Any(elements))
}

func TestAll(t *testing.T) {
	elements := []bool{true, true, true, true, true}
	assert.True(t, All(elements))

	elements = []bool{true, true, false, true, true}
	assert.False(t, All(elements))
}

func TestMapperFunction(t *testing.T) {
	elements := []int{1, 2, 3}
	mappedElements := Map(elements, func(input int) int {
		return input * 2
	})

	for i, e := range elements {
		assert.Equal(t, mappedElements[i], 2*e)
	}
}

func TestReduceFunction(t *testing.T) {
	elements := []int{1, 3, 5, 9}
	sum := Reduce(0, elements, func(total, element int) int {
		return total + element
	})

	assert.Equal(t, sum, 18)
}

func TestFilterFunction(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filteredElements := Filter(elements, func(e int) bool {
		return e%2 == 0
	})

	// All of our filtered elements should be even
	assert.True(t, All(Map(filteredElements, func(e int) bool {
		return e%2 == 0
	})))

	// We have removed half of our elements
	assert.Equal(t, len(filteredElements), len(elements)/2)

	// Not all of our initial elements were even
	assert.False(t, All(Map(elements, func(e int) bool {
		return e%2 == 0
	})))
}
