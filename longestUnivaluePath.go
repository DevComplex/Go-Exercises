package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUniValuePath(n *TreeNode) int {
	max := 0
	longestUniValuePathHelper(n, &max, nil, 0)
	return max
}

func longestUniValuePathHelper(n *TreeNode, max *int, prev *int, upToHere int) {
	if n == nil {
		return
	}

	soFar := 1

	if prev != nil && n.Val == *prev {
		soFar += upToHere
	}

	*max = int(math.Max(float64(soFar), float64(*max)))

	longestUniValuePathHelper(n.Left, max, &n.Val, soFar)
	longestUniValuePathHelper(n.Right, max, &n.Val, soFar)
}

func main() {
	n := &TreeNode{5, nil, nil}
	n.Left = &TreeNode{5, nil, nil}
	n.Left.Left = &TreeNode{5, nil, nil}

	n.Right = &TreeNode{5, nil, nil}
	n.Right.Right = &TreeNode{5, nil, nil}
	n.Right.Right.Right = &TreeNode{5, nil, nil}

	res := longestUniValuePath(n)

	fmt.Println(res)
}
