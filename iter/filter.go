package iter

type Filter[T any] struct {
	it Iterator[T]
	fn func(T) bool
}

func FilterIt[T any](it Iterator[T], fn func(t T) bool) Filter[T] {
	return newFilter(it, fn)
}

func newFilter[T any](it Iterator[T], fn func(T) bool) Filter[T] {
	return Filter[T]{
		it,
		fn,
	}
}

func (f Filter[T]) Next() (T, bool) {
	return Find(f.it, f.fn)
}
