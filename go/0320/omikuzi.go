package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 6; i++ {
		switch rand.Intn(10) {
		case 6:
			fmt.Printf("大吉%d\n", i)
		case 5, 4:
			fmt.Printf("中吉%d\n", i)
		case 3, 2:
			fmt.Printf("吉%d\n", i)
		case 1:
			fmt.Printf("凶%d\n", i)
		}
	}
}
