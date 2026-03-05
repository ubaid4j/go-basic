package main

import (
	"fmt"
)

type NodeS struct {
	Value int
	Next  *NodeS
}

var sizeS = 0
var stack = new(NodeS)

func PushS(v int) bool {
	if stack == nil {
		stack = &NodeS{v, nil}
		sizeS = 1
		return true
	}

	temp := &NodeS{v, nil}
	temp.Next = stack
	stack = temp
	sizeS++
	return true
}

func PopS(t *NodeS) (int, bool) {
	if sizeS == 0 {
		return 0, false
	}

	if sizeS == 1 {
		sizeS = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	sizeS--
	return t.Value, true
}

func traverseS(t *NodeS) {
	if sizeS == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	v, b := PopS(stack)
	if b {
		fmt.Println(v, " ")
	} else {
		fmt.Println("Pop() failed")
	}

	PushS(100)
	traverseS(stack)
	PushS(200)
	traverseS(stack)

	for i := 0; i < 10; i++ {
		PushS(i)
	}

	for i := 0; i < 15; i++ {
		v, b := PopS(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverseS(stack)
}
