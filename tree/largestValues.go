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

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func main() {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(3)
	root.Left.Left = NewTreeNode(5)
	root.Left.Right = NewTreeNode(3)
	root.Right = NewTreeNode(2)
	root.Right.Right = NewTreeNode(9)

	fmt.Println(largestValues(root))
}

func largestValues(root *TreeNode) []int {
	maxOfLevels := make(map[int]int)
	largestValuesHelper(root, 0, maxOfLevels)

	output := []int{}

	currGroup := 0

	for {
		_, ok := maxOfLevels[currGroup]

		if !ok {
			break
		}

		output = append(output, maxOfLevels[currGroup])

		currGroup += 1
	}

	return output
}

func largestValuesHelper(n *TreeNode, level int, maxOfLevels map[int]int) {
	if n == nil {
		return
	}

	_, ok := maxOfLevels[level]

	if ok {
		maxOfLevel := math.Max(float64(n.Val), float64(maxOfLevels[level]))
		maxOfLevels[level] = int(maxOfLevel)
	} else {
		maxOfLevels[level] = n.Val
	}

	largestValuesHelper(n.Left, level + 1, maxOfLevels)
	largestValuesHelper(n.Right, level + 1, maxOfLevels)
}