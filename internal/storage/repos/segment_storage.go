package repos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/modaniru/avito/internal/entity"
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

func (s *SegmentStorage) DeleteSegment(ctx context.Context, name string) error {
	op := "internal.storage.repos.SegmentStorage.DeleteSegment"
	query := "delete from segments where name = $1;"

	_, err := s.db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("%s exec error: %w", op, err)
	}

	return nil
}
