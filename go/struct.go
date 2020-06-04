package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

type Cat struct {
	Animal
	ServanName string
}

func main() {
	c := Cat{}
	c.Name = "taro"
	fmt.Println(c.Name)

	a := &Animal{Name: "tako", Age: 1}
	fmt.Println(a)

	b := new(Animal)
	fmt.Println(b)
}
