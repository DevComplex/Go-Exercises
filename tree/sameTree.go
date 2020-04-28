package main

import "fmt"

type TreeNode struct {
	Val  int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	n1 := &TreeNode{1, nil, nil}
	n1.Left = &TreeNode{2, nil, nil}
	n1.Right = &TreeNode{3, nil, nil}

	n2 := &TreeNode{1, nil, nil}
	n2.Left = &TreeNode{2, nil, nil}
	n2.Right = &TreeNode{3, nil, nil}

	same1 := isSameTree(n1, n2)

	fmt.Println(same1)

	b1 := &TreeNode{1, nil, nil}
	b1.Left = &TreeNode{2, nil, nil}

	b2 := &TreeNode{1, nil, nil}
	b2.Right = &TreeNode{2, nil, nil}

	same2 := isSameTree(b1, b2)

	fmt.Println(same2)
}