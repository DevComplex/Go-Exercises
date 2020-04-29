package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	curr := root
	stack := []*TreeNode{}

	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}

	rangeSum := 0

	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		fmt.Println(top.Val)

		if top.Val >= L && top.Val <= R {
			rangeSum += top.Val
		} else if top.Val > R {
			break
		}

		if top.Right != nil {
			curr = top.Right

			for curr != nil {
				stack = append(stack, curr)
				curr = curr.Left
			}
		}
	}

	return rangeSum
}

func initNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func main() {
	root := initNode(10)
	root.Left = initNode(5)
	root.Right = initNode(15)
	root.Left.Left = initNode(3)
	root.Left.Right = initNode(7)
	root.Right.Right = initNode(18)

	R := 15
	L := 7

	result := rangeSumBST(root, L, R)

	fmt.Println(result)
}