package bitwise

type node struct {
	isEnd    bool
	children [2]*node
}

func newNode() *node {
	return &node{
		children: [2]*node{},
	}
}

func (n *node) insertBit(b byte, i int) {
	bit := b & 0x1
	if n.children[bit] == nil {
		n.children[bit] = newNode()
	}
	if i+1 == 8 {
		n.isEnd = true
		return
	}
	n.children[bit].insertBit(b, i+1)
}

func (n *node) insert(word string, char byte, i int) {
	bit := char & 0x1
	char >>= 1

	if n.children[bit] == nil {
		n.children[bit] = newNode()
	}

	if i+1 == 8 {
		if len(word) > 1 {
			n.children[bit].insert(word[1:], word[1], 0)
		} else {
			n.children[bit].isEnd = true
		}
	} else {
		n.children[bit].insert(word, char, i+1)
	}
}

func (n *node) Insert(word string) {
	if len(word) == 0 {
		return
	}
	n.insert(word, word[0], 0)
}

func (n *node) isExist(word string, char byte, i int) bool {
	bit := char & 0x1
	char >>= 1

	if n.children[bit] == nil {
		return false
	}

	if i+1 == 8 {
		if len(word) > 1 {
			return n.children[bit].isExist(word[1:], word[1], 0)
		} else {
			return n.children[bit].isEnd
		}
	} else {
		return n.children[bit].isExist(word, char, i+1)
	}
}

func (n *node) IsExist(word string) bool {
	if len(word) == 0 {
		return false
	}
	return n.isExist(word, word[0], 0)
}
