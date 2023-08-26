package repos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/modaniru/avito/internal/entity"
)

type HistoryStorage struct {
	db *sql.DB
}

func NewHistoryStorage(db *sql.DB) *HistoryStorage {
	return &HistoryStorage{db: db}
}

func (h *HistoryStorage)  GetHistory(ctx context.Context) ([]entity.History, error){
	op := "internal.storage.repos.HistoryStorage.GetHistory"
	query := "select user_id, segment_name, operation, operation_time from history order by operation_time desc;"

	rows, err := h.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s exec error: %w", op, err)
	}

	defer rows.Close()
	history := make([]entity.History, 0)
	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.UserId, &h.SegmentName, &h.Operation, &h.OperationTime)
		if err != nil {
			return nil, fmt.Errorf("%s scan error: %w", op, err)
		}
		history = append(history, h)
	}
	return history, nil
}
