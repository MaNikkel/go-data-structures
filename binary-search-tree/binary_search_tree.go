package binarysearchtree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewTree(value int) *Node {
	return &Node{Value: value, Left: nil, Right: nil}
}

func (n *Node) InsertLeaf(value int) *Node {
	if n == nil {
		return &Node{Value: value, Left: nil, Right: nil}
	}

	if n.Value > value {
		n.Left = n.Left.InsertLeaf(value)
	} else {
		n.Right = n.Right.InsertLeaf(value)

	}

	return n
}

func (r *Node) Print() {
	if r != nil {
		r.Left.Print()

		fmt.Println(r.Value)

		r.Right.Print()
	}
}

func (r *Node) MinValue() int {
	if r.Left != nil {
		return r.Left.MinValue()
	}

	return r.Value
}

func (r *Node) MaxValue() int {
	if r.Right != nil {
		return r.Right.MaxValue()
	}

	return r.Value
}
