package slice

func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := makeMap(src)
	var ret = make([]T, 0, len(src))

	for _, val := range dst {
		if _, exist := srcMap[val]; exist {
			ret = append(ret, val)
		}
	}
	return deduplicate[T](ret)
}
