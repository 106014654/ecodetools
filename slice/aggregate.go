package slice

import "ecodetools"

func Max[T ecodetools.RealNumber](src []T) T {
	res := src[0]

	for i := 1; i < len(src); i++ {
		if src[i] > res {
			res = src[i]
		}
	}
	return res
}

func Min[T ecodetools.RealNumber](src []T) T {
	res := src[0]

	for i := 1; i < len(src); i++ {
		if src[i] < res {
			res = src[i]
		}
	}
	return res
}

func Sum[T ecodetools.RealNumber](src []T) T {
	res := src[0]

	for i := 1; i < len(src); i++ {
		res += src[i]
	}
	return res
}
