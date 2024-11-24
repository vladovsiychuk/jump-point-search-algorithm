package main

import (
	"container/heap"
	"math"
)

type Position struct {
	X, Y int
}

type Node struct {
	Position          Position
	G, H, F           int
	ParentDir         [2]int
	ForcedNeighborDir [2]int
	Parent            *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func isWalkable(x, y int, matrix [][]int) bool {
	return x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0]) && matrix[x][y] == 0
}

func heuristic(a, b *Position) int {
	// Using Manhattan distance as heuristic
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func findPathWithJPS(start, end Position, matrix [][]int) []Position {
	startNode := &Node{Position: start}
	endNode := &Node{Position: end}
	openList := &PriorityQueue{}
	heap.Init(openList)
	heap.Push(openList, startNode)

	for openList.Len() > 0 {
		currentNode := heap.Pop(openList).(*Node)

		if currentNode.Position == endNode.Position {
			return reconstructPath(currentNode)
		}

		successors := findSuccessors(currentNode, endNode, matrix)
		for _, successor := range successors {
			heap.Push(openList, successor)
		}
	}

	return nil
}

func reconstructPath(node *Node) []Position {
	var path []Position

	// Collect nodes from end to start
	for node != nil {
		path = append(path, node.Position)
		node = node.Parent
	}

	// Reverse the path to get the correct order
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func findSuccessors(currentNode *Node, endNode *Node, matrix [][]int) []*Node {
	var successors []*Node
	x, y := currentNode.Position.X, currentNode.Position.Y
	directions := getDirections(currentNode)

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		sx, sy, forcedNeighborDir, found := jump(x+dx, y+dy, dx, dy, endNode, matrix)
		if found {
			successor := &Node{Position: Position{X: sx, Y: sy}}
			successor.G = currentNode.G + heuristic(&currentNode.Position, &successor.Position)
			successor.H = heuristic(&successor.Position, &endNode.Position)
			successor.F = successor.G + successor.H
			successor.ParentDir = dir
			successor.ForcedNeighborDir = forcedNeighborDir
			successor.Parent = currentNode
			successors = append(successors, successor)
		}
	}

	return successors
}

func getDirections(node *Node) [][2]int {
	if node.ParentDir == [2]int{0, 0} {
		return [][2]int{
			{0, -1}, {0, 1}, {-1, 0}, {1, 0},
			{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
		}
	}

	pdx, pdy := node.ParentDir[0], node.ParentDir[1]
	fndx, fndy := node.ForcedNeighborDir[0], node.ForcedNeighborDir[1]

	if pdx != 0 && pdy != 0 {
		directions := [][2]int{{pdx, pdy}, {pdx, 0}, {0, pdy}}
		if fndx != 0 || fndy != 0 {
			directions = append(directions, [2]int{fndx, fndy})
		}
		return directions
	}

	directions := [][2]int{{pdx, pdy}}
	if fndx != 0 || fndy != 0 {
		directions = append(directions, [2]int{fndx, fndy})
	}
	return directions
}

func jump(x, y, dx, dy int, endNode *Node, matrix [][]int) (int, int, [2]int, bool) {
	for isWalkable(x, y, matrix) {
		if x == endNode.Position.X && y == endNode.Position.Y {
			return x, y, [2]int{}, true
		}

		if dx != 0 && dy != 0 {
			if !isWalkable(x, y+dy*-1, matrix) && isWalkable(x+dx, y+dy*-1, matrix) {
				return x, y, [2]int{dx, dy * -1}, true
			}
			if !isWalkable(x+dx*-1, y, matrix) && isWalkable(x+dx*-1, y+dy, matrix) {
				return x, y, [2]int{dx * -1, dy}, true
			}

			_, _, _, found1 := jump(x+dx, y, dx, 0, endNode, matrix)
			_, _, _, found2 := jump(x, y+dy, 0, dy, endNode, matrix)

			if found1 || found2 {
				return x, y, [2]int{}, true
			}
		} else if dx != 0 {
			if !isWalkable(x, y-1, matrix) && isWalkable(x+dx, y-1, matrix) {
				return x, y, [2]int{dx, -1}, true
			}
			if !isWalkable(x, y+1, matrix) && isWalkable(x+dx, y+1, matrix) {
				return x, y, [2]int{dx, 1}, true
			}
		} else if dy != 0 {
			if !isWalkable(x-1, y, matrix) && isWalkable(x-1, y+dy, matrix) {
				return x, y, [2]int{-1, dy}, true
			}
			if !isWalkable(x+1, y, matrix) && isWalkable(x+1, y+dy, matrix) {
				return x, y, [2]int{1, dy}, true
			}
		}

		x += dx
		y += dy
	}

	return 0, 0, [2]int{}, false
}
