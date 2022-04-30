package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	newconnection, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(newconnection, text+"\n")
		message, _ := bufio.NewReader(newconnection).ReadString('\n')
		fmt.Print("Received message from server: " + message)
	}
}
