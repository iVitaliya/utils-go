package memory

type stackInstance[T any] struct {
	items []T
}

type stackEntry[T any] struct {
	Index int
	Value T
}

func (s *stackInstance[T]) First() *stackEntry[T] {
	return &stackEntry[T]{
		Index: 0,
		Value: s.items[0],
	}
}

func (s *stackInstance[T]) Last() *stackEntry[T] {
	lastInd := s.Size() - 1

	return &stackEntry[T]{
		Index: lastInd,
		Value: s.items[lastInd],
	}
}

func (s *stackInstance[T]) Push(entry T) {
	s.items = append(s.items, entry)
}

func (s *stackInstance[T]) Pop() T {
	var (
		lastInd   = s.Size() - 1
		removable = s.items[lastInd]
	)

	s.items = s.items[:lastInd]
	return removable
}

func (s *stackInstance[T]) Size() int {
	return len(s.items)
}
