package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/storage/repos"
)

//go:generate mockery --name User
type User interface {
	SaveUser(ctx context.Context, userId int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, userId int) error
	FollowToSegments(ctx context.Context, userId int, segments []string, date *string) error
	UnFollowToSegments(ctx context.Context, userId int, segments []string) error
	GetUserSegments(ctx context.Context, id int) ([]entity.Follows, error)
	FollowRandomUsers(ctx context.Context, name string, percent float64) (int, error)
}

//go:generate mockery --name Segment
type Segment interface {
	SaveSegment(ctx context.Context, name string) (int, error)
	GetSegments(ctx context.Context) ([]entity.Segment, error)
	DeleteSegment(ctx context.Context, name string) error
}

//go:generate mockery --name History
type History interface {
	GetHistory(ctx context.Context) ([]entity.History, error)
	GetHistoryByDate(ctx context.Context, date string) ([]entity.History, error)
}

type Follow interface {
	DeleteExpiredFollows() (int, error)
}

type Storage struct {
	User
	Segment
	History
	Follow
}

func NewStorage(db *sql.DB) *Storage {
	_, err := db.Exec(`
	create table if not exists users(
		id integer primary key not null
	);
		
	create table if not exists segments(
		id serial primary key,
		name varchar not null unique
	);
		
	create table if not exists history(
		id serial primary key,
		user_id integer not null,
		segment_name varchar not null,
		operation varchar not null,
		operation_time timestamp default now()
	);
		
	create table if not exists follows(
		user_id integer not null,
		segment_id integer not null,
		expire timestamp default null,
		unique (user_id, segment_id),
		foreign key (user_id) references users (id) on delete cascade,
		foreign key (segment_id) references segments (id) on delete cascade
	);`)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Storage{
		User:    repos.NewUserStorage(db),
		Segment: repos.NewSegmentStorage(db),
		History: repos.NewHistoryStorage(db),
		Follow:  repos.NewFollowStorage(db),
	}
}
