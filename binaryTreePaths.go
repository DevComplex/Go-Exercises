package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	currPath := []string{}
	allPaths := []string{}
	binaryTreePathsHelper(root, &currPath, &allPaths)
	return allPaths
}

func removeLast(strs *[]string) {
	cp := *strs
	cp = cp[:len(cp)-1]
	*strs = cp
}

func binaryTreePathsHelper(n *TreeNode, currPath *[]string, allPaths *[]string) {
	if n == nil {
		return
	}

	*currPath = append(*currPath, strconv.Itoa(n.Val))

	if n.Left == nil && n.Right == nil {
		*allPaths = append(*allPaths, strings.Join(*currPath, "->"))
		removeLast(currPath)
		return
	}

	binaryTreePathsHelper(n.Left, currPath, allPaths)
	binaryTreePathsHelper(n.Right, currPath, allPaths)
	removeLast(currPath)
}

func main() {
	n1 := &TreeNode{1, nil, nil}
	n1.Left = &TreeNode{2, nil, nil}
	n1.Left.Right = &TreeNode{3, nil, nil}
	n1.Right = &TreeNode{4, nil, nil}

	paths := binaryTreePaths(n1)

	fmt.Println(paths)
}
