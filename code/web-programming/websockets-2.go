package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	go http.ListenAndServe(":8282", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := (&websocket.Upgrader{}).Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade error:", err)
			return
		}
		defer c.Close()

		mt, message, _ := c.ReadMessage()
		c.WriteMessage(mt, []byte(strings.ToUpper(string(message))+"!!!"))
	}))

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8282", nil)
	if err != nil {
		log.Fatal("err dial:", err)
	}
	defer c.Close()

	go func() { c.WriteMessage(websocket.TextMessage, []byte("wat")) }()

	messageType, message, err := c.ReadMessage()
	log.Printf("type: %d, message: %s; err: %v", messageType, message, err)
}
