package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6}
	for _, i := range num {
		switch i {
		case 6:
			fmt.Println("大吉")
		case 5, 4:
			fmt.Println("中吉")
		case 3, 2:
			fmt.Println("吉")
		case 1:
			fmt.Println("凶")
		}
	}
}
