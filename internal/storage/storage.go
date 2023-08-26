package storage

import (
	"context"
	"database/sql"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/storage/repos"
)

type User interface {
	SaveUser(ctx context.Context, userId int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, userId int) error
	FollowToSegments(ctx context.Context, userId int, segments []string) error
	UnFollowToSegments(ctx context.Context, userId int, segments []string) error
	GetUserSegments(ctx context.Context, id int) ([]entity.Segment, error)
}

type Segment interface {
	SaveSegment(ctx context.Context, name string) (int, error)
	GetSegments(ctx context.Context) ([]entity.Segment, error)
	DeleteSegment(ctx context.Context, name string) error
}

type History interface {
	GetHistory(ctx context.Context) ([]entity.History, error)
}

type Storage struct {
	User
	Segment
	History
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		User:    repos.NewUserStorage(db),
		Segment: repos.NewSegmentStorage(db),
		History: repos.NewHistoryStorage(db),
	}
}
