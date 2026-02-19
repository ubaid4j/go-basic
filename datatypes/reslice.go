package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println("s1", s1)
	fmt.Println("reslice", reSlice)
	fmt.Println()
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println("s1", s1)
	fmt.Println("reslice", reSlice)
}
