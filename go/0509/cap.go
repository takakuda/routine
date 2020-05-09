package main

import (
	"fmt"
)

func main() {
	s := make([]int, 4)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	t := make([]int, 5, 10)
	fmt.Println(len(t))
	fmt.Println(cap(t))
}
