package memory

type setInstance[T any] map[string]T

// !!!TODO: If item already present, give the parameter to override it
func (set *setInstance[T]) Add(key string, value T) {
	a.removeDuplicates(entry)
}

func (set *setInstance[T]) removeDuplicates(key T) {
	duplicates := set.hasDuplicates(key)

	if duplicates.HasDuplicates {
		for _, i := range duplicates.Indexs {
			set.Delete(i)
		}
	}
}
