package floodfill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloodFill(t *testing.T) {

	tests := []struct {
		name     string
		image    Image
		startRow int
		startCol int
		color    int
		want     Image
	}{
		{
			name: "Case 1",
			image: Image{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			startRow: 1,
			startCol: 1,
			color:    2,
			want: Image{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "Case 2",
			image: Image{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			startRow: 0,
			startCol: 0,
			color:    0,
			want: Image{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloodFill(tt.image, tt.startRow, tt.startCol, tt.color)
			assert.Equal(t, tt.want, got)
		})
	}
}
