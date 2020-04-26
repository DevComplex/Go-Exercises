package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func trimBST(n *TreeNode, L int, R int) *TreeNode {
	if n == nil {
		return nil
	}

	if n.Val < L {
		return trimBST(n.Right, L, R)
	}

	if n.Val > R {
		return trimBST(n.Left, L, R)
	}

	n.Left = trimBST(n.Left, L, R)
	n.Right = trimBST(n.Right, L, R)

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
	t := TreeNode{3, nil, nil}
	t.Left = &TreeNode{0, nil, nil}
	t.Right = &TreeNode{4, nil, nil}
	t.Left.Right = &TreeNode{2, nil, nil}
	t.Left.Right.Left = &TreeNode{1, nil, nil}

	n := trimBST(&t, 1, 3)

	inorder(n)
}