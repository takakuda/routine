package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	fmt.Println(s)

	a := make([]float64, 3)
	a[0] = 3.14
	fmt.Println(a)

	a[1] = 6.14
	fmt.Println(a)

	// panic: runtime error: index out of range [4] with length 3
	// fmt.Println(a[4])
}
