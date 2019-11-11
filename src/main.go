package main

import (
	"os"

	"./Conn"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "server" {
		Conn.Server()
	} else {
		Conn.Client()
	}
}
