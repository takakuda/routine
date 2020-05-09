package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	a := []int{5, 6}

	n := copy(s, a)
	fmt.Println(n)
	fmt.Println(s)
	fmt.Println(a)
}
