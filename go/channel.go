package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	ch <- 5
	i := <-ch

	fmt.Println(i)
}
