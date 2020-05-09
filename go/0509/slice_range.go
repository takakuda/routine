package main

import (
	"fmt"
)

func main() {
	s := []string{"apple", "banana", "cherry"}
	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}
}
