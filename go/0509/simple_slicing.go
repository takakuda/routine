package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:2]
	fmt.Println(s)
	t := a[2:]
	fmt.Println(t)

	b := [...]int{1, 2, 3}
	fmt.Println(len(b))
	fmt.Println(cap(b))
	// invalid slice index 4 (out of bounds for 3-element array)
	// fmt.Println(b[0:4])
}
