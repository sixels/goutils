package iter

type Iterator[T any] interface {
	Next() (T, bool)
}

func Find[T any](it Iterator[T], fn func(T) bool) (T, bool) {
	for next, ok := it.Next(); ok; next, ok = it.Next() {
		if fn(next) {
			return next, true
		}
	}
	var none T
	return none, false
}

func FindMap[T any, U any](it Iterator[T], fn func(t T) *U) (U, bool) {
	for next, ok := it.Next(); ok; next, ok = it.Next() {
		if r := fn(next); r != nil {
			return *r, true
		}
	}
	var none U
	return none, false
}

func Collect[T any](it Iterator[T]) []T {
	result := make([]T, 0)
	for next, ok := it.Next(); ok; next, ok = it.Next() {
		result = append(result, next)
	}
	return result
}

type Iter struct{}

func (it Iter) Next() (struct{}, bool) {
	return struct{}{}, false
}
