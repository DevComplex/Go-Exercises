package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	n := TreeNode{10, nil, nil}
	n.Left = &TreeNode{20, nil, nil}
	n.Left.Left = &TreeNode{30, nil, nil}
	n.Right = &TreeNode{0, nil, nil}

	fmt.Println(findTarget(&n, 30))
	fmt.Println(findTarget(&n, 33))
}

func findTarget(root *TreeNode, k int) bool {
	nums := make(map[int]bool)
	return findTargetHelper(root, k, nums)

}

func findTargetHelper(n *TreeNode, k int, nums map[int]bool) bool{
	if n == nil {
		return false
	}

	newTarget := k - n.Val

	if nums[newTarget] == true {
		return true
	}

	nums[n.Val] = true

	return findTargetHelper(n.Left, k, nums) || findTargetHelper(n.Right, k, nums)
}