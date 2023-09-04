package memory

type mapInstance[T any] map[string]struct {
	value T
	index int
}
type mapEntry[T any] struct {
	Key   string
	Value T
}

func NewMap[T any]() mapInstance[T] {
	instance := make(map[string]struct {
		value T
		index int
	})

	return instance
}

func (mi *mapInstance[T]) At(index int) T {
	var (
		ind     = index
		lastInd = len(*mi)
	)

	if ind < 0 {
		ind = 0
	} else if ind > lastInd {
		ind = lastInd
	}

	for _, v := range *mi {
		if v.index == ind {
			return v.value
		}
	}

	return *new(T)
}

func (mi *mapInstance[T]) Find(predicate func(element T, index int, iterator mapInstance[T]) T) T {
	for _, v := range *mi {
		return predicate(v.value, v.index, *mi)
	}

	return *new(T)
}

func (mi *mapInstance[T]) Every(predicate func(element T, iterator mapInstance[T]) bool) bool {
	for _, v := range *mi {
		if predicate(v.value, *mi) {
			return true
		}
	}

	return false
}

func (mi *mapInstance[T]) Size() int {
	return len(*mi)
}
