package memory

import (
	log "github.com/iVitaliya/utils-go/logger"
)

var logger = log.NewLogger()

type mapInstance[T any] map[string]T

type mapEntry[T any] struct {
	Key   string
	Value T
}

type mapEntries[T any] struct {
	Next struct {
		Done  bool
		Value mapEntry[T]
	}
	Return []mapEntry[T]
}

func NewMap[T any]() mapInstance[T] {
	instance := make(map[string]T)

	return instance
}

// Does the same as .DeleteAll() but then returns all the deleted key-value's.
func (mi *mapInstance[T]) Clear() []mapEntry[T] {
	var arr []mapEntry[T]

	for key, val := range *mi {
		arr = append(arr, mapEntry[T]{
			Key:   key,
			Value: val,
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

func (mi *mapInstance[T]) Entries() arrayInstance[struct {
	Key   string
	Value T
}] {
	arr := NewArray[struct {
		Key   string
		Value T
	}]()

	for key, val := range *mi {
		arr.Append(struct {
			Key   string
			Value T
		}{
			Key:   key,
			Value: val,
		})
	}

	return arr
}

func (mi *mapInstance[T]) Every(predicate func(element T, iterator mapInstance[T]) bool) bool {
	for _, val := range *mi {
		if predicate(val, *mi) {
			return true
		}
	}

	return false
}

func (mi *mapInstance[T]) Find(predicate func(element T, iterator mapInstance[T]) T) T {
	for _, val := range *mi {
		return predicate(val, *mi)
	}

	return *new(T)
}

func (mi *mapInstance[T]) ForEach(predicate func(element T, iterator mapInstance[T])) {
	for _, val := range *mi {
		predicate(val, *mi)
	}
}

func (mi *mapInstance[T]) Get(key string) mapEntry[T] {
	for k, val := range *mi {
		if k == key {
			return mapEntry[T]{
				Key:   key,
				Value: val,
			}
		}
	}

	logger.Error("Couldn't find the key \"" + key + "\" in the Map, are you sure you provided the right key and didn't make a typo?")

	return *new(mapEntry[T])
}

func (mi *mapInstance[T]) Has(key string) bool {
	for k, _ := range *mi {
		if k == key {
			return true
		}
	}

	return false
}

func (mi *mapInstance[T]) Keys() arrayInstance[string] {
	arr := NewArray[string]()

	for key, _ := range *mi {
		arr.Append(key)
	}

	return arr
}

func (mi *mapInstance[T]) Set(key string, value T) {
	(*mi)[key] = value
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

func (mi *mapInstance[T]) Values() arrayInstance[T] {
	arr := NewArray[T]()

	if mi.Size() < 1 {
		logger.Error("Couldn't retrieve any values as the Map is currently empty, so fetching values would be impossible")

		return arr
	}

	for _, val := range *mi {
		arr.Append(val)
	}

	return arr
}
