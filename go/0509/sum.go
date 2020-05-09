package main

import (
	"fmt"
)

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
	s := []int{1, 2, 3}
	fmt.Println(sum(s...))
}
