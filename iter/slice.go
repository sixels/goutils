package iter

type SliceIter[T any] struct {
	slice []T
}

func SliceIt[T any](slice []T) *SliceIter[T] {
	return &SliceIter[T]{
		slice: slice,
	}
}

func (it *SliceIter[T]) Next() (T, bool) {
	if len(it.slice) == 0 {
		var none T
		return none, false
	}

	next := it.slice[0]
	it.slice = it.slice[1:]

	return next, true
}
