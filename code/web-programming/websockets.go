package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("wss://echo.websocket.org", nil)
	if err != nil {
		log.Fatal("err dial:", err)
	}
	defer c.Close()

	go func() { c.WriteMessage(websocket.TextMessage, []byte("wat")) }()

	messageType, message, err := c.ReadMessage()
	log.Printf("type: %d, message: %s; err: %v", messageType, message, err)
}
