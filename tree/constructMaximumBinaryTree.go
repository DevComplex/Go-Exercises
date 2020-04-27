package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeHelper(nums, 0, len(nums) - 1)
}

func constructMaximumBinaryTreeHelper(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	maximumIndex := start

	for i := start; i <= end; i++ {
		if nums[maximumIndex] < nums[i] {
			maximumIndex = i
		}
	}

	n := &TreeNode{nums[maximumIndex], nil, nil}
	n.Left = constructMaximumBinaryTreeHelper(nums, start, maximumIndex - 1)
	n.Right = constructMaximumBinaryTreeHelper(nums, maximumIndex + 1, end)
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
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	inorder(root)
}