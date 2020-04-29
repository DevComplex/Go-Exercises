package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
    var closest *int
	
	curr := root
	stack := []*TreeNode{}

	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}

	for len(stack) > 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if closest != nil {
			currDiff := math.Abs(target - float64(*closest))
			newDiff := math.Abs(target - float64(top.Val))
	
			if newDiff < currDiff {
				closest = &top.Val
			}
		} else {
			closest = &top.Val
		}

		if float64(top.Val) > target {
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

	return *closest
}

func initNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func main() {
	root := initNode(4)
	root.Left = initNode(2)
	root.Left.Left = initNode(1)
	root.Left.Right = initNode(3)
	root.Right = initNode(5)

	target := 3.714286

	fmt.Println(closestValue(root, target))
}