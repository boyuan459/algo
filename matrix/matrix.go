package matrix

var Directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type Matrix[T any] struct {
	data [][]T
}

func New[T any](data [][]T) *Matrix[T] {
	return &Matrix[T]{
		data: data,
	}
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
