package main

import (
	"fmt"
)

func main() {
	s := make([]int, 8)
	fmt.Println(len(s))

	a := [3]int{1, 2, 3}
	fmt.Println(len(a))
}
