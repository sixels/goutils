package iter_test

import (
	"math/rand"
	"testing"

	"github.com/manekani/goutils/iter"
	"github.com/manekani/goutils/slices_ext"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	assert := assert.New(t)

	s := []int{1, 2, 3, 4, 5, 10}
	it := iter.SliceIt(s)

	next, _ := it.Next()
	assert.Equal(1, next)
	next, _ = it.Next()
	assert.Equal(2, next)
	next, _ = it.Next()
	assert.Equal(3, next)
	next, _ = it.Next()
	assert.Equal(4, next)
	next, _ = it.Next()
	assert.Equal(5, next)
	next, _ = it.Next()
	assert.Equal(10, next)

	_, ok := it.Next()
	assert.False(ok)
	_, ok = it.Next()
	assert.False(ok)
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	s := []int{1, 2, 3, 4, 5, 10}
	it := iter.MapIt[int](iter.SliceIt(s), func(v int) int {
		return v * 2
	})

	next, _ := it.Next()
	assert.Equal(2, next)
	next, _ = it.Next()
	assert.Equal(4, next)
	next, _ = it.Next()
	assert.Equal(6, next)
	next, _ = it.Next()
	assert.Equal(8, next)
	next, _ = it.Next()
	assert.Equal(10, next)
	next, _ = it.Next()
	assert.Equal(20, next)

	_, ok := it.Next()
	assert.False(ok)
	_, ok = it.Next()
	assert.False(ok)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	s := []int{1, 2, 3, 4, 5, 10}
	it := iter.FilterIt[int](iter.SliceIt(s), func(v int) bool {
		return v%2 == 0
	})

	next, _ := it.Next()
	assert.Equal(2, next)
	next, _ = it.Next()
	assert.Equal(4, next)
	next, _ = it.Next()
	assert.Equal(10, next)

	_, ok := it.Next()
	assert.False(ok)
	_, ok = it.Next()
	assert.False(ok)
}

func TestFilterMap(t *testing.T) {
	assert := assert.New(t)
	s := []int{1, 2, 3, 4, 5, 10}
	it := iter.FilterMapIt[int](iter.SliceIt(s), func(v int) *int {
		if v%2 == 0 {
			v2 := v*2 + 1
			return &v2
		}
		return nil
	})

	next, _ := it.Next()
	assert.Equal(5, next)
	next, _ = it.Next()
	assert.Equal(9, next)
	next, _ = it.Next()
	assert.Equal(21, next)

	_, ok := it.Next()
	assert.False(ok)
	_, ok = it.Next()
	assert.False(ok)
}

type RandIter struct {
	size int
	i    int
}

func (it *RandIter) Next() (float64, bool) {
	it.i += 1
	if it.i < it.size {
		return rand.Float64(), true
	}
	return 0, false
}

func BenchmarkIterMap(b *testing.B) {
	rand.Seed(int64(b.N))
	it := &RandIter{size: 1_000_000_000}
	_ = iter.Collect[float64](iter.MapIt[float64](it, func(v float64) float64 {
		return v * 2
	}))
}

func BenchmarkSliceMap(b *testing.B) {
	rand.Seed(int64(b.N))
	slice := make([]float64, 1_000_000_000)
	for i := 0; i < 1_000_000_000; i += 1 {
		slice[i] = rand.Float64()
	}
	_ = slices_ext.Map(slice, func(v float64) float64 {
		return v * 2
	})
}
