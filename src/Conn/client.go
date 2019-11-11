package Conn

import (
	"log"
	"runtime"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

var c *gosocketio.Client

func sendJoin(c *gosocketio.Client) {
	log.Println("Acking /join")
	result, err := c.Ack("/join", Channel{"main"}, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Ack result to /join: ", result)
	}
}

func Client() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 8000, false),
		transport.GetDefaultWebsocketTransport())
	log.Println("Client Created")
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("/message", func(h *gosocketio.Channel, args Message) {
		log.Println("--- Got chat message: ", args)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	go sendJoin(c)
	go sendJoin(c)
	go sendJoin(c)
	go sendJoin(c)
	go sendJoin(c)

	c.Emit("poker", "Let's play poker")

	time.Sleep(60 * time.Second)
}

func CloseClient() {
	c.Close()
	log.Println("Client Closed")
}
