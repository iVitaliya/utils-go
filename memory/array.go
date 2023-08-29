package memory

type arrayInstance[T any] []T
type arrayEntry[T any] struct {
	Index int
	Value T
}

func newArrayInstance[T any]() arrayInstance[T] {
	return arrayInstance[T]{}
}

func (a *arrayInstance[T]) At(index int) arrayEntry[T] {
	if index < 0 {
		return a.First()
	} else if index > a.Length() {
		return a.Last()
	} else {
		arr := *a

		return arrayEntry[T]{
			Index: index,
			Value: arr[index],
		}
	}
}

func (a *arrayInstance[T]) Append(entry T) arrayInstance[T] {

}

// Retrieves the First item defined in the array.
func (a *arrayInstance[T]) First() arrayEntry[T] {
	var (
		arr = *a
		val = arr[0]
	)

	return arrayEntry[T]{
		Index: 0,
		Value: val,
	}
}

func (a *arrayInstance[T]) HasDuplicates(key T) {

}

// Retrieves the last item defined in the array.
func (a *arrayInstance[T]) Last() arrayEntry[T] {
	var (
		arr = *a
		ind = len(arr) - 1
		val = arr[ind]
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