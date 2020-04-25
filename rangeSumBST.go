package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func buildTree(data []*int) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	nodes := []*TreeNode{}

	for _, val := range data {
		if val == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, initTreeNode(*val))
		}
	}

	return buildTreeHelper(nodes, 0)
}

func buildTreeHelper(data []*TreeNode, index int) *TreeNode {
	if index >= len(data) {
		return nil
	}

	n := data[index]

	if n != nil {
		n.Left = buildTreeHelper(data, 2*index + 1)
		n.Right = buildTreeHelper(data, 2*index + 2)
	}

	return n
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := 0

	if root.Val >= L && root.Val <= R {
		sum = root.Val
	}

	return sum + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
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
	n1 := 10
	n2 := 5
	n3 := 15
	n4 := 3
	n5 := 7
	n6 := 18


	data := []*int{&n1, &n2, &n3, &n4, &n5, nil, &n6}

	tree := buildTree(data)

	fmt.Printf("Printing inorder...")
	inorder(tree)

	L := 7
	R := 15

	sum := rangeSumBST(tree, L, R)

	fmt.Printf("Printing range sum %d <= x <= %d...", L, R)
	fmt.Printf("\n%d", sum)
}