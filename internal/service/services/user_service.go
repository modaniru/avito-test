package services

import (
	"context"
	"fmt"

	"github.com/modaniru/avito/internal/entity"
	"github.com/modaniru/avito/internal/storage"
)

type UserService struct {
	userStorage storage.User
}

func NewUserService(userStorage storage.User) *UserService {
	return &UserService{userStorage: userStorage}
}

func (u *UserService) SaveUser(ctx context.Context, id int) error {
	op := "internal.service.services.UserService.SaveUser"

	err := u.userStorage.SaveUser(ctx, id)
	if err != nil {
		return fmt.Errorf("%s save user error %w", op, err)
	}
	return nil
}

func (u *UserService) GetUsers(ctx context.Context) ([]entity.User, error) {
	op := "internal.service.services.UserService.GetUsers"

	users, err := u.userStorage.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s get users error %w", op, err)
	}
	return users, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	op := "internal.service.services.UserService.DeleteUser"

	err := u.userStorage.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("%s delete user error %w", op, err)
	}
	return nil
}

func (u *UserService) FollowToSegments(ctx context.Context, userId int, segments []string) error {
	op := "internal.service.services.UserService.FollowToSegments"

	err := u.userStorage.FollowToSegments(ctx, userId, segments)
	if err != nil {
		return fmt.Errorf("%s follow to segments error %w", op, err)
	}
	return nil
}

func (u *UserService) UnFollowToSegments(ctx context.Context, userId int, segments []string) error {
	op := "internal.service.services.UserService.UnFollowToSegments"

	err := u.userStorage.UnFollowToSegments(ctx, userId, segments)
	if err != nil {
		return fmt.Errorf("%s unfollow segments error %w", op, err)
	}
	return nil
}

func (u *UserService) GetUserSegments(ctx context.Context, userId int) ([]entity.Segment, error) {
	op := "internal.service.services.UserService.GetUserSegments"

	segments, err := u.userStorage.GetUserSegments(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s get user segments error %w", op, err)
	}
	return segments, nil
}
