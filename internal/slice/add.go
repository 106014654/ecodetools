package slice

import (
	"ecodetools/internal/err"
)

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, err.ErrIndexOutRange(length, index)
	}

	var zeroVal T
	src = append(src, zeroVal)

	for i := length; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}
