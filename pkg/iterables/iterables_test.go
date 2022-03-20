package iterables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceIterable(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7}

	iterX := NewSliceIterable(x)

	count := 0
	iterX.ForEach(func(i, e int) {
		count++
		assert.Equal(t, e, i)
	})
	assert.Equal(t, len(x), count)
}

func TestMappingIterator(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7}

	iterX := NewSliceIterable(x)
	mappingIter := Map(iterX, func(element int) int {
		return element * 2
	})

	count := 0
	mappingIter.ForEach(func(i, e int) {
		count++
		assert.Equal(t, e, i*2)
	})

	assert.Equal(t, len(x), count)
}

func TestFilter(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5}

	iterX := NewSliceIterable(x)
	filterIter := Filter(iterX, func(element int) bool {
		return element%2 == 0
	})

	count := 0
	filterIter.ForEach(func(i, e int) {
		count++
		assert.True(t, e%2 == 0)
	})

	assert.Equal(t, len(x)/2, count)
}

func TestLast(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6}
	iterX := NewSliceIterable(x)
	lastVal, _ := iterX.Last()

	assert.Equal(t, x[len(x)-1], lastVal)

	emptyIter := NewSliceIterable([]int{})
	_, ok := emptyIter.Last()
	assert.False(t, ok)
}

func TestReduce(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6}
	iterX := NewSliceIterable(x)

	functionalSum := Reduce(iterX, 0, func(total, element int) int {
		return total + element
	})

	imperativeSum := 0
	for _, e := range x {
		imperativeSum += e
	}

	assert.Equal(t, imperativeSum, functionalSum)
}

func TestMapIterable(t *testing.T) {
	x := map[string]string{"foo": "bar", "fizz": "buzz"}
	iterX := NewMapIterable(x)

	count := 0
	iterX.ForEach(func(index int, item *Pair[string, string]) {
		count++
		assert.True(t, x[item.Key] == item.Value)
	})
	assert.Equal(t, 2, count)

	iterX.Reset()

	iterX.ForEach(func(index int, item *Pair[string, string]) {
		count++
		assert.True(t, x[item.Key] == item.Value)
	})

	assert.Equal(t, 4, count)
}

func TestMapIterableReset(t *testing.T) {
	x := map[string]string{"foo": "bar", "fizz": "buzz"}
	iterX := NewMapIterable(x)

	iterX.Next()
	iterX.Reset()

	count := 0
	iterX.ForEach(func(index int, item *Pair[string, string]) {
		count++
		assert.True(t, x[item.Key] == item.Value)
	})

	assert.Equal(t, 2, count)
}
