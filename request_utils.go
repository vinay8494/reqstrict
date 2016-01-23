package main

import (
	"io"
	"log"
	"net"
	"strings"
)

func getHostName(request string) string {
	hostname := strings.Split((strings.Split(request, "\r\n")[1]), ": ")[1]
	return hostname
}

func modifyRequest(request string) string {
	modifiedRequest := strings.Replace(request, "HTTP/1.1", "HTTP/1.0", -1)
	return strings.Replace(modifiedRequest, "Connection: keep-alive", "", -1)
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
		if strings.HasSuffix(string(buf[:]), "\r\n\r\n") {
			break
		}
	}
	if len(buf) != 0 {
		request := string(buf[:])
		return modifyRequest(request), getHostName(request)
	}
	return "", ""
}
