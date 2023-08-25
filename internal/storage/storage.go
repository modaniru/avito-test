package storage

import (
	"context"

	"github.com/modaniru/avito/internal/entity"
)

type User interface {
	SaveUser(ctx context.Context, userId int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, userId int) error
}

type Segment interface {
	SaveSegment(ctx context.Context, name string) (int, error)
	GetSegments(ctx context.Context) ([]entity.Segment, error)
	DeleteSegment(ctx context.Context, name string) error
}
