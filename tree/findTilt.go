package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	sum := 0
	findTiltHelper(root, &sum)
	return sum
}

func findTiltHelper(n *TreeNode, sum *int) int {
	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil {
		return n.Value
	}

	left := findTiltHelper(n.Left, sum)
	right := findTiltHelper(n.Right, sum)

	diff := int(math.Abs(float64(left) - float64(right)))

	*sum += diff

	return left + right + n.Value
}

func main() {
	n := &TreeNode{1, nil, nil}
	n.Left = &TreeNode{3, nil, nil}
	n.Right = &TreeNode{2, nil, nil}

	fmt.Println(findTilt(n))
}