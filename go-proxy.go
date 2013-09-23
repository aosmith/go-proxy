package main

import (
  "code.google.com/p/go.net/websocket"
  "config"
  "fmt"
//  "io"
  "net"
  "net/http"
)

type connect struct {
  ws *websocket.Conn
  send chan string
  receive chan string
}

func ProxyServer(ws *websocket.Conn) {
  fmt.Println("Got websocket data: ")
  fmt.Println(ws)
  conn, err := net.Dial("tcp", config.DestAddress)
  if err != nil {
    panic("Error: " + err.Error())
  }
  defer conn.Close()
  for ;; {
    var msg = make([]byte, config.BufferSize)
    n := 0
    for n==0 {
      n, err = ws.Read(msg)
    }
    fmt.Printf("%s", msg[:n])
    conn, err := net.Dial("tcp", config.DestAddress)
    defer conn.Close()
    if err != nil {
      panic("Error opening tcp connection: " + err.Error())
    }
    conn.Write(msg[:n])
    n = 0
    var msg2 = make([]byte, config.BufferSize)
    for n==0 {
      n, err = conn.Read(msg)
    }
    fmt.Printf("%s", msg[:n])
    ws.Write(msg2[:n])
  }
  // io.Copy(ws,conn)

}

func main() {
  fmt.Println("Booting websocket server...")
  http.Handle("/", websocket.Handler(ProxyServer))
  err := http.ListenAndServe(":3337", nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}
