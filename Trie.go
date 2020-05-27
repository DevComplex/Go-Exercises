package main

import "fmt"

type TrieNode struct {
	IsWordComplete	bool
	children		map[rune]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		IsWordComplete: false,
		children: make(map[rune]*TrieNode),
	}
}

func Constructor() Trie {
	return Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	currNode := t.Root

	for _, char := range word {
		_, hasNode := currNode.children[char]

		if !hasNode {
			currNode.children[char] = NewTrieNode()
		}

		currNode = currNode.children[char]
	}

	currNode.IsWordComplete = true
}

func (t *Trie) Search(word string) bool {
	currNode := t.Root

	for _, char := range word {
		currNode = currNode.children[char]

		if currNode == nil {
			return false
		}
	}

	return currNode.IsWordComplete
}

func (t *Trie) StartsWith(prefix string) bool {
	currNode := t.Root

	for _, char := range prefix {
		currNode = currNode.children[char]

		if currNode == nil {
			return false
		}
	}

	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("grapefruit")

	fmt.Println(trie.StartsWith("applez"))
}