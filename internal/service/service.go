package service

import (
	"context"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/service/services"
	"github.com/modaniru/avito/internal/storage"
)

type User interface {
	SaveUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Service struct {
	User
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		User: services.NewUserService(storage.User),
	}
}
