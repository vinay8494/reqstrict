package main

import (
  "net"
  "log"
  "strings"
  "io"
)

func matchConnections ( waiting chan net.Conn, spaces chan bool) {
  for connection := range waiting {
    <-spaces
    go func (connection net.Conn) {
      handleConnection(connection)
      spaces <- true
    }(connection)
  }
}

func handleConnection(connection net.Conn) {
  defer connection.Close()

  // read request from connection
  request, hostname := readRequest(connection)
  log.Println("Request:",request)
  log.Println("Hostname:",hostname)

  // TODO : filter request

  // connect to the host if not filtered out
  remote, err := net.Dial("tcp", hostname + ":80")
  if err != nil {
    log.Println("error %s occurred while connecting to host : %s",err,hostname)
    return
  } else {
    log.Println("successfully connected to ",remote)
  }

  // TODO : Write back response from the remote to the connection

}

func getHostName(request string) (string) {
  hostname := strings.Split((strings.Split(request,"\r\n")[1]),": ")[1]
  return hostname
}

func readRequest(connection net.Conn) (string, string) {
  buf := make([]byte, 0)
  temp := make([]byte, 256)
  for {
    n, err := connection.Read(temp)
    if err != nil {
      if err != io.EOF {
        log.Println("Read error:", err)
      }
      break
    }
    buf = append(buf, temp[:n]...)
    if (strings.HasSuffix(string(buf[:]),"\r\n\r\n")) {
      break
    }
  }
  if len(buf) != 0 {
    request := string(buf[:])
    return request, getHostName(request)
  }

  return "", ""
}
