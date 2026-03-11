package service

import (
	"context"

	domainModel "github.com/SilverName608/go-chat/internal/domain/model"
)

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*domainModel.User, error)
	Login(ctx context.Context, email, password string) (token string, err error)
}
