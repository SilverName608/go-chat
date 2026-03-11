package hub

import (
	"log"

	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 4096
)

type Client struct {
	Hub    *Hub
	RoomID string
	UserID string
	Conn   *websocket.Conn
	Send   chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.Hub.Broadcast <- &BroadcastMessage{
			RoomID:  c.RoomID,
			UserID:  c.UserID,
			Payload: message,
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for {
		message, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("error: %v", err)
			return
		}
	}
}
