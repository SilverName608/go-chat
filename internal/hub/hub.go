package hub

import (
	"log"
)

type BroadcastMessage struct {
	RoomID  string
	UserID  string
	Payload []byte
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *BroadcastMessage
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *BroadcastMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			room, ok := h.Rooms[client.RoomID]
			if !ok {
				room = NewRoom(client.RoomID)
				h.Rooms[client.RoomID] = room
			}
			room.clients[client] = true
			log.Printf("client joined room %s", client.RoomID)

		case client := <-h.Unregister:
			room, ok := h.Rooms[client.RoomID]
			if !ok {
				continue
			}
			delete(room.clients, client)
			close(client.Send)
			log.Printf("client left room %s", client.RoomID)
			if len(room.clients) == 0 {
				delete(h.Rooms, client.RoomID)
			}

		case msg := <-h.Broadcast:
			room, ok := h.Rooms[msg.RoomID]
			if !ok {
				continue
			}
			for client := range room.clients {
				select {
				case client.Send <- msg.Payload:
				default:
					close(client.Send)
					delete(room.clients, client)
				}
			}
		}
	}
}
