package slices_ext

// Map returns a new slice containing the results of applying the given function
// to each of the elements of the original slice.
func Map[T any, U any](s []T, f func(v T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only the elements of s for which f returns true.
func Filter[T any](s []T, f func(v T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterMap returns a new slice containing the non nil results of s after
// applying the given function to each element of the original slice.
func FilterMap[T any, U any](s []T, f func(v T) *U) []U {
	result := make([]U, 0)
	for _, v := range s {
		if u := f(v); u != nil {
			result = append(result, *u)
		}
	}
	return result
}

// Reduce applies a function to each element of a slice and returns the reduced
// value.
func Reduce[T any, U any](s []T, f func(init U, v T) U, init U) U {
	for _, v := range s {
		init = f(init, v)
	}
	return init
}
