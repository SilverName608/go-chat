package hub

type Room struct {
	id      string
	clients map[*Client]bool
}

func NewRoom(id string) *Room {
	return &Room{
		id:      id,
		clients: make(map[*Client]bool),
	}
}
