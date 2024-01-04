package slice

import (
	"ecodetools/internal/err"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteByIndex(t *testing.T) {
	testClass := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{1, 2, 3},
			index:     0,
			wantSlice: []int{2, 3},
		},
		{
			name:      "index middle",
			slice:     []int{1, 2, 3},
			index:     1,
			wantSlice: []int{1, 3},
		},
		{
			name:      "index length",
			slice:     []int{1, 2, 3},
			index:     2,
			wantSlice: []int{1, 2},
		},
		{
			name:    "index -1",
			slice:   []int{1, 2, 3},
			index:   -1,
			wantErr: err.ErrIndexOutRange(3, -1),
		},
		{
			name:    "index more than length",
			slice:   []int{1, 2, 3},
			index:   3,
			wantErr: err.ErrIndexOutRange(3, 3),
		},
	}

	for _, tc := range testClass {
		t.Run(tc.name, func(t *testing.T) {
			res, err := DeleteByIndex(tc.slice, tc.index)
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
