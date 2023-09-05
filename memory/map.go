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

// Deletes an item from the map and returns if it was successfully deleted or not.
func (mi *mapInstance[T]) Delete(key string) bool {
	if !mi.Has(key) {
		return false
	}

	delete(*mi, key)
	return true
}

func (mi *mapInstance[T]) Every(predicate func(element T, iterator mapInstance[T]) bool) bool {
	for _, v := range *mi {
		if predicate(v.value, *mi) {
			return true
		}
	}

	return false
}

func (mi *mapInstance[T]) Find(predicate func(element T, index int, iterator mapInstance[T]) T) T {
	for _, v := range *mi {
		return predicate(v.value, v.index, *mi)
	}

	return *new(T)
}

func (mi *mapInstance[T]) ForEach(predicate func(element T, iterator mapInstance[T])) {
	for _, v := range *mi {
		predicate(v.value, *mi)
	}
}

func (mi *mapInstance[T]) Has(key string) bool {
	for k, _ := range *mi {
		if key == k {
			return true
		}
	}

	return false
}

func (mi *mapInstance[T]) Set(key string, value T) {
	m := *mi

	m[key] = struct {
		value T
		index int
	}{
		value: value,
	}
}

func (mi *mapInstance[T]) Size() int {
	return len(*mi)
}

func (mi *mapInstance[T]) ToArray() arrayInstance[T] {
	arr := NewArray[T]()

	mi.ForEach(func(element T, iterator mapInstance[T]) {
		arr.Append(element)
	})

	
}