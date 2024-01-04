package slice

import (
	"ecodetools/internal/err"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSlice(t *testing.T) {
	testClass := []struct {
		name      string
		slice     []int
		element   int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{1, 2, 3},
			element:   0,
			index:     0,
			wantSlice: []int{0, 1, 2, 3},
		},
		{
			name:      "index middl1",
			slice:     []int{1, 2, 3},
			element:   0,
			index:     1,
			wantSlice: []int{1, 0, 2, 3},
		},
		{
			name:      "index equal length",
			slice:     []int{1, 2, 3},
			element:   0,
			index:     3,
			wantSlice: []int{1, 2, 3, 0},
		},
		{
			name:    "index less than 0",
			slice:   []int{1, 2, 3},
			element: 0,
			index:   -1,
			wantErr: err.ErrIndexOutRange(3, -1),
		},
		{
			name:    "index out range",
			slice:   []int{1, 2, 3},
			element: 0,
			index:   4,
			wantErr: err.ErrIndexOutRange(3, 4),
		},
	}

	for _, tc := range testClass {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.element, tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
