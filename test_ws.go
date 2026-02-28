package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "wss", Host: "gateway.rory.server.spacebar.chat", Path: "/"}

	c, _, _ := websocket.DefaultDialer.Dial(u.String()+"?v=9&encoding=json", nil)
	defer c.Close()
	c.ReadMessage()

	identify := `{"op":2,"d":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1NiIsImlhdCI6MTYxNjIzOTAyMn0.fakesignature123456789012345678901234567890","intents":0,"properties":{"os":"linux","browser":"my_library","device":"my_library"}}}`
	c.WriteMessage(websocket.TextMessage, []byte(identify))

	for {
		c.SetReadDeadline(time.Now().Add(5 * time.Second))
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read loop:", err)
		}
		log.Printf("recv: %s", message)
	}
}
