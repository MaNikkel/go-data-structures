package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewList(value int) *Node {
	return &Node{value: value, next: nil}
}

func (n *Node) Push(value int) {

	if n.next != nil {
		n.next.Push(value)
		return
	}

	node := &Node{value: value, next: nil}
	n.next = node

}

func (h *Node) Print() string {
	formatted := ""
	for n := h; n != nil; n = n.next {
		formatted += fmt.Sprint(n.value)
		if n.next != nil {
			formatted += " -> "
		}
	}
	return formatted
}

func (h *Node) Remove(value int) *Node {
	if h.value == value {
		return h.next
	}

	for n := h; n != nil; n = n.next {

		if n.next != nil && n.next.value == value {
			newNext := n.next.next

			n.next = newNext

		}
	}

	return h
}
