package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1

	m--
	n--

	for i >= 0  && m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}

		i--
	}

	for i >= 0 && m >= 0 {
		nums1[i] = nums1[m]
		i--
		m--
	}

	for i >= 0 && n >= 0 {
		nums1[i] = nums2[n]
		i--
		n--
	}
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}       
	n := 3

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}