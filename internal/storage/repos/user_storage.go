package repos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/modaniru/avito/internal/entity"
)

const (
	FollowOperation = "ADD SEGMENT"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) SaveUser(ctx context.Context, userId int) error {
	op := "internal.storage.repos.UserStorage.SaveUser"
	query := "insert into users (id) values ($1);"

	_, err := u.db.ExecContext(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("%s exec error: %w", op, err)
	}

	return nil
}

func (u *UserStorage) GetUsers(ctx context.Context) ([]entity.User, error) {
	op := "internal.storage.repos.UserStorage.GetUsers"
	query := "select id from users;"

	users := make([]entity.User, 0)
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s exec error: %w", op, err)
	}

	defer rows.Close()
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id)
		if err != nil {
			return nil, fmt.Errorf("%s scan error: %w", op, err)
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *UserStorage) DeleteUser(ctx context.Context, userId int) error {
	op := "internal.storage.repos.UserStorage.DeleteUser"
	query := "delete from users where id = $1;"

	_, err := u.db.ExecContext(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("%s exec error: %w", op, err)
	}

	return nil
}

func (u *UserStorage) FollowToSegments(ctx context.Context, userId int, segments []string) error {
	op := "internal.storage.repos.UserStorage.FollowToSegments"
	saveFollowQuery := "insert into follows (user_id, segment_name) values ($1, $2);"
	saveHistoryQuery := "insert into history (user_id, segment_name, operation) values ($1, $2, $3)"

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s begin tx error: %w", op, err)
	}

	defer tx.Rollback()
	for _, segment := range segments {
		_, err = tx.ExecContext(ctx, saveFollowQuery, userId, segment)
		if err != nil {
			return fmt.Errorf("%s exec error: %w", op, err)
		}
		_, err := tx.ExecContext(ctx, saveHistoryQuery, userId, segment, FollowOperation)
		if err != nil {
			return fmt.Errorf("%s exec error: %w", op, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s commit transaction error: %w", op, err)
	}
	return nil
}
