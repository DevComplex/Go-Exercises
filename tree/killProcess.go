package main

import "fmt"

func killProcess(pid []int, ppid []int, kill int) []int {
	tree := make(map[int][]int)
	N := len(pid)

	for i := 0; i < N; i++ {
		if tree[pid[i]] == nil {
			tree[pid[i]] = []int{}
		}

		if tree[ppid[i]] == nil {
			tree[ppid[i]] = []int{}
		}

		tree[ppid[i]] = append(tree[ppid[i]], pid[i])
	}

	killedProcesses := []int{}
	killProcessHelper(tree, &killedProcesses, kill)
	return killedProcesses
}

func killProcessHelper(tree map[int][]int, killedProcesses *[]int, curr int) {
	*killedProcesses = append(*killedProcesses, curr)
	children := tree[curr]

	for _, child := range children {
		killProcessHelper(tree, killedProcesses, child)
	}
}

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}

	killed := killProcess(pid, ppid, 5)

	fmt.Println(killed)
}