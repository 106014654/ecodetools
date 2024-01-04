package slice

func DiffSet[T comparable](src, dst []T) []T {
	srcMap := makeMap[T](src)
	for _, val := range dst {
		delete(srcMap, val)
	}

	var ret = make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}

	return ret
}
