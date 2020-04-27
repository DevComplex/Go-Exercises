package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val	int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	maxOfTwo := int(math.Max(float64(maxDepth(n.Left)), float64(maxDepth(n.Right))))
	return maxOfTwo + 1
}

func main() {
	n := &TreeNode{3, nil, nil}
	n.Left =  &TreeNode{9, nil, nil}
	n.Right =  &TreeNode{20, nil, nil}
	n.Right.Right =  &TreeNode{7, nil, nil}
	n.Right.Left =  &TreeNode{15, nil, nil}

	fmt.Println(maxDepth(n))
}