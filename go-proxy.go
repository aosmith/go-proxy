package main

import (
  "code.google.com/p/go.net/websocket"
  "config"
  "fmt"
  "net"
  "net/http"
  "time"
)

func proxy(conn net.Conn, ws *websocket.Conn) {
  for ;; {
    var msg = make([]byte, config.BufferSize)
    var err error
    // Check websocket for data + send to server
    n := 0
    n, err = ws.Read(msg)
    if err != nil {
      panic("Error: " + err.Error())
    } else if n > 0 {
      fmt.Fprintf(conn, "%s", msg[:n])
    }
    // Check tcpsocket for data + send to client
    n = 0
    n, err = conn.Read(msg)
    if err != nil {
      panic("Error: " + err.Error())
    } else if n > 0 {
      fmt.Fprintf(ws, "%s", msg[:n])
    }
    time.Sleep(config.SleepTime * time.Millisecond)
  }
}

func ProxyServer(ws *websocket.Conn) {
  fmt.Println("New websocket connection")
  conn, err := net.Dial("tcp", config.DestAddress)
  if err != nil {
    panic("Error: " + err.Error())
  }
  defer conn.Close()
  proxy(conn, ws)
}

func main() {
  fmt.Println("Booting websocket server...")
  http.Handle("/", websocket.Handler(ProxyServer))
  err := http.ListenAndServe(config.WsPort, nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}
