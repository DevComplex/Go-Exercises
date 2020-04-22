package main

import "fmt"

func sumZero(n int) []int {
	sum := 0

	nums := []int{}

	for i := 0; i < n - 1; i++ {
		nums = append(nums, i)
		sum += i
	}

	var neg int

	if sum == 0 {
		neg = 0
	} else {
		neg = -1 * sum
	}

	nums = append(nums, neg)

	return nums
}

func main() {
	n := 1000

	fmt.Println(sumZero(n))
}