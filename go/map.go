package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[12] = "JAPAN"

	fmt.Println(m)

	n := map[int]string{1: "Taka", 2: "Kuda"}
	fmt.Println(n)

	a := map[int]string{1: "A", 2: "B", 3: "C"}
	s := a[1]
	fmt.Println(s)
	s = a[9]
	fmt.Println(s)
}
