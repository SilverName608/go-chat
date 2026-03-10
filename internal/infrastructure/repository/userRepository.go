package repository

import (
	"context"

	domainModel "github.com/SilverName608/go-chat/internal/domain/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	// Create — сохраняет нового пользователя в БД
	Create(ctx context.Context, user *domainModel.User) (*domainModel.User, error)

	// FindByID — ищет пользователя по ID
	FindByID(ctx context.Context, id uuid.UUID) (*domainModel.User, error)

	// FindByEmail — ищет пользователя по email (нужен для Login)
	FindByEmail(ctx context.Context, email string) (*domainModel.User, error)

	// Update — обновляет данные пользователя (пароль, username)
	Update(ctx context.Context, user *domainModel.User) (*domainModel.User, error)

	// Delete — удаляет пользователя по ID
	Delete(ctx context.Context, id uuid.UUID) error
}
