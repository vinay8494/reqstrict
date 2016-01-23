package main

import (
  "flag"
  "strconv"
  "net"
  "log"
)

var port int
func init() {
  const (
    defaultPort = 8000
    portUsage = "The proxy server's port"
  )
  flag.IntVar(&port, "p", defaultPort,portUsage)
}

var bindAddress string
func init() {
  const (
    defaultBindAddress = "0.0.0.0"
    addressUsage = "The proxy server's bind address"
  )
  flag.StringVar(&bindAddress, "b", defaultBindAddress, addressUsage)
}

var maxConnections int
func init() {
  const (
    defaultMaxConn = 20
    maxConnUsage = "Maximum connections to proxy"
  )
  flag.IntVar(&maxConnections, "m", defaultMaxConn, maxConnUsage)
}

func main() {
  flag.Parse()

  host := bindAddress + ":" + strconv.Itoa(port)
  server, err := net.Listen("tcp", host)
  if err != nil {
    log.Fatal("Error occured while listening:", err)
  }
  log.Println("Proxying on", host)

  waiting := make (chan net.Conn)
  spaces := make (chan bool, maxConnections)

  for i := 0; i < maxConnections; i++ {
    spaces <- true
  }

  go matchConnections( waiting , spaces)

  for {
    connection,err := server.Accept()
    if err != nil {
      log.Println(err)
    } else {
      waiting <- connection
    }
  }

}
