package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var conn net.Conn

func main() {
	conn, _ = net.Dial("tcp", "127.0.0.1:8081")
	go show()
	handle()
}

func handle() {
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
	}
}

func show() {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Close()
			os.Exit(0)
		}
		fmt.Println(message)
	}
}
