package main

import "fmt"

// Returns a new slice after appending the new integer
func appendInteger(a []int, k int) []int {
  var b []int
  bLen := len(a) + 1
  if bLen <= cap(a) {
    b = a[:bLen]
  } else {
    bCap := bLen
    if bCap < 2*len(a) {
      bCap = 2 * len(a)
    }
    b = make([]int, bLen, bCap)
    copy(b, a)
  }
  b[len(a)] = k
  return b
}

func main() {
  var x, y []int
  for i := 0; i < 10; i++ {
    y = appendInteger(x, i)
    fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
    x = y
  }
}
