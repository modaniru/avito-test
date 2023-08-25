package repos

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveUser(t *testing.T) {
	storage := NewUserStorage(db)
	insertUsers(storage, t)
	defer clearDB()
}

func TestGetUsersWhenEmpty(t *testing.T) {
	storage := NewUserStorage(db)
	res, err := storage.GetUsers(context.Background())
	require.NoError(t, err)
	require.Empty(t, res)
}

func TestGetUsers(t *testing.T) {
	storage := NewUserStorage(db)
	insertUsers(storage, t)
	defer clearDB()

	res, err := storage.GetUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, res, 10)
}

func TestDeleteUser(t *testing.T) {
	storage := NewUserStorage(db)
	insertUsers(storage, t)
	defer clearDB()
	id := 1

	err := storage.DeleteUser(context.Background(), id)
	require.NoError(t, err)

	res, err := storage.GetUsers(context.Background())
	require.NoError(t, err)

	for _, i := range res {
		require.NotEqual(t, id, i)
	}
}

func TestAddSegments(t *testing.T) {
	userStorage := NewUserStorage(db)
	segmentStorage := NewSegmentStorage(db)
	insertUsers(userStorage, t)
	insertSegments(segmentStorage, t)

	err := userStorage.FollowToSegments(context.Background(), 1, []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"})
	require.NoError(t, err)
}

func insertUsers(storage *UserStorage, t *testing.T) {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, id := range ids {
		err := storage.SaveUser(context.Background(), id)
		require.NoError(t, err)
	}
}
