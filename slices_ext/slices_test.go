package slices_ext_test

import (
	"testing"

	"github.com/manekani/goutils/slices_ext"
)

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(v int) int {
		return v * 2
	}
	result := slices_ext.Map(s, f)
	if len(result) != len(s) {
		t.Errorf("Mapped slice length doesn't match original slice length. Expected: %d, got: %d",
			len(s), len(result))
	}
	for i, v := range result {
		if v != f(s[i]) {
			t.Errorf("Mapped slice doesn't match original slice. Expected: %d, got: %d (index: %d)",
				s[i], v, i)
		}
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	f := func(v int) bool {
		return v%2 == 0
	}
	result := slices_ext.Filter(s, f)
	if len(result) != 3 {
		t.Errorf("Filtered slice length doesn't match the expected length. Expected: %d, got: %d",
			3, len(result))
	}
	for i, v := range result {
		if !f(v) {
			t.Errorf("Filtered slice have an invalid value at index: %d", i)
		}
	}
}

func TestFilterMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	m := func(v int) int {
		return v * 3
	}
	f := func(v int) *int {
		if v%2 == 0 {
			r := m(v)
			return &r
		} else {
			return nil
		}
	}

	result := slices_ext.FilterMap(s, f)
	if len(result) != 3 {
		t.Errorf("FilterMapped slice length doesn't match the expected length. Expected: %d, got: %d",
			3, len(result))
	}
	// for i, _ := range result {
	// 	if f(s[i]) == nil {
	// 	}

	// }
}
