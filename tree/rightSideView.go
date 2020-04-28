package main

import "fmt"

type TreeNode struct {
	Val int	
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(n *TreeNode) []int {
	visitedLevels := make(map[int]bool)
	view := []int{}
	rightSideViewHelper(n, 0, &view, visitedLevels)
	return view
}

func rightSideViewHelper(n *TreeNode, level int, view *[]int, visitedLevels map[int]bool) {
	if n == nil {
		return
	}

	if !visitedLevels[level] {
		*view = append(*view, n.Val)
		visitedLevels[level] = true
	}

	rightSideViewHelper(n.Right, level + 1, view, visitedLevels)
	rightSideViewHelper(n.Left, level + 1, view, visitedLevels)
}


func main() {
	r := &TreeNode{1, nil, nil}
	r.Left = &TreeNode{2, nil, nil}
	r.Left.Right = &TreeNode{5, nil, nil}
	r.Right = &TreeNode{3, nil, nil}
	r.Right.Right = &TreeNode{4, nil, nil}

	output := rightSideView(r)

	fmt.Println(output)
}