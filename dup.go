// Prints the count and text of lines that appear more than once in the named
// input file.
package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func main() {
  counts := make(map[string]int)
  // Reads input and counts.
  for _, filename := range os.Args[1:] {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup: %v\n", err)
      continue
    }

    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
    }
  }

  // Prints out the result
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
