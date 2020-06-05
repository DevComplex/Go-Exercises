package main

import "fmt"

func singleNumber(nums []int) int {
	singleNum := 0

	for _, num := range nums {
		singleNum = singleNum ^ num
	}

	return singleNum
}

func main() {
	input := []int{4, 1, 2, 1, 2}

	out := singleNumber(input)

	fmt.Println(out)
}