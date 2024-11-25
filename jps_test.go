package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpPointSearch(t *testing.T) {
	tests := []struct {
		name         string
		gridXY       [][]int
		start        Position
		end          Position
		expectedPath []Position
	}{
		{
			name: "test 1",
			gridXY: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			start: Position{X: 0, Y: 2},
			end:   Position{X: 2, Y: 0},
			expectedPath: []Position{
				{X: 0, Y: 2},
				{X: 1, Y: 2},
				{X: 2, Y: 1},
				{X: 2, Y: 0},
			},
		},
		{
			name: "test 2",
			gridXY: [][]int{
				{0, 0, 0},
				{1, 0, 0},
				{0, 1, 0},
			},
			start:        Position{X: 0, Y: 2},
			end:          Position{X: 2, Y: 0},
			expectedPath: nil,
		},
		{
			name: "test 3",
			gridXY: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			start: Position{X: 2, Y: 7},
			end:   Position{X: 6, Y: 3},
			expectedPath: []Position{
				{X: 2, Y: 7},
				{X: 3, Y: 8},
				{X: 4, Y: 8},
				{X: 5, Y: 7},
				{X: 6, Y: 6},
				{X: 6, Y: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// flip the grid to matrix
			matrix := make([][]int, len(tt.gridXY))
			for y := 0; y < len(tt.gridXY); y++ {
				column := make([]int, len(tt.gridXY[0]))
				for x := 0; x < len(tt.gridXY[y]); x++ {
					column[x] = tt.gridXY[x][y]
				}
				matrix[y] = column
			}

			path := findPathWithJPS(tt.start, tt.end, matrix)

			assert.Equal(t, tt.expectedPath, path)
		})
	}
}
