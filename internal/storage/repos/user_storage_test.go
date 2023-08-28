package repos

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestSaveUser(t *testing.T) {
// 	storage := NewUserStorage(db)
// 	insertUsers(storage, t)
// 	defer clearDB()
// }

// func TestGetUsersWhenEmpty(t *testing.T) {
// 	storage := NewUserStorage(db)
// 	res, err := storage.GetUsers(context.Background())
// 	require.NoError(t, err)
// 	require.Empty(t, res)
// }

// func TestGetUsers(t *testing.T) {
// 	storage := NewUserStorage(db)
// 	insertUsers(storage, t)
// 	defer clearDB()

// 	res, err := storage.GetUsers(context.Background())
// 	require.NoError(t, err)
// 	require.Len(t, res, 10)
// }

// func TestDeleteUser(t *testing.T) {
// 	storage := NewUserStorage(db)
// 	insertUsers(storage, t)
// 	defer clearDB()
// 	id := 1

// 	err := storage.DeleteUser(context.Background(), id)
// 	require.NoError(t, err)

// 	res, err := storage.GetUsers(context.Background())
// 	require.NoError(t, err)

// 	for _, i := range res {
// 		require.NotEqual(t, id, i)
// 	}
// }

// func TestAddSegments(t *testing.T) {
// 	userStorage := NewUserStorage(db)
// 	segmentStorage := NewSegmentStorage(db)
// 	defer clearDB()

// 	insertUsers(userStorage, t)
// 	insertSegments(segmentStorage, t)

// 	segments := []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"}
// 	err := userStorage.FollowToSegments(context.Background(), 1, segments)
// 	require.NoError(t, err)
// }

// func TestAddSegmentsTwiceThrowError(t *testing.T) {
// 	userStorage := NewUserStorage(db)
// 	segmentStorage := NewSegmentStorage(db)
// 	defer clearDB()

// 	insertUsers(userStorage, t)
// 	insertSegments(segmentStorage, t)
// 	err := userStorage.FollowToSegments(context.Background(), 1, []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"})
// 	require.NoError(t, err)
// 	err = userStorage.FollowToSegments(context.Background(), 1, []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"})
// 	require.Error(t, err)
// }

// func TestAddNonExistsSegment(t *testing.T) {
// 	userStorage := NewUserStorage(db)
// 	segmentStorage := NewSegmentStorage(db)
// 	defer clearDB()

// 	insertUsers(userStorage, t)
// 	insertSegments(segmentStorage, t)
// 	err := userStorage.FollowToSegments(context.Background(), 1, []string{"test", "test2"})
// 	require.Error(t, err)
// }

// func TestUnFollowSegment(t *testing.T) {
// 	userStorage := NewUserStorage(db)
// 	segmentStorage := NewSegmentStorage(db)
// 	defer clearDB()

// 	insertUsers(userStorage, t)
// 	insertSegments(segmentStorage, t)
// 	err := userStorage.FollowToSegments(context.Background(), 1, []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"})
// 	require.NoError(t, err)
// 	err = userStorage.UnFollowToSegments(context.Background(), 1, []string{"AVITO_PERFORMANCE_VAS"})
// 	require.NoError(t, err)
// 	list, err := userStorage.GetUserSegments(context.Background(), 1)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, list)
// 	require.Equal(t, "AVITO_VOICE_MESSAGES", list[0].Name)
// }

// func TestUnFollowUnExistsSegment(t *testing.T) {
// 	userStorage := NewUserStorage(db)
// 	segmentStorage := NewSegmentStorage(db)
// 	defer clearDB()

// 	insertUsers(userStorage, t)
// 	insertSegments(segmentStorage, t)
// 	err := userStorage.FollowToSegments(context.Background(), 1, []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"})
// 	require.NoError(t, err)
// 	err = userStorage.UnFollowToSegments(context.Background(), 1, []string{"test"})
// 	require.Error(t, err)
// }

// func insertUsers(storage *UserStorage, t *testing.T) {
// 	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	for _, id := range ids {
// 		err := storage.SaveUser(context.Background(), id)
// 		require.NoError(t, err)
// 	}
// }
