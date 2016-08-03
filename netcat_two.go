package main

import (
  "io"
  "log"
  "net"
  "os"
)

func main() {
  connection, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }
  defer connection.Close()

  go mustCopy(os.Stdout, connection)
  mustCopy(connection, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
  _, err := io.Copy(dst, src)
  if err != nil {
    log.Fatal(err)
  }
}
