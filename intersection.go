package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	i := 0
	j := 0

	intersection := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i += 1
		} else if (nums1[i] > nums2[j]) {
			j += 1
		} else {
			intersection = append(intersection, nums1[i])
			i += 1
			j += 1
		}
	}
	
	return intersection
}

func main() {
	nums1 := []int{4, 5, 9}
	nums2 := []int{4, 4, 8, 9, 9}

	out := intersect(nums1, nums2)

	fmt.Println(out)
}