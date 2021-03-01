package trie

import "fmt"

const alphabetSize = 26

// todo test benchmark if node stores value
type node struct {
	isEnd    bool
	children [alphabetSize]*node
}

func newNode() *node {
	return &node{
		children: [26]*node{},
	}
}

func hash(char uint8) uint8 {
	return char - uint8('a')
}

func (n *node) Insert(word string) {
	l := len(word)
	if l == 0 {
		return
	}

	index := hash(word[0])
	if index >= alphabetSize {
		panic(fmt.Sprintf("invalid character:<%s>", string(word[0])))
	}
	if n.children[index] == nil {
		n.children[index] = newNode()
	}
	if l == 1 {
		n.children[index].isEnd = true
	} else {
		word = word[1:]
		n.children[index].Insert(word)
	}
}

func (n *node) IsExist(word string) bool {
	l := len(word)
	if l == 0 {
		return false
	}

	i := hash(word[0])
	if n.children[i] == nil {
		return false
	}

	if l == 1 {
		return true
	}

	return n.children[i].IsExist(word[1:])
}
