package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/service/services"
	"github.com/modaniru/avito/internal/storage/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSaveUser(t *testing.T) {
	userStorage := mocks.NewUser(t)
	idToSave := 1
	userStorage.On("SaveUser", mock.Anything, idToSave).Return(nil).Once()

	userService := services.NewUserService(userStorage)
	err := userService.SaveUser(context.Background(), idToSave)
	require.NoError(t, err)
}

func TestSaveUserWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)
	idToSave := 1
	userStorage.On("SaveUser", mock.Anything, idToSave).Return(errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	err := userService.SaveUser(context.Background(), idToSave)
	require.Error(t, err)
}

func TestGetUsers(t *testing.T) {
	userStorage := mocks.NewUser(t)
	returning := []entity.User{{Id: 1}, {Id: 2}, {Id: 3}}
	excepted := []entity.User{{Id: 1}, {Id: 2}, {Id: 3}}

	userStorage.On("GetUsers", mock.Anything).Return(returning, nil).Once()

	userService := services.NewUserService(userStorage)
	actual, err := userService.GetUsers(context.Background())
	require.NoError(t, err)
	require.Equal(t, excepted, actual)
}

func TestGetUsersWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)

	userStorage.On("GetUsers", mock.Anything).Return(nil, errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	actual, err := userService.GetUsers(context.Background())
	require.Error(t, err)
	require.Nil(t, actual)
}

func TestDeleteUser(t *testing.T) {
	userStorage := mocks.NewUser(t)
	idToDelete := 1

	userStorage.On("DeleteUser", mock.Anything, idToDelete).Return(nil).Once()

	userService := services.NewUserService(userStorage)
	err := userService.DeleteUser(context.Background(), idToDelete)
	require.NoError(t, err)
}

func TestDeleteUserWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)
	idToDelete := 1

	userStorage.On("DeleteUser", mock.Anything, idToDelete).Return(errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	err := userService.DeleteUser(context.Background(), idToDelete)
	require.Error(t, err)
}

func TestFollowSegments(t *testing.T) {
	userStorage := mocks.NewUser(t)
	userId := 1
	segments := []string{"segment_1", "segment_2"}

	userStorage.On("FollowToSegments", mock.Anything, userId, segments, mock.Anything).Return(nil).Once()

	userService := services.NewUserService(userStorage)
	err := userService.FollowToSegments(context.Background(), userId, segments, nil)
	require.NoError(t, err)
}

func TestFollowSegmentsWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)
	userId := 1
	segments := []string{"segment_1", "segment_2"}

	userStorage.On("FollowToSegments", mock.Anything, userId, segments, mock.Anything).Return(errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	err := userService.FollowToSegments(context.Background(), userId, segments, nil)
	require.Error(t, err)
}

func TestUnFollowSegments(t *testing.T) {
	userStorage := mocks.NewUser(t)
	userId := 1
	segments := []string{"segment_1", "segment_2"}

	userStorage.On("UnFollowToSegments", mock.Anything, userId, segments).Return(nil).Once()

	userService := services.NewUserService(userStorage)
	err := userService.UnFollowToSegments(context.Background(), userId, segments)
	require.NoError(t, err)
}

func TestUnFollowSegmentsWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)
	userId := 1
	segments := []string{"segment_1", "segment_2"}

	userStorage.On("UnFollowToSegments", mock.Anything, userId, segments).Return(errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	err := userService.UnFollowToSegments(context.Background(), userId, segments)
	require.Error(t, err)
}

func TestGetUserSegments(t *testing.T) {
	userStorage := mocks.NewUser(t)
	userId := 1

	segments := []entity.Follows{{Id: 1, Name: "segment_1", Expire: nil}, {Id: 2, Name: "segment_2", Expire: nil}}

	userStorage.On("GetUserSegments", mock.Anything, userId).Return([]entity.Follows{{Id: 1, Name: "segment_1", Expire: nil}, {Id: 2, Name: "segment_2", Expire: nil}}, nil).Once()

	userService := services.NewUserService(userStorage)
	list, err := userService.GetUserSegments(context.Background(), userId)
	require.NoError(t, err)
	require.NotNil(t, list)
	require.Len(t, list, 2)
	require.Equal(t, segments, list)
}

func TestGetUserSegmentsWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)
	userId := 1

	userStorage.On("GetUserSegments", mock.Anything, userId).Return(nil, errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	list, err := userService.GetUserSegments(context.Background(), userId)
	require.Nil(t, list)
	require.Error(t, err)
}

func TestFollowRandomUsers(t *testing.T) {
	userStorage := mocks.NewUser(t)

	segmentName := "test"
	percent := 0.5

	userStorage.On("FollowRandomUsers", mock.Anything, segmentName, percent).Return(10, nil).Once()

	userService := services.NewUserService(userStorage)
	rowsAffected, err := userService.FollowRandomUsers(context.Background(), segmentName, percent)
	require.NoError(t, err)
	require.Equal(t, rowsAffected, 10)
}

func TestFollowRandomUsersWithError(t *testing.T) {
	userStorage := mocks.NewUser(t)

	segmentName := "test"
	percent := 0.5

	userStorage.On("FollowRandomUsers", mock.Anything, segmentName, percent).Return(0, errors.New("error")).Once()

	userService := services.NewUserService(userStorage)
	rowsAffected, err := userService.FollowRandomUsers(context.Background(), segmentName, percent)
	require.Error(t, err)
	require.Equal(t, rowsAffected, 0)
}
