package main

import "fmt"

func oddCells(n int, m int, indicies [][]int) int {
	rows := make(map[int]int)
	cols := make(map[int]int)

	for i := 0; i < len(indicies); i++ {
		indicie := indicies[i]

		row := indicie[0]
		col := indicie[1]

		rows[row]++
		cols[col]++
	}

	oddCellCount := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := rows[i] + cols[j]

			if count % 2 == 1 {
				oddCellCount++
			}
		}
	}

	return oddCellCount
}

func main() {
	n := 2
	m := 3
	
	indicies := [][]int{
		[]int{0, 1},
		[]int{1, 1},
	}

	fmt.Println(oddCells(n, m, indicies))
}