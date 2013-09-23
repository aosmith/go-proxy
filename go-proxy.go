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

func proxy(conn *net.Conn, ws *websocket.Conn) {
  for ;; {
    var msg = make([]byte, config.BufferSize)
    n := 0
    for n==0 {
      n, err = ws.Read(msg)
    }
    fmt.Printf("%s", msg[:n])
    conn.Write(msg[:n])
  }
}

func proxy(ws *websocket.Conn, conn *net.Conn) {
  for ;; {
    var msg = make([]byte, config.BufferSize)
    n := 0
    for n==0 {
      n, err = conn.Read(msg)
    }
    fmt.Printf("%s", msg[:n])
    ws.Write(msg[:n])
  }
}

func ProxyServer(ws *websocket.Conn) {
  fmt.Println("New websocket connection")
  conn, err := net.Dial("tcp", config.DestAddress)
  if err != nil {
    panic("Error: " + err.Error())
  }
  defer conn.Close()
  for ;; {
    go proxy(ws, conn)
    go proxy(conn, ws)
  }
}

func main() {
  fmt.Println("Booting websocket server...")
  http.Handle("/", websocket.Handler(ProxyServer))
  err := http.ListenAndServe(":3337", nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}
