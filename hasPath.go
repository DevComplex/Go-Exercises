package main

import "fmt"

func hasPath(grid [][]rune) bool {
	rows := len(grid)
	cols := len(grid[0])

	var queue [][]int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				queue = append(queue, []int{i, j})
				break
			}
		}
	}

	if len(queue) == 0 {
		return false
	}

	rowD := [4]int{1, -1, 0, 0}
	colD := [4]int{0, 0, 1, -1}

	for len(queue) > 0 {
		element := queue[0]

		currRow := element[0]
		currCol := element[1]

		if grid[currRow][currCol] == 'D' {
			return true
		}

		for i := 0; i < 4; i++ {
			nextRow := currRow + rowD[i]
			nextCol := currCol + colD[i]

			if nextRow < 0 || nextCol < 0 || nextRow == rows || nextCol == cols {
				continue
			}

			if grid[nextRow][nextCol] == 'W' || grid[nextRow][nextCol] == '@' {
				continue
			}

			queue = append(queue, []int{nextRow, nextCol})
		}

		grid[currRow][currCol] = '@'

		queue = queue[1:]
	}

	return false
}

func main() {
	grid := [][]rune{
		{'.', '.', '.', 'D'},
		{'.', 'W', '.', '.'},
		{'S', 'W', '.', '.'},
		{'.', 'W', '.', '.'},
	}

	fmt.Println(hasPath(grid))
}