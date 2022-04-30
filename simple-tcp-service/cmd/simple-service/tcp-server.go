package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Launching a new server")

	listener, _ := net.Listen("tcp", "127.0.0.1:8081")

	newconnection, _ := listener.Accept()

	for {
		message, _ := bufio.NewReader(newconnection).ReadString('\n')
		fmt.Print("Received message from client:", string(message))
		newconnection.Write([]byte("Hello, " + message + "\n"))
	}
}
