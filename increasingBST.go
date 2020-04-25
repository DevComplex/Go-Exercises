package main

import "fmt"

type TreeNode struct {
   Value   int
   Left    *TreeNode
   Right   *TreeNode
}

func increasingBST(n *TreeNode) *TreeNode {
   nodes := []*TreeNode{}
   inorder(n, &nodes)

   for i := 0; i < len(nodes) - 1; i++ {
      nodes[i].Right = nodes[i + 1]
      nodes[i].Left = nil      
   }  

   return nodes[0] 
}

func inorder(n *TreeNode, nodes *[]*TreeNode) {
   if n == nil {
      return
   }

   inorder(n.Left, nodes)
   *nodes = append(*nodes, n)
   inorder(n.Right, nodes)
}

func print(n *TreeNode) {
   if n == nil {
      return 
   }

   fmt.Println(n.Value)
   print(n.Right)
}

func main() {
   n := &TreeNode{20, nil, nil}
   b := &TreeNode{10, nil, nil}
   c := &TreeNode{30, nil, nil}

   n.Left = b
   n.Right = c

   t := increasingBST(n)

   print(t)
}
