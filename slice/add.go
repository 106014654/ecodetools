package slice

import "ecodetools/internal/slice"

func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	return slice.Add[Src](src, element, index)
}
