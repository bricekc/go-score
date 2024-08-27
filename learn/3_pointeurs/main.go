package main

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

func printList[T any](n *Node[T]) {
	current := n
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func reverseList[T any](n *Node[T]) *Node[T] {
	var prev *Node[T]
	current := n
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	node1 := &Node[string]{value: "coucou1"}
	node2 := &Node[string]{value: "coucou2"}
	node3 := &Node[string]{value: "coucou3"}
	node1.next = node2
	node2.next = node3
	printList(node1)
	reverseList(node1)
	printList(node3)
}
