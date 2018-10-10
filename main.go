package main

import (
	"fmt"
)

type Answer struct {
	ttl  int
	name string
}

type Node struct {
	char   byte
	answer *Answer
	next   map[byte]*Node
}

func newNode(ttl int, char byte) *Node {
	var node Node
	node.char = char
	node.answer = nil
	node.next = make(map[byte]*Node)
	return &node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node with character \"%c\"", n.char)
}

func add(root *Node, piece byte, domain string) {
	if piece == root.char {
		root.answer = &Answer{0, domain}
		return
	}

	if _, ok := root.next[piece]; !ok {
		root.next[piece] = newNode(0, piece)
	}
	add(root.next[piece], domain[0], domain[1:])
}

func main() {
	node := newNode(0, '.')
	fmt.Println(node)
	add
}
