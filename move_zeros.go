package main

import "fmt"

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func moveZeros(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			swap(nums, i, j)
			j++
		}
	}
}

func main() {
	input := []int{0, 1, 0, 3, 12}
	moveZeros(input)
	fmt.Println(input)
}