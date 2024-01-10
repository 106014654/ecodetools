package slice

func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	data := make([]Dst, len(src))
	for k, v := range src {
		data[k] = m(k, v)
	}
	return data
}

func makeMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		// 使用空结构体,减少内存消耗
		dataMap[v] = struct{}{}
	}
	return dataMap
}

func deduplicate[T comparable](data []T) []T {
	dataMap := makeMap[T](data)
	var newData = make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}
