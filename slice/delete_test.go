package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleterbyFunc(t *testing.T) {
	testClass := []struct {
		name       string
		src        []int
		deleteFunc func(index int, ele int) bool
		wantSlice  []int
	}{
		{
			name: "delete first",
			src:  []int{1, 2, 3, 4, 5},
			deleteFunc: func(index int, ele int) bool {
				return index == 0
			},
			wantSlice: []int{2, 3, 4, 5},
		},
		{
			name: "delete last",
			src:  []int{1, 2, 3, 4, 5},
			deleteFunc: func(index int, ele int) bool {
				return index == 4
			},
			wantSlice: []int{1, 2, 3, 4},
		},
		{
			name: "delete 重复",
			src:  []int{1, 2, 3, 3, 4, 5},
			deleteFunc: func(index int, ele int) bool {
				return ele == 3
			},
			wantSlice: []int{1, 2, 4, 5},
		},
		{
			name: "delete all",
			src:  []int{1, 2, 3, 3, 4, 5},
			deleteFunc: func(index int, ele int) bool {
				return true
			},
			wantSlice: []int{},
		},
	}

	for _, tc := range testClass {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterDeleteByFunc(tc.src, tc.deleteFunc)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
