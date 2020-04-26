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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	sum := 0

	var left1, left2, right1, right2 *TreeNode

	if t1 != nil {
		sum = sum + t1.Val
		left1 = t1.Left
		right1 = t1.Right
	}

	if t2 != nil {
		sum = sum + t2.Val
		left2 = t2.Left
		right2 = t2.Right
	}

	n := initTreeNode(sum)

	n.Left = mergeTrees(left1, left2)
	n.Right = mergeTrees(right1, right2)

	return n
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
	t1 := initTreeNode(1)
	t1.Left = initTreeNode(3)
	t1.Left.Left = initTreeNode(5)
	t1.Right = initTreeNode(2)

	t2 := initTreeNode(2)
	t2.Left = initTreeNode(1)
	t2.Left.Right = initTreeNode(4)
	t2.Right = initTreeNode(3)
	t2.Right.Right = initTreeNode(7)

	mergedTree := mergeTrees(t1, t2)

	bfs(mergedTree)
}