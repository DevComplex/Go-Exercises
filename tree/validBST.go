package main

import (
	"fmt"
)

type node struct {
	val int
	left *node
	right *node
}

func Node(val int) *node {
	return &node{val, nil, nil}
}

func Root(val int) node {
	return node{val, nil, nil}
}

type TestFn func (*node) bool

func isValidBST(n *node) bool {
	nodes := []*node{}

	inorder(n, &nodes)

	for i := 0; i < len(nodes) - 1; i++ {
		if nodes[i].val > nodes[i + 1].val {
			return false
		}
	}

	return true
}

func inorder(n *node, nodes *[]*node) {
	if n == nil {
		return
	}

	inorder(n.left, nodes)
	*nodes = append(*nodes, n)
	inorder(n.right, nodes)
}

type testcase struct {
	input        *node
	description  string
	assert		 bool
}

type testsuite struct {
	testcases []*testcase
}

func (t *testsuite) run(fn TestFn) {
	for _, tc := range t.testcases {
		fmt.Println(tc.description)
		result := fn(tc.input)

		if result == tc.assert {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}

		fmt.Println()
	}
}

func main() {
	r1 := Root(10)
	r1.left = Node(5)
	r1.left.left = Node(2)
	r1.left.right = Node(6)
	r1.left.left.left = Node(100)

	r2 := Root(20)
	r2.left = Node(5)
	r2.right = Node(100)

	cases := []*testcase{}

	t1 := &testcase{&r1, "Invalid binary tree should assert false...", false}
	t2 := &testcase{&r2, "Valid binary tree should assert true...", true}

	cases = append(cases, t1)
	cases = append(cases, t2)

	ts1 := testsuite{cases}

	ts1.run(isValidBST)
}