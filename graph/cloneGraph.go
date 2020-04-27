package main 

import "fmt"

type Node struct {
	Val int
	Neighbors []*Node
}

func initNode(val int) *Node {
	return &Node{val, []*Node{}}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	graph := make(map[int]*Node)
	graph[node.Val] = initNode(node.Val)

	visited := make(map[int]bool)
	visited[node.Val] = true

	queue := []*Node{node}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		parentClone := graph[next.Val]

		for _, neighbor := range next.Neighbors {
			if graph[neighbor.Val] == nil {
				graph[neighbor.Val] = initNode(neighbor.Val)
			}

			neighborClone := graph[neighbor.Val]
			parentClone.Neighbors = append(parentClone.Neighbors, neighborClone)
			
			if !visited[neighbor.Val] {
				visited[neighbor.Val] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return graph[node.Val]
}

func bfs(n *Node) {
	queue := []*Node{n}
	
	visited := make(map[int]bool)
	visited[n.Val] = true

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		fmt.Println(next.Val)

		for _, node := range next.Neighbors {
			wasVisited := visited[node.Val]

			if !wasVisited {
				visited[node.Val] = true
				queue = append(queue, node)
			}
		}
	}

}

func main() {
	n1 := initNode(1)
	n2 := initNode(2)
	n3 := initNode(3)
	n4 := initNode(4)

	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n3)
	n3.Neighbors = append(n3.Neighbors, n4)

	n2.Neighbors = append(n2.Neighbors, n1)
	n2.Neighbors = append(n2.Neighbors, n3)

	n3.Neighbors = append(n3.Neighbors, n2)
	n3.Neighbors = append(n3.Neighbors, n4)

	n4.Neighbors = append(n4.Neighbors, n3)

	n4Copy := cloneGraph(n4)

	bfs(n4Copy)
}