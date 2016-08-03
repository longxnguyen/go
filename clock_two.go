package main

import (
  "io"
  "log"
  "net"
  "time"
)

func main() {
  listener, err := net.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Fatal(err)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err)
      continue
    }

    go handleConn(conn)
  }
}

func handleConn(c net.Conn) {
  defer c.Close()
  for {
    str := time.Now().Format("15:04:05\n")
    _, err := io.WriteString(c, str)
    if err != nil {
      return // client disconnects.
    }
    time.Sleep(1 * time.Second)
  }
}
