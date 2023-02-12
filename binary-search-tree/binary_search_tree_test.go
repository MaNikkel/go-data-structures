package binarysearchtree_test

import (
	"testing"

	binarysearchtree "github.com/MaNikkel/go-binarysearchtree"
	"github.com/stretchr/testify/require"
)

func TestBST(t *testing.T) {
	r := binarysearchtree.NewTree(18)
	r.InsertLeaf(20)
	r.InsertLeaf(15)
	r.InsertLeaf(25)

	r.Print()

	require.Equal(t, 15, r.Left.Value)
	require.Equal(t, 20, r.Right.Value)
	require.Equal(t, 25, r.Right.Right.Value)

}

func TestBST_MinValue(t *testing.T) {
	r := binarysearchtree.NewTree(18)
	r.InsertLeaf(20)
	r.InsertLeaf(15)
	r.InsertLeaf(25)
	r.InsertLeaf(5)

	min := r.MinValue()

	require.Equal(t, 5, min)

}

func TestBST_MaxValue(t *testing.T) {
	r := binarysearchtree.NewTree(18)
	r.InsertLeaf(20)
	r.InsertLeaf(25)
	r.InsertLeaf(35)
	r.InsertLeaf(45)
	r.InsertLeaf(55)
	r.InsertLeaf(-5)
	r.InsertLeaf(-25)
	r.InsertLeaf(225)
	r.InsertLeaf(-25)
	r.InsertLeaf(225)
	r.InsertLeaf(35)

	max := r.MaxValue()

	require.Equal(t, 225, max)

}
