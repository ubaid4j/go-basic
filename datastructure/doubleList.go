package main

import (
	"fmt"
)

type Node2 struct {
	Value    int
	Previous *Node2
	Next     *Node2
}

var root2 = new(Node2)

func addNodeD(t *Node2, v int) int {
	if root2 == nil {
		t = &Node2{v, nil, nil}
		root2 = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists: ", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node2{v, temp, nil}
		return -2
	}

	return addNodeD(t.Next, v)
}

func traverseD(t *Node2) {
	if t == nil {
		fmt.Println("-> Empty list!")
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node2) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}

	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func sizeD(t *Node2) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNodeD(t *Node2, v int) bool {
	if root2 == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNodeD(t.Next, v)
}

func main() {
	fmt.Println(root2)
	root2 = nil
	traverseD(root2)
	addNodeD(root2, 1)
	addNodeD(root2, 1)
	traverseD(root2)
	addNodeD(root2, 10)
	addNodeD(root2, 5)
	addNodeD(root2, 0)
	addNodeD(root2, 0)
	traverseD(root2)
	fmt.Println("Size: ", sizeD(root2))
	reverse(root2)
}
