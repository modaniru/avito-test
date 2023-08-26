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
	UnFollowToSegments(ctx context.Context, userId int, segments []string) error
	GetUserSegments(ctx context.Context, userId int) ([]entity.Segment, error)
	FollowToSegments(ctx context.Context, userId int, segments []string) error
}

type Segment interface {
	SaveSegment(ctx context.Context, name string) (int, error)
	GetSegments(ctx context.Context) ([]entity.Segment, error)
	DeleteSegment(ctx context.Context, name string) error
}

type History interface {
	GetHistory(ctx context.Context) ([]entity.History, error)
}

type Service struct {
	User
	Segment
	History
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		User:    services.NewUserService(storage.User),
		Segment: services.NewSegmentService(storage.Segment),
		History: services.NewHistoryService(storage.History),
	}
}
