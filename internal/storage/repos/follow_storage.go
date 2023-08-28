package repos

import (
	"context"
	"database/sql"
	"fmt"
)

type FollowStorage struct {
	db *sql.DB
}

func NewFollowStorage(db *sql.DB) *FollowStorage {
	return &FollowStorage{db: db}
}

// Не тестируется sqlite из-за now()
func (f *FollowStorage) DeleteExpiredFollows() (int, error) {
	op := "internal.storage.repos.FollowStorage.DeleteExpiredFollows"
	getExpireFollow := "select f.user_id, s.name from follows as f inner join segments as s on f.segment_id = s.id where f.expire is not NULL and f.expire < now();"
	deleteExpireFollows := "delete from follows where expire is not null and expire < now();"
	saveHistoryQuery := "insert into history (user_id, segment_name, operation) values ($1, $2, $3);"

	type follow struct {
		id   int
		name string
	}

	rows, err := f.db.Query(getExpireFollow)
	if err != nil {
		return 0, fmt.Errorf("%s exec error: %w", op, err)
	}

	defer rows.Close()
	follows := []follow{}
	for rows.Next() {
		var f follow
		err := rows.Scan(&f.id, &f.name)
		if err != nil {
			return 0, fmt.Errorf("%s scan error: %w", op, err)
		}
		follows = append(follows, f)
	}

	tx, err := f.db.BeginTx(context.Background(), nil)
	if err != nil {
		return 0, fmt.Errorf("%s begin tx error: %w", op, err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(deleteExpireFollows)
	if err != nil {
		return 0, fmt.Errorf("%s exec error: %w", op, err)
	}

	for _, f := range follows {
		_, err := tx.Exec(saveHistoryQuery, f.id, f.name, UnFollowOperation)
		if err != nil {
			return 0, fmt.Errorf("%s exec error: %w", op, err)
		}
	}

	tx.Commit()
	return len(follows), nil
}
