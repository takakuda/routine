package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		switch i % 2 {
		case 0:
			fmt.Printf("%d-偶数\n", i)
		default:
			fmt.Printf("%d-奇数\n", i)
		}
	}
}
