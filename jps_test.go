package main

import (
	"testing"
)

func TestJumpPointSearch(t *testing.T) {
	gridXY := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	matrix := make([][]int, len(gridXY))
	for y := 0; y < len(gridXY); y++ {
		column := make([]int, len(gridXY[0]))
		for x := 0; x < len(gridXY[y]); x++ {
			column[x] = gridXY[x][y]
		}
		matrix[y] = column
	}

	start := &Node{Position: Position{X: 0, Y: 2}}
	end := &Node{Position: Position{X: 2, Y: 0}}

	path := findPathWithJPS(start, end, matrix)

	if len(path) == 0 {
		t.Errorf("No path found from start (%d, %d) to end (%d, %d)", start.Position.X, start.Position.Y, end.Position.X, end.Position.Y)
	}

	for _, node := range path {
		t.Logf("Node Position: (%d, %d)", node.Position.X, node.Position.Y)
	}
}
