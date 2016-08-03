package main

import (
  "fmt"
)

func reverse(a []int) {
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
  }
}

func main() {
  // Array
  a := [...]int{0, 1, 2, 3, 4, 5}
  reverse(a[:]) // Creates a new slice that refers to all elements
  fmt.Println(a)

  // slice
  b := []int{0, 1, 2, 3, 4, 5}
  reverse(b[:3])
  reverse(b[3:])
  reverse(b)
  fmt.Println(b)
}
