package main

import "fmt"

func plusOne(digits []int) []int {
	out := []int{}

	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		carry += digits[i]

		digit := carry % 10
		carry = carry / 10

		out = append([]int{digit}, out...)
	}

	if carry > 0 {
		out = append([]int{carry}, out... )
	}

	return out
}

func main() {
	input := []int{9, 9, 9, 9}
	out := plusOne(input)

	fmt.Println(out)
}