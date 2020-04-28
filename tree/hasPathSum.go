package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumHelper(root, sum, 0)
}

func hasPathSumHelper(root *TreeNode, sum int, currSum int) bool {
	if root == nil {
		return false
	}

	currSum += root.Val

	if root.Left == nil && root.Right == nil {
		return sum == currSum
	}

	return hasPathSumHelper(root.Left, sum, currSum) || hasPathSumHelper(root.Right, sum, currSum)
}

func main() {
	n1 := &TreeNode{5, nil, nil}
	n1.Left = &TreeNode{4, nil, nil}
	n1.Left.Left = &TreeNode{11, nil, nil}
	n1.Left.Left.Left = &TreeNode{7, nil, nil}
	n1.Left.Left.Right = &TreeNode{2, nil, nil}
	n1.Right = &TreeNode{8, nil, nil}
	n1.Right.Left = &TreeNode{13, nil, nil}
	n1.Right.Right = &TreeNode{4, nil, nil}
	n1.Right.Right.Right = &TreeNode{1, nil, nil}

	testCasesThatShouldPass := []int{22, 27, 26, 18}

	testCasesThatShouldFail := []int{0, 1223, 3343}

	for _, testCase := range testCasesThatShouldPass {
		fmt.Println(hasPathSum(n1, testCase))
	}

	for _, testCase := range testCasesThatShouldFail {
		fmt.Println(hasPathSum(n1, testCase))
	}
}