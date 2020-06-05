package main

import "fmt"

type TreeNode struct {
	Val	   int
	Left   *TreeNode
	Right  *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

type Store struct {
	Data	[]int
}

func NewStore() *Store {
	return &Store{[]int{}}
}

func getLonelyNodes(root *TreeNode) []int {
	store := NewStore()
	getLonelyNodesHelper(root, nil, store)
	return store.Data
}

func getLonelyNodesHelper(n *TreeNode, parent *TreeNode, store *Store) {
	if n == nil {
		return
	}

	parentHasOnlyLeftChild := parent != nil && parent.Left != nil && parent.Right == nil
	parentHasOnlyRightChild := parent != nil && parent.Left == nil && parent.Right != nil

	if parentHasOnlyLeftChild || parentHasOnlyRightChild {
		store.Data = append(store.Data, n.Val)
	}

	getLonelyNodesHelper(n.Left, n, store)
	getLonelyNodesHelper(n.Right, n, store)
}

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Right = NewNode(4)
	root.Right = NewNode(3)

	output := getLonelyNodes(root)

	for _, n := range output {
		fmt.Println(n)
	}
}