package main

import (
  "fmt"
)

func main() {
  s := []int{1, 2, 3}
  s = append(s, 4)
  fmt.Println(s)

  // evaluated but not used
  // append(s, 5)
}


