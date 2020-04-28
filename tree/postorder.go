package main

import "fmt"

type Node struct {
	Val int	
	Children []*Node
}

func initNode(val int) *Node {
	return &Node{val, []*Node{}}
}


func (n *Node) addChild(child *Node) {
	n.Children = append(n.Children, child)
}

func postorder(root *Node) []int {
	traversal := []int{}
	postorderHelper(root, &traversal)
	return traversal
}

func postorderHelper(root *Node, traversal *[]int) {
	if root == nil {
		return
	}

	for _, child := range root.Children {
		postorderHelper(child, traversal)
	}

	*traversal = append(*traversal, root.Val)
}

func main() {
	n1 := initNode(1)	
	n2 := initNode(3)
	n3 := initNode(2)
	n4 := initNode(4)
	n5 := initNode(5)
	n6 := initNode(6)

	n1.addChild(n2)
	n1.addChild(n3)
	n1.addChild(n4)

	n2.addChild(n5)
	n2.addChild(n6)

	traversal := postorder(n1)

	for _, node := range traversal {
		fmt.Println(node)
	}
}