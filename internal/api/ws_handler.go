package api

import (
	"net/http"

	"github.com/SilverName608/go-chat/internal/hub"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSHandler struct {
	hub *hub.Hub
}

func NewWSHandler(hub *hub.Hub) *WSHandler {
	return &WSHandler{hub: hub}
}

func (wh *WSHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	roomIDStr := chi.URLParam(r, "roomID")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	client := &hub.Client{
		Hub:    wh.hub,
		RoomID: roomID.String(),
		UserID: userID.String(),
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}

	wh.hub.Register <- client

	go client.ReadPump()
	go client.WritePump()
}
