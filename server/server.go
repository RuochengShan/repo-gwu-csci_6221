package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Let's have a party")
	for i := 0; ; i++ {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go start(conn, i)
	}
}
