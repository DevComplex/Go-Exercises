package main

import (
	"fmt"
	"math"
)

func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	currOnes := 0
	
	for index, val := range nums {
		if val == 1 {
			currOnes += 1
		} 
		
		if val == 0 || index == len(nums) - 1 {
			maxOnes = int(math.Max(float64(currOnes), float64(maxOnes)))
			currOnes = 0
		}
	}

	return maxOnes
}

func main() {
	input := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(input))
}