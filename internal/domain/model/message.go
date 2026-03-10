package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct { //Сообщение
	ID        uuid.UUID `json:"id"`
	RoomId    uuid.UUID `json:"room_id"`
	UserId    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
