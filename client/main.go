package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {

	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Create a channel to handle signals for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Start a goroutine to read messages from the WebSocket connection
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				return
			}
			fmt.Printf("Received from server: %s\n", message)
		}
	}()
	<-interrupt
	fmt.Println("Shutting down...")

}
