package memory

type mapInstance[T any] map[string]struct {
	index int
	value T
}
type mapEntry[T any] struct {
	Key   string
	Value T
}

func (mi *mapInstance[T]) At(index int) T {

}

func (mi *mapInstance[T]) Find(predicate func(element T, index int, iterator mapInstance[T]) T) T {

}

func (mi *mapInstance[T]) Every(predicate func(element T, iterator mapInstance[T])) {
	for i := 0; i < mi.Size(); i++ {
		
	}
}

func (mi *mapInstance[T]) Size() int {
	return len(*mi)
}
