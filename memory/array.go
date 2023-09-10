package memory

import "fmt"

type arrayInstance[T any] []T
type arrayEntry[T any] struct {
	Index int
	Value T
}
type arrayDuplicates struct {
	Count         int
	HasDuplicates bool
	Indexs        []int
}

func NewArray[T any]() arrayInstance[T] {
	return arrayInstance[T]{}
}

func (a *arrayInstance[T]) At(index int) arrayEntry[T] {
	if index < 0 {
		return a.First()
	} else if index >= a.Length() {
		return a.Last()
	} else {
		return arrayEntry[T]{
			Index: index,
			Value: (*a)[index],
		}
	}
}

func (a *arrayInstance[T]) Append(entry T) arrayInstance[T] {

}

// Clear all the items from the Array
func (a *arrayInstance[T]) Clear() {
	*a = arrayInstance[T]{}
}

// Returns a boolean defining whether a deletion of specified item has succeeded or not
func (a *arrayInstance[T]) Delete(index int) {
	arr := NewArray[T]()

	if index < 0 || index >= a.Length() {
		return
	}

	// Copy the elements before the index
	copy(arr, (*a)[:index])

	// Copy the elements after the index
	copy(arr[index:], (*a)[index+1:])

	// Re-assign the array with the new array
	*a = arr
}

// Retrieves the First item defined in the array.
func (a *arrayInstance[T]) First() arrayEntry[T] {
	return arrayEntry[T]{
		Index: 0,
		Value: (*a)[0],
	}
}

func (a *arrayInstance[T]) hasDuplicates(key T) arrayDuplicates {
	var (
		key_str                    = fmt.Sprint(key)
		count   int                = 0
		hasDups bool               = false
		indexs  arrayInstance[int] = NewArray[int]()
	)

	for ind, val := range *a {
		val_str := fmt.Sprint(val)

		if val_str == key_str {
			indexs.Append(ind)
			count = count + 1
		}
	}

	if count > 0 {
		hasDups = true
	}

	return arrayDuplicates{
		Count:         count,
		HasDuplicates: hasDups,
		Indexs:        indexs,
	}
}

// Retrieves the last item defined in the array.
func (a *arrayInstance[T]) Last() arrayEntry[T] {
	var (
		ind = len(*a) - 1
		val = (*a)[ind]
	)

	return arrayEntry[T]{
		Index: ind,
		Value: val,
	}
}

// Gets the length/size of the array.
func (a *arrayInstance[T]) Length() int {
	return len(*a)
}

// Pushes an item into the array, like append but then an alias.
func (a *arrayInstance[T]) Push(entries ...T) arrayInstance[T] {
	for _, val := range entries {
		a.Append(val)
	}

	return *a
}
