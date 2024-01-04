package slice

import "ecodetools/internal/slice"

func DeleteByIndex[T any](src []T, index int) ([]T, error) {
	return slice.DeleteByIndex(src, index)
}

func FilterDeleteByFunc[T any](src []T, fn func(index int, ele T) bool) []T {
	idx := 0
	for k := range src {
		if fn(k, src[k]) {
			continue
		}
		src[idx] = src[k]
		idx++
	}

	return src[:idx]
}
