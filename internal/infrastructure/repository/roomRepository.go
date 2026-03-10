package repository

import (
	"context"

	domainModel "github.com/SilverName608/go-chat/internal/domain/model"

	"github.com/google/uuid"
)

type RoomRepository interface {
	// Create — сохраняет новую комнату в БД
	Create(ctx context.Context, room *domainModel.Room) (*domainModel.Room, error)

	// FindAll — возвращает все комнаты
	FindAll(ctx context.Context) ([]*domainModel.Room, error)

	// FindByID — ищет комнату по ID
	FindByID(ctx context.Context, id uuid.UUID) (*domainModel.Room, error)

	// Update — обновляет данные комнаты (название, описание)
	Update(ctx context.Context, room *domainModel.Room) (*domainModel.Room, error)

	// Delete — удаляет комнату по ID
	Delete(ctx context.Context, id uuid.UUID) error
}
