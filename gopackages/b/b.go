package b

import (
	"a"
	"fmt"
)

func init() {
	fmt.Println("inti() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}
