package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[12] = "JAPAN"

	fmt.Println(m)
}
