package slice

import "ecodetools/internal/err"

func DeleteByIndex[T any](src []T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, err.ErrIndexOutRange(length, index)
	}

	for k, _ := range src {
		if k >= index && k < length-1 {
			src[k] = src[k+1]
		}
	}
	src = src[:length-1]
	return src, nil
}
