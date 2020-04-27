package main

import "fmt"

type TreeNode struct {
	Val	int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(temp)
	return root
}

func bfs(n *TreeNode) {
	if n == nil {
		return
	}

	queue := []*TreeNode{n}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.Left != nil {
			queue = append(queue, next.Left)
		}

		if next.Right != nil {
			queue = append(queue, next.Right)
		}

		fmt.Println(next.Val)
	}
}

func main() {
	n := &TreeNode{4, nil, nil}
	n.Left = &TreeNode{2, nil, nil}
	n.Left.Left = &TreeNode{1, nil, nil}
	n.Left.Right = &TreeNode{3, nil, nil}
	n.Right = &TreeNode{7, nil, nil}
	n.Right.Left = &TreeNode{6, nil, nil}
	n.Right.Right = &TreeNode{9, nil, nil}

	i := invertTree(n)

	bfs(i)
}