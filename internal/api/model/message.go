package model

import (
	"time"

	"github.com/google/uuid"
)

// --- Requests ---

// SendMessageRequest — приходит от клиента через WebSocket (JSON)
type SendMessageRequest struct {
	Body string `json:"body" validate:"required,min=1,max=4096"`
}

// --- Responses ---

type MessageResponse struct {
	ID        uuid.UUID `json:"id"`
	RoomID    uuid.UUID `json:"room_id"`
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"` // подтягиваем чтобы не делать запрос на фронте
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageListResponse struct {
	Messages []*MessageResponse `json:"messages"`
	Total    int                `json:"total"`
	Limit    int                `json:"limit"`
	Offset   int                `json:"offset"`
}

// WSMessage — конверт для WebSocket, type говорит фронту что делать с payload
type WSMessage struct {
	Type    string          `json:"type"` // "message", "join", "leave", "error"
	Payload MessageResponse `json:"payload"`
}
