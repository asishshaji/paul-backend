package repository

import (
	"backend/dto"
	"backend/model"
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	CheckIfUserWithUsernameExists(ctx context.Context, username string) error

	CheckIfUserWithNameAndPasswordExists(ctx context.Context, user dto.UserRegDto) error
	// FindAll() ([]entity.User, error)
}
