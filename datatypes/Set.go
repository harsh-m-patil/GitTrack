package datatypes

type Set[T comparable] map[T]bool

func (s Set[T]) Add(value T) {
	s[value] = true
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

func newSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))

	for val := range s {
		slice = append(slice, val)
	}
	return slice
}
