package memory

type instances[T any] struct {
	Array arrayInstance[T]
}

func NewInstance[T any]() instances[T] {
	return instances[T]{
		Array: newArrayInstance[T](),
	}
}
