package main

import "fmt"

type TreeNode struct {
	Val	int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(nums, 0, len(nums) - 1)
}

func sortedArrayToBSTHelper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2

	n := &TreeNode{nums[mid], nil, nil}
	n.Left = sortedArrayToBSTHelper(nums, left, mid - 1)
	n.Right = sortedArrayToBSTHelper(nums, mid + 1, right)

	return n
}

func inorder(n *TreeNode) {
	if n == nil {
		return
	}

	inorder(n.Left)
	fmt.Println(n.Val)
	inorder(n.Right)
}

func main() {
	nums := []int{1, 2, 3, 400, 532, 612, 799, 812}
	root := sortedArrayToBST(nums)
	inorder(root)
}