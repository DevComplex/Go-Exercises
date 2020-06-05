package main

import "fmt"

type Node struct {
	val 	int
	left	*Node
	right	*Node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}

func getSecondLargestNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	largestNode := findLargestNode(n)
	return inorderPredecessor(largestNode, n)
}

func inorderPredecessor(n *Node, root *Node) *Node {
	if n.left != nil {
		return findLargestNode(n.left)
	}

	curr := root
	var predecessor *Node

	for curr != n {
		predecessor = curr
		curr = curr.right
	}

	return predecessor
}

func findLargestNode(n *Node) *Node {
	curr := n

	for curr.right != nil {
		curr = curr.right
	}

	return curr
}

func main() {
	root := NewNode(10)
	root.right = NewNode(20)
	root.right.left = NewNode(18)
	root.left = NewNode(5)

	fmt.Println(getSecondLargestNode(root).val)
}