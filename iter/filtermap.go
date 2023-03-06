package iter

type FilterMap[T any, U any] struct {
	it Iterator[T]
	fn func(T) *U
}

func FilterMapIt[T any, U any](it Iterator[T], fn func(t T) *U) FilterMap[T, U] {
	return newFilterMap(it, fn)
}

func newFilterMap[T any, U any](it Iterator[T], fn func(T) *U) FilterMap[T, U] {
	return FilterMap[T, U]{
		it,
		fn,
	}
}

func (f FilterMap[T, U]) Next() (U, bool) {
	return FindMap(f.it, f.fn)
}
