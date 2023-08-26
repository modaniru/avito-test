package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/modaniru/avito/internal/entity"
)

var (
	ErrSegmentNotFound      = errors.New("segment was not found")
	ErrSegmentAlreadyExists = errors.New("segment already exists")
)

type SegmentStorage struct {
	db *sql.DB
}

func NewSegmentStorage(db *sql.DB) *SegmentStorage {
	return &SegmentStorage{db: db}
}

func (s *SegmentStorage) SaveSegment(ctx context.Context, name string) (int, error) {
	op := "internal.storage.repos.SegmentStorage.SaveSegment"
	query := "insert into segments (name) values ($1) returning id;"

	var id int
	err := s.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return 0, ErrSegmentAlreadyExists
			}
		}
		return 0, fmt.Errorf("%s exec error: %w", op, err)
	}

	return id, nil
}

func (s *SegmentStorage) GetSegments(ctx context.Context) ([]entity.Segment, error) {
	op := "internal.storage.repos.SegmentStorage.GetSegments"
	query := "select id, name from segments;"

	segments := make([]entity.Segment, 0)
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s exec error: %w", op, err)
	}

	defer rows.Close()
	for rows.Next() {
		var segment entity.Segment
		err = rows.Scan(&segment.Id, &segment.Name)
		if err != nil {
			return nil, fmt.Errorf("%s scan error: %w", op, err)
		}

		segments = append(segments, segment)
	}
	return segments, nil
}

// TODO test segment not found
func (s *SegmentStorage) DeleteSegment(ctx context.Context, name string) error {
	op := "internal.storage.repos.SegmentStorage.DeleteSegment"
	query := "delete from segments where name = $1 returning id;"
	getUsers := "select f.user_id from follows as f inner join segments as s on f.segment_id = s.id where s.name = $1;"
	saveHistoryQuery := "insert into history (user_id, segment_name, operation) values ($1, $2, $3);"

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s begin tx error: %w", op, err)
	}
	defer tx.Rollback()
	rows, err := tx.QueryContext(ctx, getUsers, name)
	if err != nil {
		return fmt.Errorf("%s exec error: %w", op, err)
	}

	defer rows.Close()
	usersId := []int{}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return fmt.Errorf("%s scan error: %w", op, err)
		}
		usersId = append(usersId, id)
	}

	for _, id := range usersId {
		_, err := tx.ExecContext(ctx, saveHistoryQuery, id, name, UnFollowOperation)
		if err != nil {
			return fmt.Errorf("%s exec error: %w", op, err)
		}
	}

	var id int
	err = tx.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrSegmentNotFound
		}
		return fmt.Errorf("%s exec error: %w", op, err)
	}
	tx.Commit()
	return nil
}
