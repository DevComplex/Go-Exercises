package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	allLevels := [][]int{}

	queue := []*TreeNode{root}
	nextLvl := []*TreeNode{}

	currLvl := []int{}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		currLvl = append(currLvl, next.Val)

		if next.Left != nil {
			nextLvl = append(nextLvl, next.Left)
		}

		if next.Right != nil {
			nextLvl = append(nextLvl, next.Right)
		}

		if len(queue) == 0 {
			allLevels = append(allLevels, currLvl[:])
			queue = nextLvl[:]

			currLvl = []int{}
			nextLvl = []*TreeNode{}
		}
	}

	lo := 0 
	hi := len(allLevels) - 1

	for lo < hi {
		temp := allLevels[lo]
		allLevels[lo] = allLevels[hi]
		allLevels[hi] = temp

		lo++
		hi--
	}

	return allLevels
}

func main() {
	r := &TreeNode{3, nil, nil}
	r.Left = &TreeNode{9, nil, nil}
	r.Right = &TreeNode{20, nil, nil}
	r.Right.Left = &TreeNode{15, nil, nil}
	r.Right.Right = &TreeNode{7, nil, nil}

	levels := levelOrderBottom(r)

	for _, level := range levels {
		for _, value := range level {
			fmt.Println(value)
		}
	}
}