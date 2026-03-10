package service

import (
	"context"

	domainModel "github.com/SilverName608/go-chat/internal/domain/model"
	"github.com/google/uuid"
)

type RoomService interface {
	// Create — создаёт новую комнату, ownerID берётся из JWT
	Create(ctx context.Context, name, description string, ownerID uuid.UUID) (*domainModel.Room, error)

	// GetAll — возвращает список всех комнат
	GetAll(ctx context.Context) ([]*domainModel.Room, error)

	// GetByID — возвращает комнату по ID
	GetByID(ctx context.Context, id uuid.UUID) (*domainModel.Room, error)

	// Delete — удаляет комнату, только владелец может удалить
	Delete(ctx context.Context, id uuid.UUID, requesterID uuid.UUID) error
}
