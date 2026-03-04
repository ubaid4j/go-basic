package main

import (
	"fmt"
)

type Node1 struct {
	Value int
	Next  *Node1
}

var root = new(Node1)

func addNode(t *Node1, v int) int {
	if root == nil {
		t = &Node1{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists: ", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node1{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverseList(t *Node1) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookupNode(t *Node1, v int) bool {
	if root == nil {
		t = &Node1{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

func size(t *Node1) int {
	if t == nil {
		fmt.Println("-> Empty List!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverseList(root)
	addNode(root, 1)
	addNode(root, -1)
	traverseList(root)
	addNode(root, 10)
	addNode(root, 11)
	addNode(root, 12)
	traverseList(root)

	if lookupNode(root, -1) {
		fmt.Println("Node exists")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists")
	} else {
		fmt.Println("Node does not exists")
	}

	fmt.Println("Size: ", size(root))
}
