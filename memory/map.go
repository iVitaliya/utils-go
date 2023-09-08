package memory

import (
	log "github.com/iVitaliya/utils-go/logger"
)

var logger = log.NewLogger()

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

// Does the same as .DeleteAll() but then returns all the deleted key-value's.
func (mi *mapInstance[T]) Clear() []*mapEntry[T] {
	var arr []*mapEntry[T]

	for key, val := range *mi {
		arr = append(arr, &mapEntry[T]{
			Key:   key,
			Value: val.value,
		})

		delete(*mi, key)
	}

	logger.Warn("All the present keys located in the Map have been deleted by using the Clear() method")

	return arr
}

// Deletes an item from the map and returns if it was successfully deleted or not.
func (mi *mapInstance[T]) Delete(key string) bool {
	if !mi.Has(key) {
		return false
	}

	delete(*mi, key)
	return true
}

// Deletes all items one-by-one from the map.
func (mi *mapInstance[T]) DeleteAll() {
	for key, _ := range *mi {
		mi.Delete(key)
	}
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

func (mi *mapInstance[T]) Get(key string) T {
	for k, val := range *mi {
		if k == key {
			return val.value
		}
	}

	logger.Error("Couldn't find the key \"" + key + "\" in the Map, are you sure you provided the right key and didn't make a typo?")
	return *new(T)
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

	return arr
}
