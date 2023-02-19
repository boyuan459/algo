package matrix

import (
	"golang.org/x/exp/constraints"
)

var Directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type Matrix[T constraints.Ordered] struct {
	data [][]T
}

func New[T constraints.Ordered](data [][]T) *Matrix[T] {
	return &Matrix[T]{
		data: data,
	}
}

func (matrix Matrix[T]) _dfs(row int, col int, values *[]T, seen *map[[2]int]bool) {
	if row < 0 || row >= len(matrix.data) || col < 0 || col >= len(matrix.data[0]) {
		return
	}
	if _, ok := (*seen)[[2]int{row, col}]; ok {
		return
	}
	*values = append(*values, matrix.data[row][col])
	(*seen)[[2]int{row, col}] = true

	for i := 0; i < len(Directions); i++ {
		var nextRow = row + Directions[i][0]
		var nextCol = col + Directions[i][1]
		matrix._dfs(nextRow, nextCol, values, seen)
	}
}

func (matrix Matrix[T]) DepthFirstTraverse() []T {
	values := []T{}
	seen := make(map[[2]int]bool)

	matrix._dfs(0, 0, &values, &seen)
	return values
}

func (matrix Matrix[T]) BreathFirstTraverse() []T {
	values := []T{}
	queue := [][]int{}
	seen := make(map[[2]int]bool)

	queue = append(queue, []int{0, 0})
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		var row = current[0]
		var col = current[1]

		if row < 0 || row >= len(matrix.data) || col < 0 || col >= len(matrix.data[0]) {
			continue
		}
		if _, ok := seen[[2]int{row, col}]; ok {
			continue
		}

		values = append(values, matrix.data[row][col])
		seen[[2]int{row, col}] = true

		for i := 0; i < len(Directions); i++ {
			var nextRow = row + Directions[i][0]
			var nextCol = col + Directions[i][1]
			queue = append(queue, []int{nextRow, nextCol})
		}
	}

	return values
}

func (matrix Matrix[T]) NumsOfIslands(island T, water T) int {
	var count = 0
	var rows = len(matrix.data)
	var cols = len(matrix.data[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix.data[i][j] == island {
				count += 1
				queue := [][]int{}
				queue = append(queue, []int{i, j})

				for len(queue) > 0 {
					current := queue[0]
					queue = queue[1:]

					var row = current[0]
					var col = current[1]

					if row < 0 || row >= rows || col < 0 || col >= cols || matrix.data[row][col] == water {
						continue
					}
					matrix.data[row][col] = water

					for k := 0; k < len(Directions); k++ {
						var nextRow = row + Directions[k][0]
						var nextCol = col + Directions[k][1]

						queue = append(queue, []int{nextRow, nextCol})
					}
				}
			}
		}
	}

	return count
}
