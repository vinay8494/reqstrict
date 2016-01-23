package main

import (
	"log"
	"net"
)

func copyData(from net.Conn, to net.Conn) {
	var bytes []byte = make([]byte, 256)
	for {
		read, err := from.Read(bytes)
		if err != nil {
			break
		}
		_, err = to.Write(bytes[:read])
		if err != nil {
			break
		}
	}
}

func matchConnections(waiting chan net.Conn, spaces chan bool) {
	for connection := range waiting {
		<-spaces
		go func(connection net.Conn) {
			handleConnection(connection)
			spaces <- true
		}(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	request, hostname := readRequest(connection)

	// TODO : filter request

	remote, err := net.Dial("tcp", hostname+":80")
	if err != nil {
		log.Println("error", err, "occurred while connecting to host :", hostname)
		return
	}
	defer remote.Close()

	_, err = remote.Write([]byte(request))
	if err != nil {
		log.Println(err)
		return
	}

	copyData(remote, connection)
}
