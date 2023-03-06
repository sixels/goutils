package iter

type Map[T any, U any] struct {
	it Iterator[T]
	fn func(T) U
}

func MapIt[T any, U any](it Iterator[T], fn func(t T) U) *Map[T, U] {
	return newMap(it, fn)
}

func newMap[T any, U any](i Iterator[T], f func(T) U) *Map[T, U] {
	return &Map[T, U]{
		i, f,
	}
}

func (m *Map[T, U]) Next() (U, bool) {
	next, ok := m.it.Next()
	if !ok {
		var none U
		return none, false
	}
	return m.fn(next), true
}
