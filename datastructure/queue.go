package main

import (
	"fmt"
)

type NodeQ struct {
	Value int
	Next  *NodeQ
}

var sizeQ = 0
var queue = new(NodeQ)

func Push(t *NodeQ, v int) bool {
	if queue == nil {
		queue = &NodeQ{v, nil}
		sizeQ++
		return true
	}

	t = &NodeQ{v, nil}
	t.Next = queue
	queue = t
	sizeQ++
	return true
}

func Pop(t *NodeQ) (int, bool) {
	if sizeQ == 0 {
		return 0, false
	}

	if sizeQ == 1 {
		queue = nil
		sizeQ--
		return t.Value, true
	}

	temp := t

	for t.Next != nil {
		temp = t
		t = t.Next
	}

	v := temp.Next.Value
	temp.Next = nil
	sizeQ--
	return v, true
}

func traverseQ(t *NodeQ) {
	if sizeQ == 0 {
		fmt.Println("Empty Queue!")
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size: ", sizeQ)
	v, b := Pop(queue)
	if b {
		fmt.Println("Pop: ", v)
	}
	fmt.Println("Size: ", sizeQ)

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	traverseQ(queue)
	fmt.Println("Size: ", sizeQ)
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop: ", v)
	}
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size: ", sizeQ)
	traverseQ(queue)
}
