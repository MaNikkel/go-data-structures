package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	linkedlist "github.com/MaNikkel/go-linkedlist"
)

func TestLinkedList(t *testing.T) {
	head := linkedlist.NewList(1)
	head.Push(2)
	head.Push(3)
	head.Push(4)

	result := head.Print()

	require.Equal(t, "1 -> 2 -> 3 -> 4", result)
}

func TestLinkedList_Remove(t *testing.T) {
	head := linkedlist.NewList(1)
	head.Push(2)
	head.Push(3)
	head.Push(4)
	head.Push(5)

	head = head.Remove(1)
	result := head.Print()

	require.Equal(t, "2 -> 3 -> 4 -> 5", result)

	head = head.Remove(3)
	result = head.Print()

	require.Equal(t, "2 -> 4 -> 5", result)

	head = head.Remove(4)
	result = head.Print()

	require.Equal(t, "2 -> 5", result)

	head = head.Remove(2)
	result = head.Print()

	require.Equal(t, "5", result)

	head = head.Remove(5)
	result = head.Print()

	require.Equal(t, "", result)
}
