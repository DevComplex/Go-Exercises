package main

import "fmt"

type TreeNode struct {
	Val	int	
	Left *TreeNode
	Right *TreeNode
}

func getBottomLeaves(n *TreeNode, parent *TreeNode, leaves *[]*TreeNode, parents map[*TreeNode]*TreeNode) {
	if n == nil {
		return
	}

	if parent != nil {
		parents[n] = parent
	}

	if n.Left == nil && n.Right == nil {
		*leaves = append(*leaves, n)
	}

	getBottomLeaves(n.Left, n, leaves, parents)
	getBottomLeaves(n.Right, n, leaves, parents)
}

func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	output := [][]int{}
	
	leaves := []*TreeNode{}
	parents := make(map[*TreeNode]*TreeNode)

	getBottomLeaves(root, nil, &leaves, parents)

	currLvl := []int{}
	nextLvl := []*TreeNode{}

	for len(leaves) > 0 {
		next := leaves[0]
		leaves = leaves[1:]

		if parents[next] != nil {
			parent := parents[next]

			if parent.Left == next {
				parent.Left = nil
			} else if parent.Right == next {
				parent.Right = nil
			}

			if parent.Left == nil && parent.Right == nil {
				nextLvl = append(nextLvl, parent)
			}
		}

		currLvl = append(currLvl, next.Val)

		if len(leaves) == 0 {
			output = append(output, currLvl[:])
			currLvl = []int{}
			leaves = nextLvl[:]
			nextLvl = []*TreeNode{}
		}
	}

	return output
}

func main() {
	n := &TreeNode{1, nil, nil}
	n.Left = &TreeNode{2, nil, nil}
	n.Left.Left = &TreeNode{4, nil, nil}
	n.Left.Right = &TreeNode{5, nil, nil}
	n.Right = &TreeNode{3, nil, nil}

	output := findLeaves(n)

	fmt.Println(output)
}