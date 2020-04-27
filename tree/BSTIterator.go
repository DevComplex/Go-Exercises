package main

import "fmt"

type TreeNode struct {
	Val	int
	Left *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Stack	[]*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	curr := root

	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}

	return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
	top := this.Stack[len(this.Stack) - 1]
	this.Stack = this.Stack[:len(this.Stack) - 1]

	if top.Right != nil {
		curr := top.Right

		for curr != nil {
			this.Stack = append(this.Stack, curr)
			curr = curr.Left
		}
	}

	return top.Val
}

func (this *BSTIterator) HasNext() bool {
    return len(this.Stack) > 0
}

func main() {
	n := &TreeNode{7, nil, nil}
	n.Left = &TreeNode{3, nil, nil}
	n.Right = &TreeNode{15, nil, nil}
	n.Right.Right = &TreeNode{20, nil, nil}
	n.Right.Left = &TreeNode{9, nil, nil}

	iter := Constructor(n)

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}